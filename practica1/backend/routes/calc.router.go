package routes

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Cal!!!"))
}

func Suma(w http.ResponseWriter, r *http.Request) {}

func Resta(w http.ResponseWriter, r *http.Request) {}

func Multiplicacion(w http.ResponseWriter, r *http.Request) {}

func Division(w http.ResponseWriter, r *http.Request) {}
