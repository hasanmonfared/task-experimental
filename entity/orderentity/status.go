package orderentity

type Status uint8

const (
	AssignedStatus = iota + 1
	AtVendorStatus
	PickedStatus
	DeliveredStatus
)

const (
	AssignedStatusStr  = "assigned"
	AtVendorStatusStr  = "at_vendor"
	PickedStatusStr    = "picked"
	DeliveredStatusStr = "delivered"
)

func (r Status) String() string {
	switch r {
	case AssignedStatus:

		return AssignedStatusStr
	case AtVendorStatus:

		return AtVendorStatusStr
	case PickedStatus:
		return PickedStatusStr
	case DeliveredStatus:
		return DeliveredStatusStr
	}
	return ""
}
func MapToStatusEntity(statusStr string) Status {
	switch statusStr {
	case AssignedStatusStr:
		return AssignedStatus
	case AtVendorStatusStr:
		return AtVendorStatus
	case PickedStatusStr:
		return PickedStatus
	case DeliveredStatusStr:
		return DeliveredStatus
	}
	return Status(0)
}
