package model

import (
	"io/ioutil"
)

//Page page structure
type Page struct {
	Title string
	Body  []byte
}

//Save save page method
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//LoadPage load page function
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
			return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}