package services

import (
	"errors"

	"github.com/Kachyr/SpyCatAgency/internal/constants"
	"github.com/Kachyr/SpyCatAgency/internal/store"
	"github.com/Kachyr/SpyCatAgency/pkg/models"
)

type MissionServiceI interface {
	CreateMission(mission *models.Mission) error
	GetMission(id uint) (*models.Mission, error)
	UpdateMission(mission *models.Mission) error
	DeleteMission(id uint) error
	CompleteMission(id uint) error
	ListMissions() ([]models.Mission, error)
	AddTarget(missionId uint, target *models.Target) error
	DeleteTarget(targetId uint) error
	UpdateTargetNotes(missionId uint, notes *models.TargetNotesJSON) error
	GetTarget(missionId uint) (*models.Target, error)
	CompleteTarget(missionId uint) error
}

type MissionService struct {
	MissionStore store.MissionStoreI
	CatStore     store.MissionStoreI
}

func NewMissionService(missionStore store.MissionStoreI) *MissionService {
	return &MissionService{MissionStore: missionStore}
}

func (s *MissionService) CreateMission(mission *models.Mission) error {
	if len(mission.Targets) > constants.MAX_MISSION_TARGETS {
		return errors.New("cannot create mission with more than 3 targets")
	}

	mission.Complete = false

	return s.MissionStore.CreateMission(mission)
}

func (s *MissionService) GetMission(id uint) (*models.Mission, error) {
	return s.MissionStore.GetMission(id)
}

func (s *MissionService) UpdateMission(mission *models.Mission) error {
	return s.MissionStore.UpdateMission(mission)
}

func (s *MissionService) DeleteMission(id uint) error {
	mission, err := s.GetMission(id)
	if err != nil {
		return err
	}

	if mission.CatID != nil {
		return errors.New("mission is assigned to a cat and cannot be deleted")
	}

	return s.MissionStore.DeleteMission(id)
}

func (s *MissionService) ListMissions() ([]models.Mission, error) {
	return s.MissionStore.ListMissions()
}

func (s *MissionService) AddTarget(missionId uint, target *models.Target) error {
	return s.MissionStore.AddTarget(missionId, target)
}

func (s *MissionService) DeleteTarget(targetId uint) error {
	return s.MissionStore.DeleteTarget(targetId)
}
func (s *MissionService) GetTarget(targetId uint) (*models.Target, error) {
	return s.MissionStore.GetTarget(targetId)
}

func (s *MissionService) UpdateTargetNotes(targetId uint, newNotes *models.TargetNotesJSON) error {
	return s.MissionStore.UpdateTargetNotes(targetId, newNotes)
}

func (s *MissionService) CompleteMission(missionID uint) error {
	return s.MissionStore.CompleteMission(missionID)
}

func (s *MissionService) CompleteTarget(targetId uint) error {
	return s.MissionStore.CompleteTarget(targetId)
}
