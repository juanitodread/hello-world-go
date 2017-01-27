package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)

  if err != nil {
    p = &Page{Title: title}
  }

  renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmplate string, p *Page) {
  t, _ := template.ParseFiles(tmplate + ".html")
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.ListenAndServe(":3535", nil)
}
