package main

import(
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from pastetext"))
}

// Menambahkan fungsi showSnippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet"))
}

// Menambahkan fungsi createSnippet handler
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}