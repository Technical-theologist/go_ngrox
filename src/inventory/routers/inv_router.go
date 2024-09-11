package routers

import (
	"inventory/controllers"
	"inventory/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", controllers.GetOneProduct).Methods("GET")

	//router.HandleFunc("/api/products", controllers.GetAllProducts).Methods("GET")
	router.Handle("/admin/products", utils.AuthMiddleware(utils.RoleAuthorizationMiddleware("admin")(http.HandlerFunc(controllers.GetAllProducts)))).Methods("GET")

	router.HandleFunc("/api/products/search", controllers.SearchByTitle).Methods("GET")
	router.HandleFunc("/api/product", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/product/update/{id}", controllers.UpdateOneProduct).Methods("PUT")
	router.HandleFunc("/api/product/filter", controllers.FilterBy).Methods("POST")
	router.HandleFunc("/api/bulkdata", controllers.BulkDataUpload).Methods("POST")
	router.HandleFunc("/api/createuser", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/userlogin", controllers.UserLogin).Methods("POST")
	router.HandleFunc("/api/upload", controllers.FileUpload).Methods("POST")
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("C:\\"))))
	return router
}
