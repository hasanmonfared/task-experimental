package mysqlagent

import (
	"database/sql"
	"gameapp/adapter/mysql"
	"gameapp/entity/agententity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"time"
)

func (d DB) CheckExistsAgentID(agentID uint) (bool, error) {
	const op = "mysqldelayreport.CheckExistsAgentID"

	row := d.adapter.Conn().QueryRow(`select * from agents where id= ?`, agentID)
	_, err := scanDelayReport(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return true, nil
}

func scanDelayReport(scanner mysql.Scanner) (agententity.Agent, error) {
	var agent agententity.Agent
	var createdAt time.Time
	err := scanner.Scan(&agent.ID, &agent.Firstname, &agent.Lastname, &createdAt)
	agent.CreatedAt = createdAt
	return agent, err
}
