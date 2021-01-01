package services

import "ecommerce-backend/models"

type UserHandler interface {
	Register(string, string) (*models.User, error)
	Login(string, string) (*models.User, error)
}
