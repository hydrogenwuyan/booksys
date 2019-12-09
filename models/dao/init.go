package dao

const (
	database = "booksys"
)

func Init() {
	AdminDaoEntity.Init(database)
}
