package companyname

type Logic struct {
	repo *Repo
}

func NewLogic(repo *Repo) *Logic {
	return &Logic{
		repo: repo,
	}
}

func (m *Logic) Create(input *models.CompanyName) error {
	return m.repo.Create(input)
}

func (m *Logic) read(id *models.PK) (*models.CompanyName, error) {
	return m.repo.Read(id)
}
