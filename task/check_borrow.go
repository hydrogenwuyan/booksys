package task

import (
	"github.com/astaxie/beego/toolbox"
	"project/booksys/common"
	"project/booksys/models/dao"
	"project/booksys/models/entity"
)

// 30s执行一次
func init() {
	toolbox.AddTask("check_borrow", toolbox.NewTask("check_borrow", "0/30 * * * * *", toolbox.TaskFunc(checkBorrow)))
}

func checkBorrow() (err error) {
	common.LogFuncDebug("checkBorrow run")

	idList, err := dao.BookDaoEntity.FetchExpireStuIdList()
	if err != nil {
		return
	}

	for _, id := range idList {
		var stuEntity *entity.StudentEntity
		stuEntity, err = dao.StudentDaoEntity.Info(id)
		if err != nil {
			return
		}

		if stuEntity.IsBlack == dao.StudentIsBlack {
			continue
		}

		stuEntity.IsBlack = dao.StudentIsBlack
		err = dao.StudentDaoEntity.Update(stuEntity)
		if err != nil {
			return
		}
	}
	return
}
