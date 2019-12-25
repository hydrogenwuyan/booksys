package dao

const (
	Database = "booksys"
)

func Init() {
	AdminDaoEntity = NewAdminDao(Database)
	BookDaoEntity = NewBookDao(Database)
	StudentDaoEntity = NewStudentDao(Database)
}
