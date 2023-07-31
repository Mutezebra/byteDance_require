package api

import (
	"AlittleRequire/pkg/ctl"
	"AlittleRequire/pkg/e"
	"AlittleRequire/pkg/logger"
	"AlittleRequire/service"
	"AlittleRequire/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTopicHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateTopicReq
		var code int
		err := c.ShouldBind(&req)
		if err != nil {
			logger.LogrusObj.Info(err)
			code = e.InvalidParam
			c.JSON(http.StatusBadRequest, ErrorResponse(err, code))
			return
		}

		l := service.GetTaskSrv()
		_, code, err = l.CreateTopic(c, &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err, code))
			return
		}
		c.JSON(200, ctl.RespSuccess())
	}
}

func CreatePostHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreatePostReq
		var code int
		err := c.ShouldBind(&req)
		if err != nil {
			logger.LogrusObj.Info(err)
			code = e.InvalidParam
			c.JSON(http.StatusBadRequest, ErrorResponse(err, code))
			return
		}

		l := service.GetTaskSrv()
		_, code, err = l.CreatePost(c, &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err, code))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess())
	}
}

func GetTopicInfoHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.GetTopicInfoReq
		var code int
		err := c.ShouldBind(&req)
		if err != nil {
			logger.LogrusObj.Info(err)
			code = e.InvalidParam
			c.JSON(http.StatusBadRequest, ErrorResponse(err, code))
			return
		}

		l := service.GetTaskSrv()
		resp, code, err := l.GetTopicInfo(c, &req)
		if err != nil {
			logger.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err, code))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
	}
}
