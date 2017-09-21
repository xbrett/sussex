package service

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xbrett/sussex/logic"
)

//DisplayInfo is for accessing logic data
type DisplayInfo struct {
	logicDataAccess logic.Logic
	gl              logic.GroceryList
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func Chores(w http.ResponseWriter, r *http.Request) {
	//add time switch to rotate chore assignments
	tpl.ExecuteTemplate(w, "chores.gohtml", logic.GetCurrentChores())
	ticker := time.NewTicker(7 * 24 * time.Second)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				// do something every week as defined by ticker above
				logic.RotateChores()
			}
		}
	}(ticker)
	//also need a call to save current chores somewhere in the db
}

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

	//.save line
	err := ioutil.WriteFile("./resources/grocery-list.txt", []byte(body), 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/grocery-list", http.StatusFound)
}

func insertNewLine(str string) string {
	return ""
}
