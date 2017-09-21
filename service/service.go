package service

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	bindTo string
)

func init() {
	flag.StringVar(&bindTo, "listen", ":9000", "host:port to bind to")
}

func ListenAndServe() {

	//disInfo := DisplayInfo{logicDataAccess: logic.New(), gl: logic.NewGroceryList()}
	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	router.HandleFunc("/chores", Chores)
	router.HandleFunc("/view/grocery-list", ViewGroceries)
	router.HandleFunc("/edit/grocery-list", EditGroceries)
	router.HandleFunc("/save/grocery-list", SaveGroceries)

	log.Println("Listening...(port: 9000)")
	http.ListenAndServe(bindTo, router)
}
