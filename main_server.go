package cooking4mula

import (
  "fmt"
  "html/template"
  "net/http"
)

func init() {
  http.HandleFunc("/", homepageHandler)
}

func homepageHandler(writer http.ResponseWriter,
                     reader *http.Request) {
  writer.Header().Set("Content-Type", "text/html")
  err := homepageTemplate.Execute(writer, nil)

  if err != nil {
    http.Error(writer,
               "Sorry we are having technical issues. Please try again later. ",
               http.StatusInternalServerError)
    fmt.Println("Error: Could not create template: ", err.Error())
  }
}

const homepageTemplateFilePath = "src/homepage_template.html"
var homepageTemplate = template.Must(
    template.ParseFiles(homepageTemplateFilePath))
