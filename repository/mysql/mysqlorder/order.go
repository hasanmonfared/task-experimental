package mysqlorder

import (
	"database/sql"
	"gameapp/adapter/mysql"
	"gameapp/entity/orderentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (d DB) GetDetailOrderByID(ctx context.Context, orderID uint) (orderentity.Order, error) {
	const op = "mysqlorder.GetDetailOrderByID"
	row := d.adapter.Conn().QueryRowContext(ctx, `select * from orders where id= ?`, orderID)
	order, err := scanOrder(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return orderentity.Order{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotfound).WithKind(richerror.KindNotFound)
		}
		return orderentity.Order{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindNotFound)
	}
	return order, nil
}
func scanOrder(scanner mysql.Scanner) (orderentity.Order, error) {
	var createdAt time.Time
	var order orderentity.Order
	var statusStr string
	err := scanner.Scan(&order.ID, &order.UserID, &order.VendorID, &order.DeliveryTime, &statusStr, &createdAt)
	order.Status = orderentity.MapToStatusEntity(statusStr)
	order.CreatedAt = createdAt
	return order, err
}
