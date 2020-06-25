package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/ArturoArreola/demo-social-network/bd"
)

// ListaUsuarios
func ListaUsuarios (w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page :=  r.URL.Query().Get("page")
	search :=  r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar la página como parámetro entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeerUsuariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Erro al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
