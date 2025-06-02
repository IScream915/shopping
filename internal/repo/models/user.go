package models

const (
	TableNameUser = "t_users"
)

var UserColumn = &userColumn{
	ID:        "id",
	Account:   "account",
	Email:     "email",
	NickName:  "nickname",
	Password:  "password",
	Age:       "age",
	Sex:       "sex",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

type User struct {
	BaseModel
	Account  string `gorm:"column:account;type:varchar(50);not null;uniqueIndex:uniq_account;comment:账号" json:"account"` // 账号
	Email    string `gorm:"column:email;type:varchar(50);not null;uniqueIndex:uniq_email;comment:邮箱" json:"email"`       // 邮箱
	NickName string `gorm:"column:nickname;type:varchar(255);not null;comment:昵称" json:"nick_name"`                      // 昵称
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`                       // 密码
	Age      uint8  `gorm:"column:age;type:int unsigned;default:18;not null;comment:年龄" json:"age"`                      // 年龄
	Sex      string `gorm:"column:sex;type:enum('0','1');not null;comment:性别" json:"sex"`                                // 性别
}

func (*User) TableName() string { return TableNameUser }

type userColumn struct {
	ID        string
	Account   string
	Email     string
	NickName  string
	Password  string
	Age       string
	Sex       string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
