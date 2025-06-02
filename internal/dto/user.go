package dto

type AccountLoginReq struct {
	Account  string `json:"account" binding:"required"`  // 账号
	Password string `json:"password" binding:"required"` // 密码
}

type EmailSendReq struct {
	Email string `json:"email" binding:"required,email"` //邮箱
}

type EmailLoginReq struct {
	Email            string `json:"email" binding:"required,email"`       // 邮箱
	VerificationCode string `json:"verification_code" binding:"required"` // 验证码
}

type CreateUserReq struct {
	Account  string `json:"account" binding:"required"`           // 账号
	Email    string `json:"email" binding:"required,email"`       // 邮箱
	NickName string `json:"nickname" binding:"required"`          // 昵称
	Password string `json:"password" binding:"required"`          // 密码
	Age      uint8  `json:"age" binding:"required,gte=0,lte=100"` // 年龄
	Sex      string `json:"sex" binding:"required,oneof=0 1"`     // 性别
}

type UpdateUserReq struct {
	Email    string `json:"email"`    // 邮箱
	NickName string `json:"nickname"` // 昵称
	Password string `json:"password"` // 密码
	Age      uint8  `json:"age"`      // 年龄
	Sex      string `json:"sex"`      // 性别
}

type DeleteUserReq struct {
	IDs []uint64 `json:"ids" binding:"required"` // 用户ID
}
