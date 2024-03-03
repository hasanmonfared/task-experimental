package agentservice

type Repository interface {
	CheckExistsAgentID(agentID uint) (bool, error)
}
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
