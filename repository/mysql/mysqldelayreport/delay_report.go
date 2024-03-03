package mysqldelayreport

import (
	"database/sql"
	"fmt"
	"gameapp/adapter/mysql"
	"gameapp/entity/delayreportentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (d DB) InsertDelayReport(ctx context.Context, vendorID uint, orderID uint, deliveryTime uint) error {
	const op = "mysqldelayreport.InsertDelayReport"
	var query string
	var args []interface{}

	if deliveryTime != 0 {
		query = `INSERT INTO delay_reports (vendor_id,order_id, delivery_time) VALUES (?,?, ?)`
		args = []interface{}{vendorID, orderID, deliveryTime}
	} else {
		query = `INSERT INTO delay_reports (vendor_id,order_id) VALUES (?,?)`
		args = []interface{}{vendorID, orderID}
	}

	_, err := d.adapter.Conn().ExecContext(ctx, query, args...)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
func (d DB) HasPendingDelayReport(ctx context.Context, orderID uint) (bool, error) {
	const op = "mysqldelayreport.HasPendingDelayReport"
	fmt.Println(orderID)
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

func (d DB) CheckAgentBusyInQueue(AgentID uint) (bool, error) {
	const op = "mysqldelayreport.CheckAgentBusyInQueue"
	row := d.adapter.Conn().QueryRow(`select * from delay_reports where delay_check = false And  agent_id =?`, AgentID)
	_, err := scanDelayReport(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return true, nil
}

func scanDelayReport(scanner mysql.Scanner) (delayreportentity.DelayReport, error) {
	var report delayreportentity.DelayReport
	var createdAt time.Time
	var deliveryTime sql.NullInt64
	var agentID sql.NullInt64
	err := scanner.Scan(&report.ID, &report.VendorID, &report.OrderID, &agentID, &report.DelayCheck, &deliveryTime, &createdAt)

	if agentID.Valid {
		report.AgentID = uint(agentID.Int64)
	}
	if deliveryTime.Valid {
		report.DeliveryTime = uint(deliveryTime.Int64)
	}
	report.CreatedAt = createdAt
	return report, err
}
