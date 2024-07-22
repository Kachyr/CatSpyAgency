package models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name              string   `gorm:"not null"`
	YearsOfExperience int      `gorm:"not null"`
	Breed             string   `gorm:"not null"`
	Salary            float64  `gorm:"not null"`
	Mission           *Mission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CatJSON struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name" binding:"required"`
	YearsOfExperience int     `json:"yearsOfExperience" binding:"required"`
	Breed             string  `json:"breed" binding:"required"`
	Salary            float64 `json:"salary" binding:"required"`
}

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateSalaryJSON struct {
	Salary float64 `json:"salary" binding:"required"`
}

type AssignMissionJSON struct {
	MissionId uint `json:"missionId" binding:"required"`
}

func ToCatJSON(cat Cat) CatJSON {
	return CatJSON{
		ID:                cat.ID,
		Name:              cat.Name,
		YearsOfExperience: cat.YearsOfExperience,
		Breed:             cat.Breed,
		Salary:            cat.Salary,
	}
}

func FromCatJSON(catJSON CatJSON) Cat {
	return Cat{
		Name:              catJSON.Name,
		YearsOfExperience: catJSON.YearsOfExperience,
		Breed:             catJSON.Breed,
		Salary:            catJSON.Salary,
	}
}

func ToCatsJsonArray(cats []Cat) []CatJSON {
	catsJson := []CatJSON{}
	for _, cat := range cats {
		catsJson = append(catsJson, ToCatJSON(cat))
	}
	return catsJson
}
