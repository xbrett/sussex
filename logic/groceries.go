package logic

import (
	"io/ioutil"
	"os"
)

type GroceryList struct {
	Title string
	Body  []byte
}

func (gl GroceryList) Save() error {
	return ioutil.WriteFile("./resources/grocery-list.txt", gl.Body, os.ModeDevice)
}

func NewGroceryList() GroceryList {
	var gl GroceryList
	gl.Title = "grocery-list"
	return gl
}
