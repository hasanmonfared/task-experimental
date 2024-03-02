package mysqlorderdelay

import (
	"fmt"
	"golang.org/x/net/context"
)

func (d DB) InsertDelayReport(ctx context.Context, orderID uint) error {
	_, err := d.adapter.Conn().ExecContext(ctx, `insert into delay_reports(order_id)values(?)`, orderID)
	if err != nil {
		return fmt.Errorf("can't execute command:%w", err)
	}
	return nil
}
