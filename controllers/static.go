package controllers

import (
	"html/template"
	"net/http"

	"github.com/sxc/oishifood/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	question := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is Oishi Food?",
			Answer:   "Oishi Food is a food delivery service that delivers food to your door.",
		},
		{
			Question: "How do I order food?",
			Answer:   "You can order food by visiting our website and selecting the food you want to order.",
		},
		{
			Question: "How do I pay for my order?",
			Answer:   "You can pay for your order by cash or credit card.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, question)
	}
}
