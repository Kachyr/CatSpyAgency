package initializers

import (
	"github.com/Kachyr/SpyCatAgency/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	CatHandler     handlers.CatHandlerI
	MissionHandler handlers.MissionHandlerI
}

func NewRouter(catHandler handlers.CatHandlerI, missionHandler handlers.MissionHandlerI) *Router {

	return &Router{
		CatHandler:     catHandler,
		MissionHandler: missionHandler,
	}
}

func (r *Router) SetupAPIs(e *gin.Engine) {
	e.POST("/cats", r.CatHandler.AddCat)
	e.GET("/cats", r.CatHandler.GetAllCats)
	e.GET("/cats/:id", r.CatHandler.GetCat)
	e.PUT("/cats/:id", r.CatHandler.UpdateCat)
	e.PUT("/cats/:id/salary", r.CatHandler.UpdateCatSalary)
	e.DELETE("/cats/:id", r.CatHandler.DeleteCat)
	e.PUT("/cats/:id/assign-mission", r.CatHandler.AssignMission)

	e.POST("/mission", r.MissionHandler.CreateMission)
	e.DELETE("/mission/:id", r.MissionHandler.DeleteMission)
	e.GET("/mission/:id", r.MissionHandler.GetMission)
	e.GET("/mission", r.MissionHandler.ListMissions)
	e.PUT("/mission/:id/add-target", r.MissionHandler.AddTarget)
	e.PUT("/target/:id/notes", r.MissionHandler.UpdateTargetNotes)
	e.GET("/target/:id", r.MissionHandler.GetTarget)
	e.PUT("/mission/:id/complete", r.MissionHandler.CompleteMission)
	e.PUT("/target/:id/complete", r.MissionHandler.CompleteTarget)
	e.DELETE("/target/:id", r.MissionHandler.DeleteTarget)
}
