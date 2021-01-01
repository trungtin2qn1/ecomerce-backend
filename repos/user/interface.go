package repos

import "ecommerce-backend/models"

type DbReader interface {
	FindByID(interface{}) (*models.User, error)
}

type DbWriter interface {
	Create(*models.User) (*models.User, error)
}
