package mysqlorderdelay

import (
	"database/sql"
	"gameapp/adapter/mysql"
	"gameapp/entity/orderdelayentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (d DB) InsertDelayReport(ctx context.Context, orderID uint) error {
	const op = "mysqlorderdelay.InsertDelayReport"
	_, err := d.adapter.Conn().ExecContext(ctx, `insert into delay_reports(order_id)values(?)`, orderID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
func (d DB) HasPendingDelayReport(ctx context.Context, orderID uint) (bool, error) {
	const op = "mysqlorderdelay.HasPendingDelayReport"

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
func scanDelayReport(scanner mysql.Scanner) (orderdelayentity.DelayReport, error) {
	var report orderdelayentity.DelayReport
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
