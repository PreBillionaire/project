package routers

import (
	"github.com/PreBillionaire/mongoAPI/controllers"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/userSignup", controllers.UserSignup).Methods("POST")
	r.HandleFunc("/userLogin", controllers.UserLogin).Methods("POST")
	r.HandleFunc("/addEmployee", controllers.AddEmployee).Methods("POST")
	r.HandleFunc("/getAllEmp", controllers.GetAllEmployee).Methods("GET")
	r.HandleFunc("/updateEmp", controllers.UpdateEmployee).Methods("POST")
	r.HandleFunc("/delAEmp/{username}", controllers.DeleteOneEmployee).Methods("DELETE","POST")
	r.HandleFunc("/delAllEmp", controllers.DeleteAllEmployees).Methods("DELETE")
	return r
}
