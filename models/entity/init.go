package entity

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(AdminEntity))
	orm.RegisterModel(new(BookEntity))
	orm.RegisterModel(new(StudentEntity))
}
