package mysqltrip

import (
	"gameapp/adapter/mysql"
	"gameapp/entity/tripentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (d DB) GetTripByOrderID(ctx context.Context, orderID uint) ([]tripentity.Trip, error) {
	const op = "mysqlorderdelay.GetTripByOrderID"

	trips := make([]tripentity.Trip, 0)

	rows, err := d.adapter.Conn().QueryContext(ctx, `select * from trips where order_id= ?`, orderID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan()
		acl, err := scanTrip(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}
		trips = append(trips, acl)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)

	}
	return trips, nil
}
func scanTrip(scanner mysql.Scanner) (tripentity.Trip, error) {
	var createdAt time.Time
	var trip tripentity.Trip
	err := scanner.Scan(&trip.ID, &trip.OrderID, &trip.Status, &createdAt)
	return trip, err
}
