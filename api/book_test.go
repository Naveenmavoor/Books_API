package api

import (
	"fmt"
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "ML REIMER", ISBN: "012345678"}
	json := book.ToJSON()

	if book == fromJSON(json) {
		fmt.Print("success")
	}
}
func TestBookFromJSON(t *testing.T) {

	json := []byte(`{"Title":"Cloud Native Go","Author":"ML Reimer","ISBN":"0123456789"}`)
	book := fromJSON(json)
	checkbook := Book{Title: "Cloud Native Go", Author: "ML REIMER", ISBN: "012345678"}
	if book == checkbook {
		fmt.Print("success")
	}

}
