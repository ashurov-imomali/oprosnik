package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"main/models"
	"strconv"
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
	user.Role = "user"
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

func (s *Service) UpdateQuestions(quest []models.Questions) ([]models.Questions, error) {
	err := s.Db.Model(&models.Questions{}).Save(quest).Error
	if err != nil {
		return nil, err
	}
	return quest, nil
}

func (s *Service) DeleteQuestions(id int) error {
	err := s.Db.Where("question_id=?", id).Delete(&models.Answers{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return s.Db.Where("id = ?", id).Delete(&models.Questions{}).Error
}

func (s *Service) AddAnswers(mp map[string]interface{}) ([]models.Answers, error) {
	strId, ok := mp["user_id"].(string)
	if !ok {
		return nil, errors.New("USERID !!!")
	}
	log.Println(mp)
	delete(mp, "user_id")
	var ansArray []models.Answers
	userId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return nil, err
	}
	for key, value := range mp {
		var ans models.Answers
		questionId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, err
		}
		ans.UserId = userId
		ans.QuestionId = questionId
		ans.Variant = value.(string)
		ansArray = append(ansArray, ans)
	}
	if err := s.Db.Create(ansArray).Error; err != nil {
		return nil, err
	}
	return ansArray, nil
}

type data struct {
	Question string `gorm:"column:name"`
	Variant  string `gorm:"column:variant"`
}

type just struct {
	FullName string `gorm:"column:full_name"`
	Id       int64  `json:"id" gorm:"column:id"`
}

func (s *Service) GetAnswers() ([]models.UserAns, error) {
	var anss []models.UserAns
	var j []just
	err := s.Db.Model(&models.User{}).Where("role=?", "user").
		Select("full_name, id").Scan(&j).Error
	if err != nil {
		return nil, err
	}
	log.Println(j)
	for _, j2 := range j {
		var data []data
		var items models.UserAns
		err = s.Db.Table("answers a").Select("q.name, a.variant").
			Joins("join questions q on q.id = a.question_id").Where("user_id = ?", j2.Id).
			Scan(&data).Error
		if err != nil {
			return nil, err
		}
		for _, datum := range data {
			items.Questions = append(items.Questions, datum.Question)
			items.Variants = append(items.Variants, datum.Variant)
		}
		items.FullName = j2.FullName
		anss = append(anss, items)
	}
	return anss, nil
}
