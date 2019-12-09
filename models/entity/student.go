package entity

// 学生数据
type StudentEntity struct {
	Id          int64  `orm:"column(id);pk" json:"id,omitempty"`
	User        string `orm:"column(user);pk" json:"user,omitempty"`
	Password    string `orm:"column(password);pk" json:"password,omitempty"`
	Sex         int8   `orm:"column(sex);pk" json:"sex,omitempty"`
	Age         int32  `orm:"column(age);pk" json:"age,omitempty"`
	Phone       int32  `orm:"column(phone);pk" json:"phone,omitempty"`
	Name        string `orm:"column(name);pk" json:"name,omitempty"`
	BorrowBooks string `orm:"column(borrowBooks);pk" json:"borrowBooks,omitempty"`
	CreateTime  int64  `orm:"column(createTime)" json:"createTime,omitempty"`
	UpdateTime  int64  `orm:"column(updateTime)" json:"updateTime,omitempty"`
	DeleteTime  int64  `orm:"column(deleteTime)" json:"deleteTime,omitempty"`
}

func (e *StudentEntity) TableName() string {
	return "t_student_entity"
}

const TABLE_StudentEntity = "t_student_entity"

const COLUMN_StudentEntity_Id = "id"
const COLUMN_StudentEntity_User = "user"
const COLUMN_StudentEntity_Password = "password"
const COLUMN_StudentEntity_Sex = "sex"
const COLUMN_StudentEntity_Age = "age"
const COLUMN_StudentEntity_Phone = "phone"
const COLUMN_StudentEntity_Name = "name"
const COLUMN_StudentEntity_BorrowBooks = "borrowBooks"
const COLUMN_StudentEntity_CreateTime = "createTime"
const COLUMN_StudentEntity_UpdateTime = "updateTime"
const COLUMN_StudentEntity_DeleteTime = "deleteTime"

const ATTRIBUTE_StudentEntity_Id = "Id"
const ATTRIBUTE_StudentEntity_User = "User"
const ATTRIBUTE_StudentEntity_Password = "Password"
const ATTRIBUTE_StudentEntity_Sex = "Sex"
const ATTRIBUTE_StudentEntity_Age = "Age"
const ATTRIBUTE_StudentEntity_Phone = "Phone"
const ATTRIBUTE_StudentEntity_Name = "Name"
const ATTRIBUTE_StudentEntity_BorrowBooks = "BorrowBooks"
const ATTRIBUTE_StudentEntity_CreateTime = "CreateTime"
const ATTRIBUTE_StudentEntity_UpdateTime = "UpdateTime"
const ATTRIBUTE_StudentEntity_DeleteTime = "DeleteTime"
