package delayreportvalidator

type Repository interface {
	CheckQueueAgentID(agentID uint) (bool, error)
}
type Order interface {
	IsOrderExceedingTheTimeDelivery(orderID uint) (bool, error)
}
type Agent interface {
	ExistsAgentID(agentID uint) (bool, error)
}
type Validator struct {
	repo  Repository
	order Order
	agent Agent
}

func New(repo Repository, agent Agent, order Order) Validator {
	return Validator{repo: repo, agent: agent, order: order}
}
