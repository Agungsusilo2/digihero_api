package service

import (
	"digihero/model/applicant/domain"
	"digihero/model/applicant/dto"
	"digihero/model/applicant/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type DigiHeroServiceImp struct {
	digiHeroRepository repository.DigiHeroRepository
}

func NewDigiHeroService(digiHeroRepository repository.DigiHeroRepository) DigiHeroService {
	return &DigiHeroServiceImp{digiHeroRepository: digiHeroRepository}
}
func (a *DigiHeroServiceImp) GetAll() []domain.DigiHero {
	return a.digiHeroRepository.FindAll()
}

func (a *DigiHeroServiceImp) GetByID(id string) domain.DigiHero {
	return a.digiHeroRepository.FindById(id)
}

func (a *DigiHeroServiceImp) Create(ctx *gin.Context) (*domain.DigiHero, error) {
	var input dto.CreateDigiHeroReq

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}
	uniqueID := uuid.New()
	unique := (uniqueID).String()

	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	if err != nil {
		return nil, err
	}
	password := string(bytes)
	digiHero := domain.DigiHero{
		ID:             unique,
		Fullname:       input.Fullname,
		Name:           "DigiHero",
		Gmail:          input.Gmail,
		Telepon:        input.Telepon,
		Age:            input.Age,
		Job:            input.Job,
		RegionalOrigin: input.RegionalOrigin,
		Skill:          input.Skill,
		Username:       input.Username,
		Password:       password,
	}
	result, err := a.digiHeroRepository.Save(digiHero)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *DigiHeroServiceImp) Update(ctx *gin.Context) (*domain.DigiHero, error) {
	var update dto.UpdateDigiHero

	if err := ctx.ShouldBindJSON(&update); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(update)

	if err != nil {
		return nil, err
	}
	id := ctx.Param("id")
	newPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(update.Password), 14)
	if err != nil {
		return nil, err
	}
	password := string(newPasswordBytes)
	digiHero := domain.DigiHero{
		ID:             id,
		Fullname:       update.Fullname,
		Name:           "DigiHero",
		Gmail:          update.Gmail,
		Telepon:        update.Telepon,
		Age:            update.Age,
		Job:            update.Job,
		RegionalOrigin: update.RegionalOrigin,
		Skill:          update.Skill,
		Username:       update.Username,
		Password:       password,
	}

	result, err := a.digiHeroRepository.Update(digiHero)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *DigiHeroServiceImp) Delete(ctx *gin.Context) (*domain.DigiHero, error) {
	id := ctx.Param("id")
	digiHero := domain.DigiHero{
		ID: id,
	}

	result, err := a.digiHeroRepository.Delete(digiHero)

	if err != nil {
		return nil, err
	}

	return result, nil
}
