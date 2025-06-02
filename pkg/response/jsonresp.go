package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code       int         `json:"code"`        // 错误码((0:成功, 1:失败, >1:错误码))
	Msg        string      `json:"msg"`         // 提示信息
	Data       interface{} `json:"data"`        // 返回数据(业务接口定义具体数据结构)
	ServerTime int64       `json:"server_time"` // 服务器时间
	Version    string      `json:"version"`     // 版本号
}

type JsonResponseFunc func(obj *JsonResponse)

type JsonResponseFuncs []JsonResponseFunc

func (funcs JsonResponseFuncs) apply(obj *JsonResponse) {
	for _, f := range funcs {
		f(obj)
	}
}

func WithCode(code int) JsonResponseFunc {
	return func(obj *JsonResponse) {
		obj.Code = code
	}
}

func WithMsg(msg string) JsonResponseFunc {
	return func(obj *JsonResponse) {
		obj.Msg = msg
	}
}

func WithData(data interface{}) JsonResponseFunc {
	return func(obj *JsonResponse) {
		obj.Data = data
	}
}

func WithPageData(list interface{}, total int64) JsonResponseFunc {
	return func(obj *JsonResponse) {
		obj.Data = map[string]interface{}{
			"list":  list,
			"total": total,
		}
	}
}

func WithErr(err error) JsonResponseFunc {
	return func(obj *JsonResponse) {
		if err == nil {
			return
		} else {
			obj.Code = -1
			obj.Msg = err.Error()
		}
	}
}

func Json(c *gin.Context, funcs ...JsonResponseFunc) {
	resp := &JsonResponse{
		Code: 0,
		Msg:  "success",
		//Version:    version.Get(),
		ServerTime: time.Now().Unix(),
	}
	JsonResponseFuncs(funcs).apply(resp)
	//if gin.Mode() == gin.DebugMode {
	//	marshal, _ := json.Marshal(resp)
	//	log.ZDebug("接口: %s | 响应: %s", c.Request.RequestURI, string(marshal))
	//}
	c.JSON(http.StatusOK, resp)
}
