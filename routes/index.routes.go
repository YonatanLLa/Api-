package routes

import "net/http"

// r -> tener acceso al parametro que te envian los clientes.

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello wold! 3"))
}
