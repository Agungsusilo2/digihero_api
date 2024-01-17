package routes

import (
	"digihero/app/controller"
	"digihero/middleware"
	"digihero/model/applicant/repository"
	"digihero/model/applicant/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ctx *gin.Context

func Api(router *gin.Engine, db *gorm.DB, apiKey string) {
	digiHeroRepository := repository.NewApplicantRepository(db)
	digiHeroService := service.NewDigiHeroService(digiHeroRepository)
	digiHeroController := controller.NewDigiHeroController(digiHeroService, ctx)

	v1 := router.Group("/api/v1", middleware.AuthMiddleware(apiKey))
	{
		v1.GET("/digihero", digiHeroController.Index)
		v1.GET("/digihero/:id", digiHeroController.GetByID)
		v1.POST("/digihero", digiHeroController.Create)
		v1.PATCH("/digihero/:id", digiHeroController.Update)
		v1.DELETE("/digihero/:id", digiHeroController.Delete)
	}
}
