package legacy_http

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func ServeApplication() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./cmd/ui/static/"))
	//For matching paths, we also strip the "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/downloadFile", downloadFileHandler)

	mux.HandleFunc("/", Home)
	mux.HandleFunc(" /snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	port := AssignPort()
	log.Printf(fmt.Sprintf("Starting server on %s", port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal().Err(err)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./cmd/ui/html/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, err = fmt.Fprintf(w, "Display a specific snippet with ID %d...", id) // calls w.Write([]byte("Display a...")) internally
	if err != nil {
		return
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Changing the response header map after a call to w.WriteHeader() or
		//w.Write() will have no effect on the headers that the user receives. You need to make
		//sure that your response header map contains all the headers you want before you call
		//these methods.

		// Suppressing system generated header
		w.Header()["Date"] = nil

		//Add to response header map
		w.Header().Set("Allow", http.MethodPost)
		// It’s only possible to call w.WriteHeader() once per response, and after the status code has
		// been written it can’t be changed.
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		//The last 2 lines can be replaced with: http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

/*
* Supports Range requests (https://web.archive.org/web/20230918195519/https://benramsey.com/blog/2008/05/206-partial-content-and-range-requests/)
* http.ServeFile() does not automatically sanitize the file path. If you’re constructing a file path from untrusted
* user input, to avoid directory traversal attacks you must sanitize the input with filepath.Clean() before using it.
 */
func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Query().Get("fileName") == "fileA.zip" {
		http.ServeFile(w, r, "./cmd/ui/static/file1.7z")
	} else {
		http.NotFound(w, r)
	}
}

func AssignPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //defaultPort uint16
	}
	return port
}
