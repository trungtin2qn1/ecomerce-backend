package repos

import (
	"ecommerce-backend/models"

	"gorm.io/gorm"
)

type Mongo struct {
}

type PostgreSQL struct {
	db *gorm.DB
}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{}
}

func NewPostgreSQLWithValue(db *gorm.DB) *PostgreSQL {
	return &PostgreSQL{
		db: db,
	}
}

type MySQL struct {
	db *gorm.DB
}

func (p *PostgreSQL) FindByID(id interface{}) (*models.User, error) {
	return nil, nil
}

func (p *PostgreSQL) Create(user *models.User) (*models.User, error) {
	res := user
	tx := p.db.Begin()
	defer func() {
		if p.db.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	tx = tx.Create(res)
	return res, tx.Error
}

/*func (m *MySQL) FindByID(id string) (*models.User, error) {*/
//return nil, nil
//}

//func (mongo *Mongo) FindByID(id string) (*models.User, error) {
//return nil, nil
/*}*/
