package service

import (
	"digihero/model/applicant/domain"
	"github.com/gin-gonic/gin"
)

type DigiHeroService interface {
	GetAll() []domain.DigiHero
	GetByID(id string) domain.DigiHero
	Create(ctx *gin.Context) (*domain.DigiHero, error)
	Update(ctx *gin.Context) (*domain.DigiHero, error)
	Delete(ctx *gin.Context) (*domain.DigiHero, error)
}
