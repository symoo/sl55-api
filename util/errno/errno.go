package errno

import (
	"fmt"
)

type Errno struct {
	Code int
	Msg string
}

func (err Errno) Error() string {
	return err.Msg
}

