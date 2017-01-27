package main

import (
  "encoding/json"
  "html/template"
  "io/ioutil"
  "net/http"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type Page struct {
  Title string `json:"title"`
  Body  []byte `json:"body"`
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
  p, err := loadPage(title)

  if err != nil {
    http.Redirect(w, r, "/edit/" + title, http.StatusFound)
    return
  }

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

func saveHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/save/"):]
  body := r.FormValue("body")
  p := &Page{Title: title, Body: []byte(body)}
  err := p.save()

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func jsonViewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/"):]
  p, _ := loadPage(title)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(p)
}

func renderTemplate(w http.ResponseWriter, tmplate string, p *Page) {
  err := templates.ExecuteTemplate(w, tmplate + ".html", p)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.HandleFunc("/json/", jsonViewHandler)

  http.ListenAndServe(":3535", nil)
}
