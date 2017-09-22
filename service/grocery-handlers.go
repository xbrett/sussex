package service

import (
	"io/ioutil"
	"net/http"
)

func ViewGroceries(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("./resources/grocery-list.txt")
	tpl.ExecuteTemplate(w, "view-groceries.gohtml", string(body))
}

func EditGroceries(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("./resources/grocery-list.txt")
	tpl.ExecuteTemplate(w, "edit-groceries.gohtml", string(body))
}

func SaveGroceries(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")

	err := ioutil.WriteFile("./resources/grocery-list.txt", []byte(body), 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/grocery-list", http.StatusFound)
}
