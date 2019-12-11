package entity

// 管理员数据
type AdminEntity struct {
	Id         int64  `orm:"column(id);pk" json:"id,omitempty"`
	User       string `orm:"column(user)" json:"user,omitempty"`
	Password   string `orm:"column(password)" json:"password,omitempty"`
	Sex        int8   `orm:"column(sex)" json:"sex,omitempty"`
	Age        int8  `orm:"column(age)" json:"age,omitempty"`
	Phone      string  `orm:"column(phone)" json:"phone,omitempty"`
	Name       string `orm:"column(name)" json:"name,omitempty"`
	CreateTime int64  `orm:"column(createTime)" json:"createTime,omitempty"`
	UpdateTime int64  `orm:"column(updateTime)" json:"updateTime,omitempty"`
	DeleteTime int64  `orm:"column(deleteTime)" json:"deleteTime,omitempty"`
}

func (e *AdminEntity) TableName() string {
	return "t_admin_entity"
}

const TABLE_AdminEntity = "t_admin_entity"

const COLUMN_AdminEntity_Id = "id"
const COLUMN_AdminEntity_User = "user"
const COLUMN_AdminEntity_Password = "password"
const COLUMN_AdminEntity_Sex = "sex"
const COLUMN_AdminEntity_Age = "age"
const COLUMN_AdminEntity_Phone = "phone"
const COLUMN_AdminEntity_Name = "name"
const COLUMN_AdminEntity_CreateTime = "createTime"
const COLUMN_AdminEntity_UpdateTime = "updateTime"
const COLUMN_AdminEntity_DeleteTime = "deleteTime"

const ATTRIBUTE_AdminEntity_Id = "Id"
const ATTRIBUTE_AdminEntity_User = "User"
const ATTRIBUTE_AdminEntity_Password = "Password"
const ATTRIBUTE_AdminEntity_Sex = "Sex"
const ATTRIBUTE_AdminEntity_Age= "Age"
const ATTRIBUTE_AdminEntity_Phone = "Phone"
const ATTRIBUTE_AdminEntity_Name = "Name"
const ATTRIBUTE_AdminEntity_CreateTime = "CreateTime"
const ATTRIBUTE_AdminEntity_UpdateTime = "UpdateTime"
const ATTRIBUTE_AdminEntity_DeleteTime = "DeleteTime"
