package orderdelayvalidator

type Repository interface {
	IsOrderByStatusesExist(orderID uint) (bool, error)
}
type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
