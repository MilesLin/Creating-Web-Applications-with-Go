package model

import(
	"fmt"
	"database/sql"
	"time"
	"log"
)

const passwordSalt = "sdfjklsdjfvoi72384234uklsjdszdxncv09823oweqiproweqxdcvxcz"

type User struct {
	Email		string
	Password	string
}

func Login(email, password string) (*User, error) {
	result := &User {}

	row := db.QueryRow(`
		SELECT email, password
		FROM myuser
		WHERE email = $1
		  AND password = $2`, email,password)
	
		  fmt.Println(row)

	err := row.Scan(&result.Email, &result.Password)

	fmt.Println(result)
	fmt.Println(err)

	switch {
		case err == sql.ErrNoRows:
			return nil, fmt.Errorf("User not found")
		case err != nil:
			return nil, err
	}

	t := time.Now()
	
	ur, err := db.Exec(`
		UPDATE myuser
		SET lastlogin = $1
		WHERE email = $2
		AND password = $3`, 
	  t, email, password)
	
	if err != nil {
		log.Printf("failed to update login time")
	}

	fmt.Println(ur)

	return result, nil


}