package delayreportvalidator

import (
	"fmt"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateDelayReportRequest(req delayreportparam.DelayReportRequest) (map[string]string, error) {
	const op = "delayreportvalidator.ValidateDelayReportRequest"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.OrderID,
			validation.Required,
			validation.By(v.checkIsOrderTimeDelivery)),
	); err != nil {
		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}

func (v Validator) checkIsOrderTimeDelivery(value interface{}) error {
	orderID := value.(uint)
	isExists, orderFound, err := v.order.IsOrderExceedingTheTimeDelivery(orderID)
	if err != nil {
		if !orderFound {
			return fmt.Errorf(errmsg.ErrorMsgOrderIDNotValid)
		}
		if !isExists {
			return fmt.Errorf(errmsg.ErrorMsgOrderMayDelay)
		}
	} else {
		if !isExists {
			return fmt.Errorf(errmsg.ErrorMsgOrderMayDelay)
		}
	}
	return nil
}
