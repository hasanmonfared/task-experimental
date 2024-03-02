package orderentity

type Status uint8

const (
	SubmitStatus = iota + 1
	InDoingStatus
	ReadyToSendStatus
)

const (
	SubmitStatusStr      = "submit"
	InDoingStatusStr     = "in_doing"
	ReadyToSendStatusStr = "ready_to_send"
)

func (r Status) String() string {
	switch r {
	case SubmitStatus:

		return SubmitStatusStr
	case InDoingStatus:

		return InDoingStatusStr
	case ReadyToSendStatus:
		return ReadyToSendStatusStr
	}
	return ""
}
func MapToStatusEntity(statusStr string) Status {
	switch statusStr {
	case SubmitStatusStr:
		return SubmitStatus
	case InDoingStatusStr:
		return InDoingStatus
	case ReadyToSendStatusStr:
		return ReadyToSendStatus
	}
	return Status(0)
}
