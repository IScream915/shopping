package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strings"
)

func StartWith(f validator.FieldLevel) bool {
	// 获取字段值
	// f.Field() 返回当前正在验证的字段的 reflect.Value
	field := f.Field().String()
	// 获取标签中的值
	// 当你在结构体标签中写 startWith=abc 时，f.Param() 会返回 "abc"
	tagValue := f.Param()
	return strings.HasPrefix(field, tagValue)
}

func RegisterCustomValidation() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这是一个可选的布尔参数（variadic 参数）
		//用于指定在字段值为 null 或空时是否也执行验证。通常默认是不调用验证函数，但你可以通过传入 true 来改变这个行为。
		_ = validate.RegisterValidation("startWith", StartWith)
	}
}
