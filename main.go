package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// penggunaan hello word
	// helloword.HelloWord()

	// penggunaan assets
	// sets.Assets()

	// penggunaan html
	// htmls.Htmls()

	// penggunaan render partial html
	// renderspartial.ParseGlob()
	// renderspartial.ParseGlob()

	// penggunaan template actions and variabel
	// templatesaction.Templates()

	// penggunaan predefiend
	// predefiend.Predefiend()

	// penggunaan custome function
	// customefunction.Custom()

	// penggunaan render template
	// renders.Renders()

	// penggunaan render string
	// renderstrings.RenderStrings()

	// penggunaan form value
	// formvalues.FormValues()

	// penggunaan form value untuk upload gambar
	// uploads.Uploads()

	// penggunaan payload ajax json
	// payloads.Payloads()

	// penggunaan ajax response
	// ajxresponse.Responses()

	// penggunaan basic auth
	// http.HandleFunc("/student", ActionStudent)

	// server := new(http.Server)
	// server.Addr = ":8081"

	// fmt.Println("server started at localhost:8081")
	// server.ListenAndServe()

	// penggunaan muddleware
	// mux := http.DefaultServeMux

	// mux.HandleFunc("/student", ActionStudents)

	// var handler http.Handler = mux
	// handler = MiddlewareAuth(handler)
	// handler = MiddlewareAllowOnlyGet(handler)

	// server := new(http.Server)
	// server.Addr = ":9000"
	// server.Handler = handler

	// fmt.Println("server started at localhost:9000")
	// server.ListenAndServe()

	// penggunaan custome mux
	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)

	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}

func ActionStudents(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}
