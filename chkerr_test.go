package chkerr

import (
	"strings"
	"testing"
	"errors"
)

type ErrBase struct {
	msg string 
}

func (err *ErrBase) Error() string {
	return err.msg
}

func TestHandle(t *testing.T) {
	errMsg := "Base error"
	baseErr := &ErrBase{errMsg}
	var err error = baseErr
	Handle(&err)
	if e := err.Error(); e != errMsg {
		t.Errorf("Expected '%s', but got %s", errMsg, e)
	}
	altMsg := "Test msg"
	Handle(&err, altMsg)
	if e := err.Error(); strings.Index(e, altMsg) != 0 {
		t.Errorf("Error doesn't contain '%s' message", altMsg)
	}
}


func TestHandle_CheckErrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Handle func has not recovered Check-panic")
		}
	}()
	func() {
		var err error
		// setup err handler
		defer Handle(&err)
		// call some func which returns err
		err = errors.New("Test error")
		// and after that Check call simulation 
		panic("check")
	}()
}


func TestHandle_OtherPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Handle func has recovered Other panics")
		}
	}()
	var err error
	defer Handle(&err)
	// simulate other panic before err is assigned
	panic("exception")
}
