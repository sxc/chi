package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
