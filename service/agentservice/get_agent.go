package agentservice

import (
	"gameapp/pkg/richerror"
)

func (s Service) ExistsAgentID(agentID uint) (bool, error) {
	const op = "agentservice.ExistsAgentID"

	exists, err := s.repo.CheckExistsAgentID(agentID)
	if err != nil {
		return false, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if exists {
		return true, nil

	}
	return false, nil
}
