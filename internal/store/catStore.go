package store

import (
	"github.com/Kachyr/SpyCatAgency/pkg/models"
	"gorm.io/gorm"
)

type CatStoreI interface {
	AddCat(cat *models.Cat) error
	DeleteCat(id uint) error
	GetCat(id uint) (*models.Cat, error)
	ListCats() ([]models.Cat, error)
	UpdateCat(catId uint, cat *models.Cat) error
	UpdateCatSalary(catId uint, salary float64) error
	AssignMission(catId uint, missionId *models.Mission) error
}

type CatStore struct {
	db *gorm.DB
}

func NewCatStore(db *gorm.DB) *CatStore {
	return &CatStore{db: db}
}

func (s *CatStore) AddCat(cat *models.Cat) error {

	return s.db.Create(&cat).Error
}

func (s *CatStore) GetCat(id uint) (*models.Cat, error) {
	var cat models.Cat
	if err := s.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (s *CatStore) UpdateCat(catId uint, cat *models.Cat) error {
	var c models.Cat
	if err := s.db.First(&c, "id = ?", catId).Error; err != nil {
		return err
	}
	c.Breed = cat.Breed
	c.Name = cat.Name
	c.Salary = cat.Salary
	c.YearsOfExperience = cat.YearsOfExperience

	return s.db.Save(c).Error
}

func (s *CatStore) UpdateCatSalary(catId uint, salary float64) error {
	return s.db.Model(&models.Cat{}).Where("id = ?", catId).Update("salary", salary).Error
}

func (s *CatStore) AssignMission(catId uint, mission *models.Mission) error {
	var cat models.Cat
	if err := s.db.First(&cat, "id = ?", catId).Error; err != nil {
		return err
	}
	cat.Mission = mission

	return s.db.Save(cat).Error
}

func (s *CatStore) DeleteCat(id uint) error {
	return s.db.Delete(&models.Cat{}, id).Error
}

func (s *CatStore) ListCats() ([]models.Cat, error) {
	var cats []models.Cat
	if err := s.db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}
