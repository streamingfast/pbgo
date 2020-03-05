package v1

import (
	"fmt"
)

func (t *TransactionCursor) Key() string {
	return fmt.Sprintf("%s", t.TransactionHash)
}
