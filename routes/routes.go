package routes

import (
	"net/http"
	"sistem-peminjaman-lab/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Home
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/lab", http.StatusSeeOther)
	})

	// Lab CRUD
	r.HandleFunc("/lab", controllers.LabIndex)
	r.HandleFunc("/lab/create", controllers.LabCreate)
	r.HandleFunc("/lab/edit/{id:[0-9]+}", controllers.LabEdit)
	r.HandleFunc("/lab/delete/{id:[0-9]+}", controllers.LabDelete)

	// Alat CRUD
	r.HandleFunc("/alat", controllers.AlatIndex)
	r.HandleFunc("/alat/create", controllers.AlatCreate)
	r.HandleFunc("/alat/edit/{id:[0-9]+}", controllers.AlatEdit)
	r.HandleFunc("/alat/delete/{id:[0-9]+}", controllers.AlatDelete)

	// Cetak PDF & Excel
	r.HandleFunc("/lab/pdf", controllers.LabPDF)
	r.HandleFunc("/lab/excel", controllers.LabExcel)
	r.HandleFunc("/alat/pdf", controllers.AlatPDF)
	r.HandleFunc("/alat/excel", controllers.AlatExcel)
	r.HandleFunc("/report", controllers.ReportIndex)

	return r
}
