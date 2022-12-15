package controllers

import (
	"bytes"
	"html/template"
	"net/http"
	"time"

	"github.com/danakin/festor.info/ui"
)

type Controllers struct {
	Homepage   *Homepage
	Blog       *Blog
	Technology *Technology
	Contact    *Contact
	CV         *CV
	Project    *Project
	Error      *Error
}

func NewControllers() *Controllers {
	return &Controllers{
		Homepage:   NewHomepageController(),
		Blog:       NewBlogController(),
		Technology: NewTechnologyController(),
		Contact:    NewContactController(),
		CV:         NewCVController(),
		Project:    NewProjectController(),
		Error:      NewErrorController(),
	}
}

type templateData struct {
	CurrentYear int
	Slug        string
}

func view(w http.ResponseWriter, r *http.Request, view string, data *templateData) {
	ts, err := template.
		New(view).
		ParseFS(
			ui.EmbeddedFiles,
			"templates/layouts/app.layout.tmpl",
			"templates/partials/*.partial.tmpl",
			view,
		)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buf, "app", addDefaultData(data, r))
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	buf.WriteTo(w)
}

func addDefaultData(data *templateData, r *http.Request) *templateData {
	if data == nil {
		data = &templateData{}
	}

	data.CurrentYear = time.Now().Year()
	return data
}
