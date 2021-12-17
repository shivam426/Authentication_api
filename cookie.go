package main

import (
	"fmt"
	"net/http"
)

func addCookie(w http.ResponseWriter, name, value string) {
	fmt.Println(name, value)
	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: value,
	})
}
