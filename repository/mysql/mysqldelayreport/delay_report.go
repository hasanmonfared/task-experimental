package mysqldelayreport

import (
	"database/sql"
	"gameapp/adapter/mysql"
	"gameapp/entity/delayreportentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (d DB) InsertDelayReport(ctx context.Context, orderID uint, deliveryTime time.Time) error {
	const op = "mysqldelayreport.InsertDelayReport"
	var query string
	var args []interface{}

	if !deliveryTime.IsZero() {
		query = `INSERT INTO delay_reports (order_id, delivery_time) VALUES (?, ?)`
		args = []interface{}{orderID, deliveryTime}
	} else {
		query = `INSERT INTO delay_reports (order_id) VALUES (?)`
		args = []interface{}{orderID}
	}

	_, err := d.adapter.Conn().ExecContext(ctx, query, args...)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
func (d DB) HasPendingDelayReport(ctx context.Context, orderID uint) (bool, error) {
	const op = "mysqldelayreport.HasPendingDelayReport"

	row := d.adapter.Conn().QueryRowContext(ctx, `select * from delay_reports where order_id= ? AND delay_check = false`, orderID)
	_, err := scanDelayReport(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return true, nil
}

func (d DB) GetFirstDelayReport(ctx context.Context) (delayreportentity.DelayReport, error) {
	const op = "mysqldelayreport.GetFirstDelayReport"
	row := d.adapter.Conn().QueryRowContext(ctx, `select * from delay_reports where delay_check = false And  agent_id is null ORDER BY created_at ASC LIMIT 1`)
	report, err := scanDelayReport(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return delayreportentity.DelayReport{}, nil
		}
		return delayreportentity.DelayReport{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return report, nil
}

func (d DB) AddAgentDelayReport(ctx context.Context, AgentID uint, DelayReportID uint) error {
	const op = "mysqldelayreport.AddAgentDelayReport"
	_, err := d.adapter.Conn().ExecContext(ctx, `update delay_reports set agent_id=? where id=?`, AgentID, DelayReportID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
func scanDelayReport(scanner mysql.Scanner) (delayreportentity.DelayReport, error) {
	var report delayreportentity.DelayReport
	var createdAt time.Time
	var deliveryTime sql.NullTime
	var agentID sql.NullInt64
	err := scanner.Scan(&report.ID, &report.OrderID, &agentID, &report.DelayCheck, &deliveryTime, &createdAt)

	if agentID.Valid {
		report.AgentID = uint(agentID.Int64)
	}
	if deliveryTime.Valid {
		report.DeliveryTime = deliveryTime.Time
	}
	report.CreatedAt = createdAt
	return report, err
}