package response

import (
	"github.com/kataras/iris"
)

// Response : JSON Response Object
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    interface{} `json:"page,omitempty"`
	Te      string      `json:"test,omitempty"`
}

// JSON ...
func JSON(ctx iris.Context, d interface{}) {
	ctx.JSON(&Response{
		Code:    ErrorCode.SUCCESS,
		Message: "success",
		Data:    d,
	})
	ctx.Application().Logger().Info("response is success data: " + d.(string))
}

// JSONPage ...
func JSONPage(ctx iris.Context, d interface{}, p interface{}) {
	ctx.JSON(&Response{
		Code:    ErrorCode.SUCCESS,
		Message: "success",
		Data:    d,
		Page:    p,
	})
	ctx.Application().Logger().Info("response is success data: " + d.(string))
}

// JSONError ...
func JSONError(ctx iris.Context, err string) {
	ctx.JSON(&Response{
		Code:    ErrorCode.ERROR,
		Message: err,
	})
	ctx.Application().Logger().Info("response is error : " + err)
}

// JSONErrorCode ...
func JSONErrorCode(ctx iris.Context, err string, code int) {
	ctx.JSON(&Response{
		Code:    code,
		Message: err,
	})
	ctx.Application().Logger().Info("response is errorCode : " + err)
}

// JSONBad is ...
func JSONBad(ctx iris.Context, err string) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(&Response{
		Code:    ErrorCode.ERROR,
		Message: err,
	})
	ctx.Application().Logger().Info("response is bad data: " + err)
}
