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
		//if perr := recover(); perr != nil {
			//if e, ok := perr.(error); ok && *err != e {
			//	*err = e
			//}
		//}
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
