package repository

import (
	"digihero/model/applicant/domain"
	"gorm.io/gorm"
)

type DigiHeroRepositoryImp struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) DigiHeroRepository {
	return &DigiHeroRepositoryImp{db}
}

func (d *DigiHeroRepositoryImp) FindAll() []domain.DigiHero {
	var applicants []domain.DigiHero

	_ = d.db.Find(&applicants)

	return applicants
}

func (d *DigiHeroRepositoryImp) FindById(id string) domain.DigiHero {
	var applicant domain.DigiHero

	_ = d.db.First(&applicant, "id = ?", id)

	return applicant
}

func (d *DigiHeroRepositoryImp) Save(applicant domain.DigiHero) (*domain.DigiHero, error) {
	result := d.db.Create(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (d *DigiHeroRepositoryImp) Update(applicant domain.DigiHero) (*domain.DigiHero, error) {
	result := d.db.Save(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (d *DigiHeroRepositoryImp) Delete(applicant domain.DigiHero) (*domain.DigiHero, error) {
	result := d.db.Delete(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}
