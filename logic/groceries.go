package logic

import (
	"io/ioutil"
	"os"
	"sync"
)

type GroceryList struct {
	Title string
	Body  []byte
	mu    sync.Mutex
}

func (gl GroceryList) Save() error {
	gl.mu.Lock()
	err := ioutil.WriteFile("./resources/grocery-list.txt", gl.Body, os.ModeDevice)
	gl.mu.Unlock()
	return err
}

func NewGroceryList() GroceryList {
	var gl GroceryList
	gl.Title = "grocery-list"
	return gl
}
