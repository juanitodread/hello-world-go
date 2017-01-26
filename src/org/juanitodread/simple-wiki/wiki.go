package main

import (
  "fmt"
  "io/ioutil"
)

type Page struct {
  Title string
  Body  []byte
}

// Saves the Page body object to a text file. The textfile is named according
// the page title.
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600) // 0600 are the permisions for the file
}

// Reads a page from file.
func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)

  if err != nil {
    return nil, err
  }

  return &Page{Title: title, Body: body}, nil
}

func main() {
  p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page")}
  p1.save()
  p2, _ := loadPage("TestPage")
  fmt.Println(string(p2.Body))
}
