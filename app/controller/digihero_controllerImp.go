package controller

import (
	"digihero/model/applicant/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DigiHeroControllerImp struct {
	digiHeroService service.DigiHeroService
	ctx             *gin.Context
}

func NewDigiHeroController(digiHeroService service.DigiHeroService, ctx *gin.Context) DigiHeroController {
	return &DigiHeroControllerImp{digiHeroService: digiHeroService, ctx: ctx}
}

func (a *DigiHeroControllerImp) Index(ctx *gin.Context) {
	data := a.digiHeroService.GetAll()

	ctx.JSON(http.StatusOK, data)
}

func (a *DigiHeroControllerImp) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data := a.digiHeroService.GetByID(id)

	ctx.JSON(http.StatusOK, data)
}

func (a *DigiHeroControllerImp) Create(ctx *gin.Context) {
	data, err := a.digiHeroService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
			"code":   http.StatusInternalServerError,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (a *DigiHeroControllerImp) Delete(ctx *gin.Context) {
	data, err := a.digiHeroService.Delete(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (a *DigiHeroControllerImp) Update(ctx *gin.Context) {
	data, err := a.digiHeroService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}
