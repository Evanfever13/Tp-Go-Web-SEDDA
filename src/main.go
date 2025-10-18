package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	type Produits struct {
		Img    string
		ImgAlt string
		Name   string
		Price  string
	}
	Pull1 := Produits{"./static/img/products/16A.webp", "Image du pull 16A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", "146€"}
	Pull2 := Produits{"./static/img/products/18A.webp", "Image du pull 18A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", "146€"}
	Pull3 := Produits{"./static/img/products/19A.webp", "Image du pull 19A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", "146€"}
	Pull4 := Produits{"./static/img/products/21A.webp", "Image du pull 21A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", "146€"}
	Pull5 := Produits{"./static/img/products/22A.webp", "Image du pull 22A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", "146€"}
	Pull6 := Produits{"./static/img/products/33A.webp", "Image du pull 33A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", "146€"}
	ListProduits := []Produits{Pull1, Pull2, Pull3, Pull4, Pull5, Pull6}

	temp, errtemp := template.ParseGlob("./templates/*.html")
	if errtemp != nil {
		fmt.Println(errtemp)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Home", ListProduits)

	})

	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8000", nil)
}
