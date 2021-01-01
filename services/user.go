package services

import (
	"ecommerce-backend/database"
	"ecommerce-backend/models"
	userRepo "ecommerce-backend/repos/user"
)

const (
	postgreSQL = "postgres"
)

type User struct {
	dbReader userRepo.DbReader
	dbWriter userRepo.DbWriter
}

func NewUser() *User {
	return &User{}
}

func (user *User) Register(email, password string) (*models.User, error) {
	//TODO: Verify input (email and password) if it is valid

	hashPwd := password
	userModel := models.NewUser()
	userModel.Email = email
	userModel.Password = hashPwd
	err := user.setDbWriter(postgreSQL)
	if err != nil {
		return nil, err
	}
	return user.dbWriter.Create(userModel)
}

func (user *User) Login(email, password string) (*models.User, error) {
	return nil, nil
}

func (user *User) setDbWriter(Type string) error {
	var dbWriter userRepo.DbWriter
	switch Type {
	case postgreSQL:
		dbWriter = userRepo.NewPostgreSQLWithValue(database.SqlDB())
	default:
		panic("Not implemented yet")
	}
	user.dbWriter = dbWriter
	return nil
}
