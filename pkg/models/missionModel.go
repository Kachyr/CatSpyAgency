package models

import (
	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	CatID    *uint
	Targets  []Target `gorm:"foreignKey:MissionID"`
	Complete bool     `gorm:"not null"`
}

type MissionJSON struct {
	ID       uint         `json:"id"`
	CatID    *uint        `json:"catId"`
	Targets  []TargetJSON `json:"targets"`
	Complete bool         `json:"complete"`
}

type CreateMissionJSON struct {
	CatID   *uint        `json:"catId"`
	Targets []TargetJSON `json:"targets"`
}

func ToMissionJSON(mission Mission) MissionJSON {
	return MissionJSON{
		ID:       mission.ID,
		CatID:    mission.CatID,
		Targets:  ToTargetJSONList(mission.Targets),
		Complete: mission.Complete,
	}
}

func FromMissionJSON(missionJSON CreateMissionJSON) Mission {
	return Mission{
		CatID:   missionJSON.CatID,
		Targets: FromTargetJSONList(missionJSON.Targets),
	}
}

func ToMissionsJSON(missions []Mission) []MissionJSON {
	var missionsJSON []MissionJSON

	for _, mission := range missions {
		missionsJSON = append(missionsJSON, ToMissionJSON(mission))
	}

	return missionsJSON
}
