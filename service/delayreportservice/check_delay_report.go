package delayreportservice

import (
	"gameapp/pkg/richerror"
)

func (s Service) CheckQueueAgentID(agentID uint) (bool, error) {
	const op = "delayreportservice.CheckQueueAgentID"

	exists, err := s.repo.CheckAgentBusyInQueue(agentID)
	if err != nil {
		return false, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if exists {
		return true, nil
	}
	return false, nil
}
