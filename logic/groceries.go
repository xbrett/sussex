package logic

import (
	"io/ioutil"
	"os"
	"sync"
)

//GroceryList is used to store list info
type GroceryList struct {
	Title string
	Body  []byte
	mu    *sync.Mutex
}

//Save writes to the file grocery-list.txt with the updated GroceryList
func (gl GroceryList) Save() error {
	gl.mu.Lock()
	err := ioutil.WriteFile("./resources/grocery-list.txt", gl.Body, os.ModeDevice)
	gl.mu.Unlock()
	return err
}

//NewGroceryList generates a new empty body GroceryList titled "grocery-list"
func NewGroceryList() GroceryList {
	var gl GroceryList
	gl.Title = "grocery-list"
	return gl
}
