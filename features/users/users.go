package users

type Users struct {
	dbName string
}

func New(dbName string) Users {
	return Users{dbName}
}
