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


//func main() {
//	length, err := test()
//	if err != nil {
//		log.Fatal(err)
//	} 
//	fmt.Printf("length=%d\n", length)
//}
//
//func test() (read int, err error) {
//	// turn on custom error handling
//	defer Handle(&err, "test() - Can't test file")
//
//	file, err := os.Open("go.mod0");
//	Check(&err) // Use custom error handling
//	defer file.Close()
//
//	buf := make([]byte, 100)
//	c, err := file.Read(buf);
//	Check(&err)
//
//	read = c
//
//	return 
//}
