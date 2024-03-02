package estimateentity

import "time"

type Estimate struct {
	NewEstimate time.Duration `json:"new_estimate"`
}
