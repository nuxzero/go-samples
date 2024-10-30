package foo

import "net/http"

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func SayGoodbye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye, World!"))
}

func Talk(w http.ResponseWriter, r *http.Request) {
	message := r.PostFormValue("message")

	if message == "" {
		w.Write([]byte("Talk, nothing to say!"))
	}

	w.Write([]byte("Talk, " + message))
}
