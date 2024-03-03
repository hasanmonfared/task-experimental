package delayreportvalidator

import (
	"fmt"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateGetDelayReportRequest(req delayreportparam.GetDelayReportRequest) (map[string]string, error) {
	const op = "delayreportvalidator.ValidateGetDelayReportRequest"
	if err := validation.ValidateStruct(&req,

		validation.Field(&req.AgentID,
			validation.Required,
			validation.By(v.checkExistsAgent),
			validation.By(v.checkAgentInQueue),
		),
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

func (v Validator) checkExistsAgent(value interface{}) error {
	AgentID := value.(uint)

	if isExists, err := v.agent.ExistsAgentID(AgentID); err == nil {
		if !isExists {
			return fmt.Errorf(errmsg.ErrorMsgAgentIDNotValid)
		}
	}
	return nil
}

func (v Validator) checkAgentInQueue(value interface{}) error {
	AgentID := value.(uint)
	if isExists, err := v.repo.CheckQueueAgentID(AgentID); err == nil {
		if isExists {
			return fmt.Errorf(errmsg.ErrorMsgOrderReview)
		}
	}
	return nil
}
