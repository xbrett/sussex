package service

import (
	"net/http"
	"time"

	"github.com/xbrett/sussex/logic"
)

func (di DisplayInfo) Chores(w http.ResponseWriter, r *http.Request) {
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
	di.logicDataAccess.SaveStateOfChores()
	//also need a call to save current chores somewhere in the db
}
