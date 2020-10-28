package register

import "testing"

func Test_Register(t *testing.T) {
	r := Register()
	if r != true {
		t.Errorf("Register(), Expected true, actual %t\n", r)
	}
}

