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
	const op = "mysqltrip.GetTripByOrderID"

	row := d.adapter.Conn().QueryRowContext(ctx, `select * from trips where order_id= ?`, orderID)
	trip, err := scanTrip(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return tripentity.Trip{}, nil
		}
		return tripentity.Trip{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return trip, nil
}
func scanTrip(scanner mysql.Scanner) (tripentity.Trip, error) {
	var trip tripentity.Trip
	var createdAt time.Time
	var tripStatusStr string
	err := scanner.Scan(&trip.ID, &trip.OrderID, &tripStatusStr, &createdAt)
	trip.Status = tripentity.MapToStatusEntity(tripStatusStr)
	trip.CreatedAt = createdAt
	return trip, err
}
