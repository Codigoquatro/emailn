package campanha

type Repository interface {
	Save(campanha *Campanha) error
}
