package handlers

import (
	"net/http"
	"strconv"

	"github.com/Kachyr/SpyCatAgency/internal/services"
	"github.com/Kachyr/SpyCatAgency/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type MissionHandlerI interface {
	CreateMission(c *gin.Context)
	DeleteMission(c *gin.Context)
	AddTarget(c *gin.Context)
	GetMission(c *gin.Context)
	UpdateTargetNotes(c *gin.Context)
	ListMissions(c *gin.Context)
	CompleteMission(c *gin.Context)
	CompleteTarget(c *gin.Context)
	GetTarget(c *gin.Context)
	DeleteTarget(c *gin.Context)
}

type MissionHandler struct {
	MissionService services.MissionServiceI
}

func NewMissionHandler(missionService services.MissionServiceI) *MissionHandler {
	return &MissionHandler{MissionService: missionService}
}

func (h *MissionHandler) CreateMission(c *gin.Context) {
	var body models.CreateMissionJSON
	if err := c.BindJSON(&body); err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	mission := models.FromMissionJSON(body)

	if err := h.MissionService.CreateMission(&mission); err != nil {
		log.Info().Err(err).Msg("Cant add cat")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *MissionHandler) DeleteMission(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(missionId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}
	err = h.MissionService.DeleteMission(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *MissionHandler) AddTarget(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(missionId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}
	var body models.TargetJSON
	if err := c.BindJSON(&body); err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	target := models.FromTargetJSON(body)
	err = h.MissionService.AddTarget(uint(id), &target)
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *MissionHandler) GetMission(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(missionId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	mission, err := h.MissionService.GetMission(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ToMissionJSON(*mission))
}

func (h *MissionHandler) UpdateTargetNotes(c *gin.Context) {
	targetId := c.Param("id")
	if targetId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(targetId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	var body models.TargetNotesJSON
	if err := c.BindJSON(&body); err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	err = h.MissionService.UpdateTargetNotes(uint(id), &body)
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *MissionHandler) ListMissions(c *gin.Context) {
	missions, err := h.MissionService.ListMissions()
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ToMissionsJSON(missions))
}

func (h *MissionHandler) CompleteMission(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(missionId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.MissionService.CompleteMission(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *MissionHandler) CompleteTarget(c *gin.Context) {
	targetId := c.Param("id")
	if targetId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(targetId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.MissionService.CompleteTarget(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *MissionHandler) GetTarget(c *gin.Context) {
	targetId := c.Param("id")
	if targetId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(targetId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	target, err := h.MissionService.GetTarget(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ToTargetJSON(*target))
}

func (h *MissionHandler) DeleteTarget(c *gin.Context) {
	targetId := c.Param("id")
	if targetId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(targetId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}
	err = h.MissionService.DeleteTarget(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
