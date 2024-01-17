package repository

import "digihero/model/applicant/domain"

type DigiHeroRepository interface {
	FindAll() []domain.DigiHero
	FindById(id string) domain.DigiHero
	Save(applicant domain.DigiHero) (*domain.DigiHero, error)
	Update(applicant domain.DigiHero) (*domain.DigiHero, error)
	Delete(applicant domain.DigiHero) (*domain.DigiHero, error)
}
