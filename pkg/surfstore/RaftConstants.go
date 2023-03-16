package surfstore

import (
	"fmt"
)

var (
	ERR_SERVER_CRASHED   = fmt.Errorf("Server is crashed.")
	ERR_NOT_LEADER       = fmt.Errorf("Server is not the leader")
	ERR_MAJORITY_CRASHED = fmt.Errorf("Majority servers are crashed")
)
