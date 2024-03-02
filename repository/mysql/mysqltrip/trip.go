package mysqltrip

import (
	"database/sql"
	"gameapp/adapter/mysql"
	"gameapp/entity/tripentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (d DB) GetTripByOrderID(ctx context.Context, orderID uint) (tripentity.Trip, error) {
	const op = "mysqlorderdelay.GetTripByOrderID"

	row, err := d.adapter.Conn().QueryContext(ctx, `select * from trips where order_id= ?`, orderID)
	trip, err := scanTrip(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return tripentity.Trip{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotfound).WithKind(richerror.KindNotFound)
		}
		return tripentity.Trip{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return trip, nil
}
func scanTrip(scanner mysql.Scanner) (tripentity.Trip, error) {
	var createdAt time.Time
	var trip tripentity.Trip
	var tripStatusStr string
	err := scanner.Scan(&trip.ID, &trip.OrderID, &tripStatusStr, &createdAt)
	trip.Status = tripentity.MapToStatusEntity(tripStatusStr)
	return trip, err
}
