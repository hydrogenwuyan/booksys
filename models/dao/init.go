package dao

const (
	database = "booksys"
)

func Init() {
	AdminDaoEntity = NewAdminDao(database)
	BookDaoEntity = NewBookDao(database)
}
