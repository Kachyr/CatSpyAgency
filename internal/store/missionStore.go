package store

import (
	"errors"

	"github.com/Kachyr/SpyCatAgency/internal/constants"
	"github.com/Kachyr/SpyCatAgency/pkg/models"
	"gorm.io/gorm"
)

type MissionStoreI interface {
	CreateMission(mission *models.Mission) error
	GetMission(id uint) (*models.Mission, error)
	UpdateMission(mission *models.Mission) error
	DeleteMission(id uint) error
	CompleteMission(id uint) error
	ListMissions() ([]models.Mission, error)
	AddTarget(missionId uint, target *models.Target) error
	DeleteTarget(targetId uint) error
	UpdateTargetNotes(targetId uint, notes *models.TargetNotesJSON) error
	GetTarget(targetId uint) (*models.Target, error)
	CompleteTarget(targetId uint) error
}

type MissionStore struct {
	db *gorm.DB
}

func NewMissionStore(db *gorm.DB) *MissionStore {
	return &MissionStore{db: db}
}

func (s *MissionStore) CreateMission(mission *models.Mission) error {
	return s.db.Create(mission).Error
}

func (s *MissionStore) GetMission(id uint) (*models.Mission, error) {
	var mission models.Mission
	if err := s.db.Preload("Targets").First(&mission, id).Error; err != nil {
		return nil, err
	}
	return &mission, nil
}

func (s *MissionStore) UpdateMission(mission *models.Mission) error {
	return s.db.Save(mission).Error
}

func (s *MissionStore) DeleteMission(id uint) error {
	return s.db.Delete(&models.Mission{}, id).Error
}

func (s *MissionStore) ListMissions() ([]models.Mission, error) {
	var missions []models.Mission
	if err := s.db.Preload("Targets").Find(&missions).Error; err != nil {
		return nil, err
	}
	return missions, nil
}

func (s *MissionStore) AddTarget(missionId uint, target *models.Target) error {
	mission, err := s.GetMission(missionId)
	if err != nil {
		return err
	}

	if mission.Complete {
		return errors.New("cannot add targets to a completed mission")
	}
	if len(mission.Targets) >= constants.MAX_MISSION_TARGETS {
		return errors.New("mission already has maximum amount of targets")
	}
	return s.db.Model(&mission).Association("Targets").Append(target)
}

func (s *MissionStore) CompleteMission(missionId uint) error {
	tx := s.db.Begin()

	var mission models.Mission
	if err := tx.Preload("Targets").First(&mission, missionId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if mission.Complete {
		return errors.New("cannot update a target of a completed mission")
	}

	for i := range mission.Targets {
		mission.Targets[i].Complete = true
		if err := tx.Save(&mission.Targets[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	mission.Complete = true
	if err := tx.Save(&mission).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *MissionStore) GetTarget(targetId uint) (*models.Target, error) {
	var target models.Target
	if err := s.db.First(&target, targetId).Error; err != nil {
		return nil, err
	}
	return &target, nil
}

func (s *MissionStore) DeleteTarget(targetId uint) error {
	target, err := s.GetTarget(targetId)
	if err != nil {
		return err
	}
	return s.db.Delete(&target).Error
}

func (s *MissionStore) UpdateTargetNotes(targetId uint, newNotes *models.TargetNotesJSON) error {
	target, err := s.GetTarget(targetId)

	if err != nil {
		return err
	}

	target.Notes = newNotes.Notes

	return s.db.Save(target).Error
}

func (s *MissionStore) CompleteTarget(targetId uint) error {
	target, err := s.GetTarget(targetId)
	if err != nil {
		return err
	}
	target.Complete = true
	return s.db.Save(target).Error
}
