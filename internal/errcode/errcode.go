package errcode

import "base_frame/pkg/errs"

var (
	DBRecordNotFound     = errs.New(2003, "no relevant record")     // 数据库中没有相关记录
	EntityParameterError = errs.New(2007, "entity parameter error") // 实体参数错误
)
