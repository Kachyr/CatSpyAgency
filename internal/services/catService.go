package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Kachyr/SpyCatAgency/internal/store"
	"github.com/Kachyr/SpyCatAgency/pkg/models"
	"github.com/go-resty/resty/v2"
)

const theCatAPIUrl = "https://api.thecatapi.com/v1/breeds"

type CatServiceI interface {
	AddCat(cat *models.Cat) error
	DeleteCat(id uint) error
	GetCat(id uint) (*models.Cat, error)
	ListCats() ([]models.Cat, error)
	UpdateCat(catId uint, cat *models.Cat) error
	UpdateCatSalary(catId uint, salary float64) error
	AssignMission(catId uint, missionId uint) error
}

type CatService struct {
	CatStore     store.CatStoreI
	MissionStore store.MissionStoreI
}

func NewCatService(catStore store.CatStoreI, missionStore store.MissionStoreI) *CatService {
	return &CatService{CatStore: catStore, MissionStore: missionStore}
}

func (s *CatService) AddCat(cat *models.Cat) error {
	if err := s.validateCatBreed(cat.Breed); err != nil {
		return err
	}
	return s.CatStore.AddCat(cat)
}

func (s *CatService) GetCat(id uint) (*models.Cat, error) {
	return s.CatStore.GetCat(id)
}

func (s *CatService) UpdateCat(catId uint, cat *models.Cat) error {
	if err := s.validateCatBreed(cat.Breed); err != nil && cat.Breed != "" {
		return err
	}
	return s.CatStore.UpdateCat(catId, cat)
}
func (s *CatService) UpdateCatSalary(catId uint, newSalary float64) error {
	return s.CatStore.UpdateCatSalary(catId, newSalary)
}

func (s *CatService) DeleteCat(id uint) error {
	return s.CatStore.DeleteCat(id)
}

func (s *CatService) ListCats() ([]models.Cat, error) {
	return s.CatStore.ListCats()
}
func (s *CatService) AssignMission(catId uint, missionId uint) error {
	mission, err := s.MissionStore.GetMission(missionId)
	if err != nil {
		return err
	}

	if mission.Complete {
		return errors.New("cant assign cat on completed mission")
	}

	cat, err := s.CatStore.GetCat(catId)
	if err != nil {
		return err
	}
	if cat.Mission != nil {
		return errors.New("cat already have assigned mission")
	}

	return s.CatStore.AssignMission(catId, mission)
}

func (s *CatService) validateCatBreed(breedName string) error {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(theCatAPIUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New("failed to fetch breeds from TheCatAPI")
	}

	var breeds []models.Breed
	if err := json.Unmarshal(resp.Body(), &breeds); err != nil {
		return err
	}

	for _, breed := range breeds {
		if breed.Name == breedName {
			return nil
		}
	}

	return errors.New("invalid breed")
}
