package logic

//ChoreList contains all of the chores to be done at Sussex
var ChoreList []string

func init() {
	ChoreList = append(ChoreList, "Vacuum")
	ChoreList = append(ChoreList, "Wipe Table/Counters")
	ChoreList = append(ChoreList, "Trash")
	ChoreList = append(ChoreList, "Dishwasher")
}

func RotateChores() {
	ChoreList = append(ChoreList[1:], ChoreList[1])
}

func GetCurrentChores() []string {
	return ChoreList
}
