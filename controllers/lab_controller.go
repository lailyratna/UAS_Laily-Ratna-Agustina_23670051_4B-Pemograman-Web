package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"

	"sistem-peminjaman-lab/config"
)

func LabIndex(w http.ResponseWriter, r *http.Request) {
	var labs []config.Lab
	config.DB.Find(&labs)
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/lab/index.html"))
	tmpl.Execute(w, labs)
}

func LabCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		lab := config.Lab{
			NamaLab:   r.FormValue("nama_lab"),
			Lokasi:    r.FormValue("lokasi"),
			Kapasitas: Atoi(r.FormValue("kapasitas")),
		}
		config.DB.Create(&lab)
		http.Redirect(w, r, "/lab", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Title":  "Tambah Lab",
		"Action": "/lab/create",
		"Lab":    config.Lab{},
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/lab/form.html"))
	tmpl.Execute(w, data)
}

func LabEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/lab/edit/"):len(r.URL.Path)]
	var lab config.Lab
	config.DB.First(&lab, id)

	if r.Method == http.MethodPost {
		lab.NamaLab = r.FormValue("nama_lab")
		lab.Lokasi = r.FormValue("lokasi")
		lab.Kapasitas = Atoi(r.FormValue("kapasitas"))
		config.DB.Save(&lab)
		http.Redirect(w, r, "/lab", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Title":  "Edit Lab",
		"Action": "/lab/edit/" + id,
		"Lab":    lab,
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/lab/form.html"))
	tmpl.Execute(w, data)
}

func LabDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/lab/delete/"):len(r.URL.Path)]
	var lab config.Lab
	config.DB.Delete(&lab, id)
	http.Redirect(w, r, "/lab", http.StatusSeeOther)
}

func LabPDF(w http.ResponseWriter, r *http.Request) {
	var labs []config.Lab
	config.DB.Find(&labs)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Laporan Data Lab")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	pdf.CellFormat(10, 8, "ID", "1", 0, "", false, 0, "")
	pdf.CellFormat(50, 8, "Nama Lab", "1", 0, "", false, 0, "")
	pdf.CellFormat(50, 8, "Lokasi", "1", 0, "", false, 0, "")
	pdf.CellFormat(30, 8, "Kapasitas", "1", 1, "", false, 0, "")

	for _, l := range labs {
		pdf.CellFormat(10, 8, fmt.Sprint(l.ID), "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 8, l.NamaLab, "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 8, l.Lokasi, "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprint(l.Kapasitas), "1", 1, "", false, 0, "")
	}

	_ = pdf.Output(w)
}

func LabExcel(w http.ResponseWriter, r *http.Request) {
	var labs []config.Lab
	config.DB.Find(&labs)

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	headers := []string{"ID", "Nama Lab", "Lokasi", "Kapasitas"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, l := range labs {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), l.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), l.NamaLab)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), l.Lokasi)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), l.Kapasitas)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=lab.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	_ = f.Write(w)
}
