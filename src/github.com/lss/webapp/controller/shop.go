package controller

import (
	"creating-web-applications-go/src/github.com/lss/webapp/viewmodel"
	"html/template"
	"net/http"
)

type shop struct {
	shopTemplate *template.Template
}

func (h shop) registerRoutes() {
	http.HandleFunc("/shop", h.handleShop)
}

func (h shop) handleShop(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewShop()
	h.shopTemplate.Execute(w, vm)
}
