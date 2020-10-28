package login

import "testing"

func Test_Login(t *testing.T) {
	r := Login()
	if r != true {
		t.Errorf("Login(),Expected true,actually got %t\n", r)
	}
}
