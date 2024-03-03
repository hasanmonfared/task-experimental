package estimateentity

import "time"

type Estimate struct {
	NewEstimate time.Time `json:"new_estimate"`
}
