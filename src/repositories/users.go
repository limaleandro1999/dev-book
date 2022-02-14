package repositories

import (
	"dev-book/src/models"
	"errors"

	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func CreateUsersRepository(db *gorm.DB) *users {
	return &users{db}
}

func (userRepository users) Create(user models.User) (uint64, error) {
	result := userRepository.db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (userRepository users) GetUsers() ([]models.User, error) {
	var users []models.User
	result := userRepository.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (userRepository users) GetUser(userId string) (models.User, error) {
	var user models.User
	result := userRepository.db.Find(&user, userId)

	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (userRepository users) Update(userId string, user models.User, userData interface{}) (models.User, error) {
	result := userRepository.db.Model(&user).Updates(userData)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (userRepository users) Delete(userId string) error {
	result := userRepository.db.Delete(&models.User{}, userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
