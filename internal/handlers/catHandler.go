package handlers

import (
	"net/http"
	"strconv"

	"github.com/Kachyr/SpyCatAgency/internal/services"
	"github.com/Kachyr/SpyCatAgency/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CatHandlerI interface {
	GetCat(c *gin.Context)
	AddCat(c *gin.Context)
	UpdateCat(c *gin.Context)
	UpdateCatSalary(c *gin.Context)
	DeleteCat(c *gin.Context)
	GetAllCats(c *gin.Context)
	AssignMission(c *gin.Context)
}

type CatHandler struct {
	CatService services.CatServiceI
}

func NewCatHandler(catService services.CatServiceI) *CatHandler {
	return &CatHandler{CatService: catService}
}

func (h *CatHandler) AddCat(c *gin.Context) {
	var body models.CatJSON
	if err := c.BindJSON(&body); err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	cat := models.FromCatJSON(body)

	if err := h.CatService.AddCat(&cat); err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CatHandler) GetCat(c *gin.Context) {
	catId := c.Param("id")
	if catId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(catId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}
	cat, err := h.CatService.GetCat(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ToCatJSON(*cat))
}

func (h *CatHandler) GetAllCats(c *gin.Context) {
	cats, err := h.CatService.ListCats()
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ToCatsJsonArray(cats))
}

func (h *CatHandler) UpdateCat(c *gin.Context) {
	catId := c.Param("id")
	var body models.CatJSON
	if err := c.BindJSON(&body); err != nil || catId == "" {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(catId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	cat := models.FromCatJSON(body)

	err = h.CatService.UpdateCat(uint(id), &cat)
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CatHandler) UpdateCatSalary(c *gin.Context) {
	catId := c.Param("id")
	var body models.UpdateSalaryJSON
	if err := c.BindJSON(&body); err != nil || catId == "" {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(catId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.CatService.UpdateCatSalary(uint(id), body.Salary)
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CatHandler) DeleteCat(c *gin.Context) {
	catId := c.Param("id")
	if catId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(catId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.CatService.DeleteCat(uint(id))
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *CatHandler) AssignMission(c *gin.Context) {
	catId := c.Param("id")
	var body models.AssignMissionJSON
	if err := c.BindJSON(&body); err != nil || catId == "" {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(catId, 10, 32)
	if err != nil {
		log.Info().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.CatService.AssignMission(uint(id), body.MissionId)
	if err != nil {
		log.Info().Err(err).Send()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
