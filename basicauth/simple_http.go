package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// func handle(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(`{"name": "jetsadawwts"}`))
// }

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Health struct {
	Status string `json:"status"`
}

var users = []User{
	{ID: 1, Name: "jetsadawwts", Age: 26},
	{ID: 2, Name: "wongwit", Age: 25},
}

var healths = []Health{
	{Status: "OK"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	//Log Auth
	// u, p , ok := r.BasicAuth()
	// log.Println("auth:", u, p, ok)

	if r.Method == "GET" {
		b, err := json.Marshal(users)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(b)
		return
	}

	if r.Method == "POST" {
		// b, err := json.Marshal(users)
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		users = append(users, u)
		fmt.Fprintf(w, "hello %s create users", "POST")
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(healths)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

//Middleware Individual endpoint
// func logMiddleware(Handler http.HandlerFunc) http.HandlerFunc {
// 	return func(res http.ResponseWriter, req *http.Request) {
// 		start := time.Now()
// 		Handler.ServeHTTP(res, req)
// 		log.Printf("Server http middleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
// 	}
// }

// func main() {
// 	http.HandleFunc("/users",logMiddleware(userHandler))
// 	http.HandleFunc("/health", logMiddleware(healthHandler))

// 	log.Println("Server started at :2565")
// 	log.Fatal(http.ListenAndServe(":2565" ,nil))
// 	log.Println("bye bye!")

// }

//Auth Middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user, pass, ok := req.BasicAuth()
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte(`{"message": "can't parse the basic auth."}`))
			return
		}
		if user != "apidesign" || pass != "45678" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"massage": "username or password incorrect."}`))
			return
		}
		fmt.Println("Auth passed.")
		next(w, req)
	}
}

// Middleware with Mux
type Logger struct {
	Handler http.Handler
}

// Medhod
func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", AuthMiddleware(userHandler))
	mux.HandleFunc("/health", healthHandler)

	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}

	log.Println("Server started at :2565")
	log.Fatal(srv.ListenAndServe())
	log.Println("bye bye!")
}
