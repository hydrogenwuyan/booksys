package entity

// 管理员数据
type AdminEntity struct {
	Id         int64 `orm:"column(id);pk" json:"id,omitempty"`
	CreateTime int64 `orm:"column(createTime)" json:"createTime,omitempty"`
	UpdateTime int64 `orm:"column(updateTime)" json:"updateTime,omitempty"`
	DeleteTime int64 `orm:"column(deleteTime)" json:"deleteTime,omitempty"`
}

func (e *AdminEntity) TableName() string {
	return "t_admin_entity"
}

const TABLE_AdminEntity = "t_admin_entity"

const COLUMN_AdminEntity_Id = "id"
const COLUMN_AdminEntity_CreateTime = "createTime"
const COLUMN_AdminEntity_UpdateTime = "updateTime"
const COLUMN_AdminEntity_DeleteTime = "deleteTime"

const ATTRIBUTE_AdminEntity_Id = "Id"
const ATTRIBUTE_AdminEntity_CreateTime = "CreateTime"
const ATTRIBUTE_AdminEntity_UpdateTime = "UpdateTime"
const ATTRIBUTE_AdminEntity_DeleteTime = "DeleteTime"
