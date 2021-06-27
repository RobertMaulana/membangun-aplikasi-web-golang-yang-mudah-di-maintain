package main

import (
	"log"
	"net/http"
)

// Membuat fungsi handler home yang akan menulis byte slice berisi
// "Hello from Pastetext" pada respon body untuk client.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Pastetext"))
}

func main() {
	// Penggunaan fungsi http.NewServeMux() untuk inisialisasi awal servemux yang baru, kemudian
	// daftarkan fungsi home diatas sebagai handler route dengan URL pattern "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Penggunaan fungsi http.ListenAndServe() untuk memulai web server baru.
	// Ada 2 paramater: alamat TCP network untuk listen di port (di case ini ":8000")
	// dan servemux yang baru saja kita buat.
	// Jika http.ListenAndServe() terjadi error, kita menggunakan fungsi log.Fatal()
	// untuk menapilkan log error kemudian keluar.
	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}