package main;

import (
    "net/http"
    "html/template"
);

var templates map[string]*template.Template;

func main() {
    templates = make(map[string]*template.Template);
    templates["index"] = template.Must(template.ParseFiles("template/index.html", "template/base.html"));
    
    http.HandleFunc("/", hIndex);
    http.ListenAndServe("127.0.0.1:8018", nil);
}

func hIndex(writer http.ResponseWriter, request *http.Request) {
    renderTemplate(writer, "index", nil);
}

func renderTemplate(writer http.ResponseWriter, templateResource string, data interface{}) {
    templateRenderingException := templates[templateResource].ExecuteTemplate(writer, templateResource + ".html", data);
    if templateRenderingException != nil {
        http.Error(writer, templateRenderingException.Error(), http.StatusInternalServerError);
    }
}