package constant

const (
	DemandAssignCancel       = -1 // 已取消
	DemandAssignPending      = 0  // 待开始
	DemandAssignProcessing   = 1  // 进行中
	DemandAssignOrder        = 4  // 已下单
	DemandAssignComment      = 7  // 已留评
	DemandAssignOrderSettle  = 10 // 订单费用结算
	DemandAssignRewardSettle = 13 // 佣金结算
	DemandAssignFinish       = 15 // 完成
	DemandAssignRefund       = 18 //  退款
)
