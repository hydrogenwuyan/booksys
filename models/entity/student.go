package entity

// 学生数据
type StudentEntity struct {
	Id         int64  `orm:"column(id);pk" json:"id,omitempty"`
	User       string `orm:"column(user)" json:"user,omitempty"`
	Password   string `orm:"column(password)" json:"password,omitempty"`
	IsBlack    int8   `orm:"column(isBlack)" json:"isBlack,omitempty"`
	Sex        int8   `orm:"column(sex)" json:"sex,omitempty"`
	Age        int8   `orm:"column(age)" json:"age,omitempty"`
	Phone      string `orm:"column(phone)" json:"phone,omitempty"`
	Name       string `orm:"column(name)" json:"name,omitempty"`
	CreateTime int64  `orm:"column(createTime)" json:"createTime,omitempty"`
	UpdateTime int64  `orm:"column(updateTime)" json:"updateTime,omitempty"`
	DeleteTime int64  `orm:"column(deleteTime)" json:"deleteTime,omitempty"`
}

func (e *StudentEntity) TableName() string {
	return "t_student_entity"
}

const TABLE_StudentEntity = "t_student_entity"

const COLUMN_StudentEntity_Id = "id"
const COLUMN_StudentEntity_User = "user"
const COLUMN_StudentEntity_Password = "password"
const COLUMN_StudentEntity_IsBlack = "isBlack"
const COLUMN_StudentEntity_Sex = "sex"
const COLUMN_StudentEntity_Age = "age"
const COLUMN_StudentEntity_Phone = "phone"
const COLUMN_StudentEntity_Name = "name"
const COLUMN_StudentEntity_CreateTime = "createTime"
const COLUMN_StudentEntity_UpdateTime = "updateTime"
const COLUMN_StudentEntity_DeleteTime = "deleteTime"
