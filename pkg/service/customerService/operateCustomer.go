package customerService

import (
	"fmt"
	"go-admin/common/global"
	"go-admin/pkg/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func OperateCustomer(tasks ...*models.CustomerOperation) error {
	tx := global.Eloquent.Set("gorm:query_option", "FOR UPDATE")

	if err := OperateCustomerWithTx(tx, tasks...); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func OperateCustomerWithTx(tx *gorm.DB, tasks ...*models.CustomerOperation) error {
	customers, err := findCustomersByOperation(tx, tasks)
	if err != nil {
		return err
	}

	logs := make([]*models.CustomerLog, 0)
	for _, task := range tasks {
		if task.Amount == 0 {
			continue
		}
		if _, ok := customers[task.Key()]; !ok {
			continue
		}
		log, err := newCustomerLog(customers[task.Key()], task)
		if err != nil {
			return err
		}
		logs = append(logs, log)
	}

	if err = insertCustomerLogBatch(tx, logs); err != nil {
		return err
	}

	if err = updateCustomersBatch(tx, customers); err != nil {
		return err
	}

	return nil
}

func findCustomersByOperation(db *gorm.DB, tasks []*models.CustomerOperation) (map[string]*models.Customer, error) {
	mustFind := make(map[int]struct{})
	for _, t := range tasks {
		mustFind[t.CustomerId] = struct{}{}
	}

	if db == nil {
		db = global.Eloquent
	}

	where := fmt.Sprintf("(customer_id='%v')", tasks[0].CustomerId)
	for i := 1; i < len(tasks); i++ {
		where += fmt.Sprintf("or (user_id='%v')", tasks[i].CustomerId)
	}

	customers := make([]*models.Customer, 0)
	err := db.Where(where).Find(&customers).Error
	if err != nil {
		return nil, err
	}

	ret := make(map[string]*models.Customer)
	for _, c := range customers {
		delete(mustFind, int(c.ID))
		ret[strconv.Itoa(int(c.ID))] = c
	}

	for key, _ := range mustFind {
		var (
			customerId = key
		)
		c := &models.Customer{}
		c.ID = uint(customerId)
		if err := db.Create(c).Error; err != nil {
			return nil, models.ErrAccountNotFound
		}
		ret[strconv.Itoa(customerId)] = c
	}

	return ret, nil
}

func newCustomerLog(customer *models.Customer, operate *models.CustomerOperation) (*models.CustomerLog, error) {
	log := &models.CustomerLog{
		CustomerId:   int64(operate.CustomerId),
		OpType:       operate.OpType,
		BsType:       operate.BsType,
		Detail:       operate.Detail,
		Ext:          operate.Ext,
		BeforeAmount: customer.AvailAmount,
	}
	log.CreateBy = operate.CreateBy

	if operate.OpType == models.OpTypeAdd {
		customer.AvailAmount = customer.AvailAmount + operate.Amount
		customer.TotalAmount = customer.TotalAmount + operate.Amount
	} else {
		customer.AvailAmount = customer.AvailAmount - operate.Amount
		if customer.AvailAmount-customer.PrepayAmount < 0 {
			return nil, models.ErrAmountNotEnough
		}
	}
	log.Amount = operate.Amount
	log.AfterAmount = customer.AvailAmount
	return log, nil
}

func insertCustomerLogBatch(db *gorm.DB, logs []*models.CustomerLog) error {
	if len(logs) == 0 {
		return nil
	}

	if db == nil {
		db = global.Eloquent
	}

	sql := "insert into customer_log (customer_id, amount, op_type, " +
		"bs_type, detail, before_amount, after_amount, ext, create_by, created_at) values "
	values := []string{}
	for _, log := range logs {
		values = append(values, fmt.Sprintf("('%v', %v, %v, %v, '%v', %v, %v, '%v', %v, now())",
			log.CustomerId, log.Amount, log.OpType, log.BsType, log.Detail,
			log.BeforeAmount, log.AfterAmount, log.Ext, log.CreateBy))
	}

	sql += strings.Join(values, ", ")
	// println(sql)
	return db.Exec(sql).Error
}

func updateCustomersBatch(db *gorm.DB, customers map[string]*models.Customer) error {
	if len(customers) == 0 {
		return nil
	}

	if db == nil {
		db = global.Eloquent
	}

	ids := make([]string, 0)
	sql := "update account set avail_amount = case id "
	for _, customer := range customers {
		ids = append(ids, strconv.Itoa(int(customer.ID)))
		sql += fmt.Sprintf("when %v then '%v' ", customer.ID, customer.AvailAmount)
	}
	sql += "end, "
	sql += "prepay_amount = case id "
	for _, customer := range customers {
		sql += fmt.Sprintf("when %v then '%v' ", customer.ID, customer.PrepayAmount)
	}
	sql += "end, "
	sql += "total_amount = case id "
	for _, customer := range customers {
		sql += fmt.Sprintf("when %v then '%v' ", customer.ID, customer.TotalAmount)
	}
	sql += fmt.Sprintf("end where id in (%v)", strings.Join(ids, ","))

	return db.Exec(sql).Error
}
