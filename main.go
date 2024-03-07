package main

import (
	"encoding/json"
	"go-web-dasar/conf"
	"go-web-dasar/handles"
	"log"
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
	// mux := new(CustomMux)

	// mux.HandleFunc("/student", ActionStudent)

	// mux.RegisterMiddleware(MiddlewareAuth)
	// mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	// server := new(http.Server)
	// server.Addr = ":9000"
	// server.Handler = mux

	// fmt.Println("server started at localhost:9000")
	// server.ListenAndServe()

	// penggunaan cookie
	// cokies.Cookies()

	// penggunaan cimple configuration
	// router := new(CustomMux)
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	// router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("How are you?"))
	// })

	// server := new(http.Server)
	// server.Handler = router
	// server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
	// server.WriteTimeout = conf.Configuration().Server.WriteTimeout * time.Second
	// server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	// if conf.Configuration().Log.Verbose {
	// 	log.Printf("Starting server at %s \n", server.Addr)
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }

	// Server Handler HTTP Request Cancellation
	handles.Handle()
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

type CustomMuxs struct {
	http.ServeMux
}

func (c CustomMuxs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration().Log.Verbose {
		log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
	}

	c.ServeMux.ServeHTTP(w, r)
}
