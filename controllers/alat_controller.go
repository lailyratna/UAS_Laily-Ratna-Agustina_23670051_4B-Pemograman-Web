package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"

	"sistem-peminjaman-lab/config"
)

func AlatIndex(w http.ResponseWriter, r *http.Request) {
	var alats []config.Alat
	config.DB.Preload("Lab").Find(&alats)
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/alat/index.html"))
	tmpl.Execute(w, alats)
}

func AlatCreate(w http.ResponseWriter, r *http.Request) {
	var labs []config.Lab
	config.DB.Find(&labs)

	if r.Method == http.MethodPost {
		alat := config.Alat{
			NamaAlat: r.FormValue("nama_alat"),
			Jumlah:   Atoi(r.FormValue("jumlah")),
			Kondisi:  r.FormValue("kondisi"),
			LabID:    uint(Atoi(r.FormValue("lab_id"))),
		}
		config.DB.Create(&alat)
		http.Redirect(w, r, "/alat", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Title":  "Tambah Alat",
		"Action": "/alat/create",
		"Alat":   config.Alat{},
		"Labs":   labs,
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/alat/form.html"))
	tmpl.Execute(w, data)
}

func AlatEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/alat/edit/"):len(r.URL.Path)]
	var alat config.Alat
	var labs []config.Lab
	config.DB.First(&alat, id)
	config.DB.Find(&labs)

	if r.Method == http.MethodPost {
		alat.NamaAlat = r.FormValue("nama_alat")
		alat.Jumlah = Atoi(r.FormValue("jumlah"))
		alat.Kondisi = r.FormValue("kondisi")
		alat.LabID = uint(Atoi(r.FormValue("lab_id")))
		config.DB.Save(&alat)
		http.Redirect(w, r, "/alat", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Title":  "Edit Alat",
		"Action": "/alat/edit/" + id,
		"Alat":   alat,
		"Labs":   labs,
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/alat/form.html"))
	tmpl.Execute(w, data)
}

func AlatDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/alat/delete/"):len(r.URL.Path)]
	var alat config.Alat
	config.DB.Delete(&alat, id)
	http.Redirect(w, r, "/alat", http.StatusSeeOther)
}

func AlatPDF(w http.ResponseWriter, r *http.Request) {
	var alats []config.Alat
	config.DB.Preload("Lab").Find(&alats)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Laporan Data Alat")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	pdf.CellFormat(10, 8, "ID", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 8, "Nama Alat", "1", 0, "", false, 0, "")
	pdf.CellFormat(20, 8, "Jumlah", "1", 0, "", false, 0, "")
	pdf.CellFormat(30, 8, "Kondisi", "1", 0, "", false, 0, "")
	pdf.CellFormat(50, 8, "Lab", "1", 1, "", false, 0, "")

	for _, a := range alats {
		pdf.CellFormat(10, 8, fmt.Sprint(a.ID), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 8, a.NamaAlat, "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprint(a.Jumlah), "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 8, a.Kondisi, "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 8, a.Lab.NamaLab, "1", 1, "", false, 0, "")
	}

	_ = pdf.Output(w)
}

func AlatExcel(w http.ResponseWriter, r *http.Request) {
	var alats []config.Alat
	config.DB.Preload("Lab").Find(&alats)

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	headers := []string{"ID", "Nama Alat", "Jumlah", "Kondisi", "Lab"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, a := range alats {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), a.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), a.NamaAlat)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), a.Jumlah)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), a.Kondisi)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), a.Lab.NamaLab)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=alat.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	_ = f.Write(w)
}
