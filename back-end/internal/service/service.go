package service

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"main/models"
)

type Service struct {
	Db *gorm.DB
}

func InitialSrv(db *gorm.DB) *Service {
	return &Service{Db: db}
}

func (s *Service) CheckNewUser(user *models.User) (*models.User, error) {
	var usr models.User
	err := s.Db.Where("username = ?", user.Username).First(&usr).Error
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (s *Service) CreateUser(user *models.User) (*models.User, error) {
	data := []byte(user.Password)
	hash := md5.Sum(data)
	user.Password = fmt.Sprintf("%x", hash)
	err := s.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) VerifyPass(user *models.User, passHash string) bool {
	hash := md5.Sum([]byte(user.Password))
	return fmt.Sprintf("%x", hash) == passHash
}

func (s *Service) GetQuestions() ([]models.Questions, error) {
	var quests []models.Questions
	err := s.Db.Model(&models.Questions{}).Find(&quests).Error
	if err != nil {
		return nil, err
	}
	return quests, nil
}

func (s *Service) UpdateQuestions(quest *models.Questions) (*models.Questions, error) {
	err := s.Db.Model(&models.Questions{}).Save(quest).Error
	if err != nil {
		return nil, err
	}
	return quest, nil
}

func (s *Service) DeleteQuestions(id int) error {
	return s.Db.Where("id = ?", id).Delete(&models.Questions{}).Error
}