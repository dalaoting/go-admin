package media

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/pkg/models"
	"go-admin/pkg/service/mediaService"
	"go-admin/pkg/uuid"
	"go-admin/tools"
	"net/http"
	"strconv"
)

type Media struct {
	apis.Api
}

// swagger:parameters  UploadFileRequest
type UploadFileRequest struct {
	// 上传文件类型 1:mp4 2:mp3 3:jpg 4:log
	// required: true
	Type int64 `form:"type" binding:"required"`
	// 当为媒体文件时，代表媒体的宽
	Width int `form:"width"`
	// 当为媒体文件时，代表媒体的高
	Height int `form:"height"`
	// 当为视频时，代表视频的时长
	MediaTime int `form:"media_time"`
}

// swagger:model UploadFileResponse
type UploadFileResponse struct {
	// 文件的ID
	Id int64 `json:"id"`
	// 文件名称
	Name string `json:"name"`
	// 文件的链接
	Url string `json:"url"`
	// 文件的大小
	Size int64 `json:"size"`
}

func (e *Media) UploadFile(c *gin.Context) {
	req := &UploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数错误")
		return
	}
	// 获取上传类型
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}
	mediaType, err := mediaService.GetMediaType(db, req.Type)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusUnprocessableEntity, err, "获取文件类型失败")
		return
	}
	_, header, err := c.Request.FormFile("file")

	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusUnprocessableEntity, err, "文件为空")
		return
	}

	if header.Size > 1024*1024*5 {
		e.Error(c, http.StatusUnprocessableEntity, err, "文件过大")
		return
	}

	//uploader, err := objectStorageService.GetOss()
	//if err != nil {
	//	common.ErrorLog("UploadFile", h.GetRequestId(), "GetOss error", err.Error())
	//	h.ErrResponse(exp.New("上传失败", exp.CodeServiceBusy))
	//	return
	//}
	u, _ := uuid.UUID()
	uid := strconv.FormatUint(u, 10)
	//option := objectStorageService.SetFileOption(mediaType.Folder, mediaType.Suffix)
	//f, err := uploader.UploadFile(file, uid, option)
	//if err != nil {
	//	common.ErrorLog("UploadFile", h.GetRequestId(), "UploadFile error", err.Error())
	//	h.ErrResponse(exp.New("上传失败", exp.CodeServiceBusy))
	//	return
	//}
	objKey := mediaType.Folder + "/" + uid + mediaType.Suffix
	//if err := o.bucket.PutObject(objKey, fd); err != nil {
	//	return nil, err
	//}
	if err = c.SaveUploadedFile(header, "/usr/local/nginx/html/img/"+objKey); err != nil {
		log.Error(err)
		e.Error(c, http.StatusUnprocessableEntity, err, "上传失败")
		return
	}

	path := "https://rt.oldthree.cn" + "/" + objKey

	// 保存媒体文件
	media := &models.Media{
		MediaType:   int(req.Type),
		OriginalURL: path,
		Width:       req.Width,
		Height:      req.Height,
		MediaTime:   req.MediaTime,
		MediaSize:   int(header.Size),
		IP:          c.Request.RemoteAddr,
	}

	response := &UploadFileResponse{
		Name: uid,
		Url:  path,
		Size: 0,
	}
	if err = mediaService.Save(db, media); err == nil {
		response.Id = media.ID
	}
	e.OK(c, response, "上传成功")
}
