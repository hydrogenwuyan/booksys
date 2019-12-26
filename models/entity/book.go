package entity

// 图书数据
type BookEntity struct {
	Id         int64  `orm:"column(id);pk" json:"id,omitempty"`
	IsBorrow   int8   `orm:"column(isBorrow)" json:"isBorrow,omitempty"`
	StuId      int64  `orm:"column(stuId)" json:"stuId,omitempty"`
	StuUser    string `orm:"column(stuUser)" json:"stuUser,omitempty"`
	ExpireTime int64  `orm:"column(expireTime)" json:"expireTime,omitempty"`
	Type       int32  `orm:"column(type)" json:"type,omitempty"`
	Name       string `orm:"column(name)" json:"name,omitempty"`
	Author     string `orm:"column(author)" json:"author,omitempty"`
	Press      string `orm:"column(press)" json:"press,omitempty"`
	CreateTime int64  `orm:"column(createTime)" json:"createTime,omitempty"`
	UpdateTime int64  `orm:"column(updateTime)" json:"updateTime,omitempty"`
	DeleteTime int64  `orm:"column(deleteTime)" json:"deleteTime,omitempty"`
}

func (e *BookEntity) TableName() string {
	return "t_book_entity"
}

const TABLE_BookEntity = "t_book_entity"

const COLUMN_BookEntity_Id = "id"
const COLUMN_BookEntity_IsBorrow = "isBorrow"
const COLUMN_BookEntity_StuId = "stuId"
const COLUMN_BookEntity_StuUser = "stuUser"
const COLUMN_BookEntity_ExpireTime = "expireTime"
const COLUMN_BookEntity_Type = "type"
const COLUMN_BookEntity_Name = "name"
const COLUMN_BookEntity_Author = "author"
const COLUMN_BookEntity_Press = "press"
const COLUMN_BookEntity_CreateTime = "createTime"
const COLUMN_BookEntity_UpdateTime = "updateTime"
const COLUMN_BookEntity_DeleteTime = "deleteTime"
