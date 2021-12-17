package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	costumers := &Costumer{}
	dbcostumers := &Costumer{}

	err := json.NewDecoder(r.Body).Decode(&costumers)
	if err != nil {
		panic(err)
	}

	// param := r.URL.Query().Get("id")
	result, err := conn.Query(`SELECT * FROM costumer where email='` + costumers.Email + `'`)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		err = result.Scan(&dbcostumers.Id, &dbcostumers.Name, &dbcostumers.Password, &dbcostumers.Email)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(dbcostumers)
		default:
			panic(err)
		}

	}

	costumerpass := []byte(costumers.Password)
	dbcostumerpass := []byte(dbcostumers.Password)

	passErr := bcrypt.CompareHashAndPassword(dbcostumerpass, costumerpass)
	if passErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"response":"Wrong Password!"}`))
		fmt.Println(passErr)
	} else {

		jwtToken, err := GenerateJWT(costumers.Email)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"` + err.Error() + `"}`))
			return
		}

		addCookie(w, "Bearer", jwtToken)
	}
	w.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["response"] = "Login Successfully"
	des, _ := json.Marshal(res)
	w.Write([]byte(des))
	// json.NewEncoder(w).Encode(res)

	fmt.Println(result)
}
