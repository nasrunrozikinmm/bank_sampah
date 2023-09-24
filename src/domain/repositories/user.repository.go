package repositories

import (
	"fmt"

	models "github.com/iyunrozikin26/bank_sampah.git/src/domain/models"
	"gorm.io/gorm"
)

// kontrak methods
type UserRepository interface {
	FindAll() []models.User
	FindOne(id int) models.User
	Save(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(user models.User) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (ur *UserRepositoryImpl) FindAll() []models.User {
	var users []models.User

	_ = ur.db.Find(&users) // menggunakan pointer agar merujuk ke memory yang sama
	return users

}
func (ur *UserRepositoryImpl) FindOne(id int) models.User {
	var user models.User
	_ = ur.db.First(&user, id)
	return user
}
func (ur *UserRepositoryImpl) Save(user models.User) (*models.User, error) {
	fmt.Println(user, "userrrrrrrrrrrrrrrrrrrrrrrrrrrr")
	result := ur.db.Create(&user)
	if result != nil {
		return nil, result.Error
	}
	
	return &user, nil
}
func (ur *UserRepositoryImpl) Update(user models.User) (*models.User, error) {
	result := ur.db.Save(&user)
	if result != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (ur *UserRepositoryImpl) Delete(user models.User) (*models.User, error) {
	result := ur.db.Delete(&user)
	if result != nil {
		return nil, result.Error
	}
	return &user, nil
}