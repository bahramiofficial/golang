package companyname

import (
	"gorm.io/gorm"
	models "pejamn.ir/pers/models"
)

type Repo struct {
	db *gorm.DB
}

//facktory patten
func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (m *Repo) Create(input *models.CompanyName) error {
	if err := m.db.Debug().Create(input).Error; err != nil {
		return err
	}
	return nil
}

func (m *Repo) Read(id *models.PK) (*models.CompanyName, error) {
	var (
		//cn = new(models.CompanyName)
		//cn =  models.CompanyName and first &cn
		cn = &models.CompanyName{} //best prastic
	)
	//err := m.db.where('id = ?' ,id).First(cn).Error
	err := m.db.First(cn, id).Error
	if err != nil {
		return nil, err
	}
	return cn, nil
}
