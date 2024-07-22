package models

import (
	"errors"

	"gorm.io/gorm"
)

type Target struct {
	gorm.Model
	MissionID uint   `gorm:"not null"`
	Name      string `gorm:"not null"`
	Country   string `gorm:"not null"`
	Notes     string
	Complete  bool `gorm:"not null"`
}

func (t *Target) BeforeDelete(tx *gorm.DB) (err error) {
	var target Target
	if err := tx.First(&target, t.ID).Error; err != nil {
		return err
	}
	if target.Complete {
		return errors.New("cant delete completed target")
	}
	return nil
}

type TargetJSON struct {
	ID        uint   `json:"id"`
	MissionID uint   `json:"missionId"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Notes     string `json:"notes"`
	Complete  bool   `json:"complete"`
}
type TargetNotesJSON struct {
	Notes string `json:"notes" binding:"required"`
}

func ToTargetJSON(target Target) TargetJSON {
	return TargetJSON{
		ID:        target.ID,
		MissionID: target.MissionID,
		Name:      target.Name,
		Country:   target.Country,
		Notes:     target.Notes,
		Complete:  target.Complete,
	}
}

func FromTargetJSON(targetJSON TargetJSON) Target {
	return Target{
		Name:    targetJSON.Name,
		Country: targetJSON.Country,
		Notes:   targetJSON.Notes,
	}
}

func ToTargetJSONList(targets []Target) []TargetJSON {
	targetJSONList := make([]TargetJSON, len(targets))
	for i, target := range targets {
		targetJSONList[i] = ToTargetJSON(target)
	}
	return targetJSONList
}

func FromTargetJSONList(targetJSONs []TargetJSON) []Target {
	targetList := make([]Target, len(targetJSONs))
	for i, targetJSON := range targetJSONs {
		targetList[i] = FromTargetJSON(targetJSON)
	}
	return targetList
}
