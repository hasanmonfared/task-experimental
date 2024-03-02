package mysqlorder

import (
	"database/sql"
	"gameapp/adapter/mysql"
	"gameapp/entity/orderentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"time"
)

func (d DB) IsOrderExceedingTheTimeDelivery(orderID uint) (bool, error) {
	const op = "mysqlorder.IsOrderExceedingTheTimeDelivery"
	row := d.adapter.Conn().QueryRow(`select * from orders where id= ? AND delivery_time<= NOW() AND status = ?`, orderID, orderentity.ReadyToSendStatusStr)
	_, err := scanOrder(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotfound).WithKind(richerror.KindNotFound)
		}
		return false, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindNotFound)
	}
	return true, nil

}
func scanOrder(scanner mysql.Scanner) (orderentity.Order, error) {
	var createdAt time.Time
	var order orderentity.Order
	var statusStr string
	err := scanner.Scan(&order.ID, &order.UserID, &order.VendorID, &order.DeliveryTime, &statusStr, &createdAt)
	order.Status = orderentity.MapToStatusEntity(statusStr)
	return order, err
}
