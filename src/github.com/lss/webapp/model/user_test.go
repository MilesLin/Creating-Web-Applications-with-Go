package model
import "testing"

func TestLoginSendsCorrectPasswordHash(t *testing.T) {
	testDB := new(mockDB)
	testDB.returnedRow = &mockRow{}
	db = testDB

	password := "the pwd"
	email := "the email"

	Login(email, password)

	if testDB.lastArgs[1] != password {
		t.Errorf("login function failed")
	}

}