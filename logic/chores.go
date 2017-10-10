package logic

import (
	"log"

	"github.com/xbrett/sussex/database"
)

//ChoreList contains all of the chores to be done at Sussex
var ChoreList []database.ChorePair

func init() {
	brett := database.ChorePair{Name: "Brett", Chore: "Vacuum", ID: 1}
	quinn := database.ChorePair{Name: "Quinn", Chore: "Wipe Counters", ID: 2}
	jason := database.ChorePair{Name: "Jason", Chore: "Trash", ID: 3}
	paul := database.ChorePair{Name: "Paul", Chore: "Dishwasher", ID: 4}

	ChoreList = append(ChoreList, brett)
	ChoreList = append(ChoreList, quinn)
	ChoreList = append(ChoreList, jason)
	ChoreList = append(ChoreList, paul)
}

func RotateChores() {
	ChoreList = append(ChoreList[1:], ChoreList[0])
}

func GetCurrentChores() []database.ChorePair {
	return ChoreList
}

func (l *Logic) SaveStateOfChores() {
	var err error
	for i := 0; i <= 4; i++ {
		err = l.mydb.UpdateChores(ChoreList[i])
		if err != nil {
			log.Fatal("ChoreList not saved")
		}
	}
}
