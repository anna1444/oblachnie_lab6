package main
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		bakes, err := DbGetBakes()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, bakes)
	}
}

func addBakeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		typename := r.Form.Get("typename")
		price := r.Form.Get("price")
		madedate :=r.Form.Get("madedate")
		expiration := r.Form.Get("expiration")
		err := DbAddBake(typename,price, madedate,expiration)
		if err != nil {
			log.Fatal(err) }
	} }

func getMaxHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_max_form.html")
		if err != nil {
			log.Fatal(err)
		}
		bakes, err := DbGetMaximum()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, bakes)
	}}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port }


func main() {
	err := DbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addBakeHandler)
	http.HandleFunc("/max", getMaxHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}