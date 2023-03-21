package chkerr

import (
	"fmt"
)

func Handle(err *error, msg ...string) {
	if *err != nil {
		if len(msg) > 0 {
			*err = fmt.Errorf("%s: %w", msg[0], *err)
		}
		// Intercept only our custom panic from Check
		recover()
	}
}

func Check(err *error, msg ...string) {
	if *err != nil {
		if len(msg) > 0 {
			*err = fmt.Errorf("%s: %w", msg[0], *err)
		}
		panic(*err)
	}
}
