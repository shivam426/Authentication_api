package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	costumer := &Costumer{}
	err := json.NewDecoder(r.Body).Decode(&costumer)

	if err != nil {
		log.Fatal(err)
	}
	costumer.Password = getHash([]byte(costumer.Password))
	sqlStatement := `INSERT INTO "costumer" (email,name,pwd) VALUES ($1,$2,$3)`
	result, err := conn.Exec(sqlStatement, costumer.Email, costumer.Name, costumer.Password)
	if err != nil {
		panic(err)
	}

	fmt.Println(costumer)

	w.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["response"] = "registered Successfully"
	des, _ := json.Marshal(res)
	w.Write([]byte(des))
	fmt.Println(result)

}
