package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ArturoArreola/demo-social-network/middlew"
	"github.com/ArturoArreola/demo-social-network/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores es la función donde se setea el puerto y el Handler, y escucha el servidor
func Manejadores() {
	router := mux.NewRouter()

	// API Endpoints
	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/perfil", middlew.CheckDB(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.CheckDB(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")

	router.HandleFunc("/nuevoPost", middlew.CheckDB(middlew.ValidarJWT(routers.NuevaPublicacion))).Methods("POST")
	router.HandleFunc("/obtenerPosts", middlew.CheckDB(middlew.ValidarJWT(routers.ObtenerPublicacion))).Methods("GET")
	router.HandleFunc("/borrarPost", middlew.CheckDB(middlew.ValidarJWT(routers.BorrarPublicacion))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.CheckDB(middlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.CheckDB(middlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.CheckDB(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckDB(middlew.ValidarJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckDB(middlew.ValidarJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckDB(routers.ConsultaRelacion)).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.CheckDB(routers.ListaUsuarios)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
