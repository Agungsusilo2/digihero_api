package domain

type DigiHero struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Fullname       string `json:"full_name" gorm:"type:varchar(255);not null"`
	Name           string `json:"name" gorm:"type:varchar(255);not null"`
	Gmail          string `json:"gmail" gorm:"type:varchar(255);not null"`
	Telepon        int    `json:"telepon" gorm:"type:int;not null"`
	Age            int    `json:"age" gorm:"type:int;not null"`
	Job            string `json:"job" gorm:"type:varchar(255);not null"`
	RegionalOrigin string `json:"regional_origin" gorm:"type:varchar(255);not null"`
	Skill          string `json:"skill" gorm:"type:text"`
	Username       string `json:"username" gorm:"type:varchar(255);not null"`
	Password       string `json:"password" gorm:"type:varchar(255);not null"`
}
