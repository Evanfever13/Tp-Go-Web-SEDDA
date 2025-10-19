package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func main() {
	type Produits struct {
		Img              string
		ImgAlt           string
		Name             string
		Prix             int
		PrixReduit       int
		PourcentageReduc int
		Description      string
		Taille           string
		IsReduc          bool
		Id               int
	}
	Pull1 := Produits{"./static/img/products/16A.webp", "Image du pull 16A", "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", 146, 0, 0, "Un pull sympa", "XS/S/M/L/XL", false, 0}
	Pull2 := Produits{"./static/img/products/18A.webp", "Image du pull 18A", "PALACE PULL A CAPUCHE UNISEXE GOTHIC", 150, 0, 0, "Un pull sympa", "XS/S/M/L/XL", false, 1}
	Pull3 := Produits{"./static/img/products/19A.webp", "Image du pull 19A", "PALACE PULL A CAPUCHE UNISEXE VERDAD", 100, 75, 25, "Un pull sympa", "XS/S/M/L/XL", true, 2}
	Pull4 := Produits{"./static/img/products/21A.webp", "Image du pull 21A", "PALACE PULL UNISEXE MARINE", 132, 0, 0, "Un pull sympa", "XS/S/M/L/XL", false, 3}
	Pull5 := Produits{"./static/img/products/22A.webp", "Image du pull 22A", "PALACE PULL UNISEXE NOIR", 92, 0, 0, "Un pull sympa", "XS/S/M/L/XL", false, 4}
	Pantalon1 := Produits{"./static/img/products/33B.webp", "Image du pull 33A", "PANTALON UNISEXE GOTHIC", 112, 0, 0, "Un pantalon sympa", "XS/S/M/L/XL", false, 5}
	ListProduits := []Produits{Pull1, Pull2, Pull3, Pull4, Pull5, Pantalon1}

	temp, errtemp := template.ParseGlob("./templates/*.html")
	if errtemp != nil {
		fmt.Println(errtemp)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Home", ListProduits)

	})
	http.HandleFunc("/ajouter", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Impossible de lire le formulaire", http.StatusBadRequest)
				return
			}

			priceStr := r.FormValue("price")
			price, err := strconv.Atoi(priceStr)

			if err != nil {
				price = 0
			}

			priceRedStr := r.FormValue("priceRed")
			priceRed, err := strconv.Atoi(priceRedStr)
			if err != nil {
				priceRed = 0
			}

			pourcentageStr := r.FormValue("Pourcentage")
			pourcentage, err := strconv.Atoi(pourcentageStr)
			if err != nil {
				pourcentage = 0
			}

			isReduc := priceRed > 0 || pourcentage > 0

			nextID := 0

			for _, p := range ListProduits {
				if p.Id >= nextID {
					nextID = p.Id + 1
				}
			}

			NewProduit := Produits{
				Img:              r.FormValue("image"),
				ImgAlt:           r.FormValue("imageAlt"),
				Name:             r.FormValue("name"),
				Prix:             price,
				PrixReduit:       priceRed,
				PourcentageReduc: pourcentage,
				Description:      r.FormValue("description"),
				Taille:           r.FormValue("sizes"),
				IsReduc:          isReduc,
				Id:               nextID,
			}

			ListProduits = append(ListProduits, NewProduit)
		}

		temp.ExecuteTemplate(w, "Ajouter", ListProduits)
		priceStr := r.FormValue("price")
		price, _ := strconv.Atoi(priceStr)

		priceRedStr := r.FormValue("priceRed")
		priceRed, _ := strconv.Atoi(priceRedStr)

		pourcentageStr := r.FormValue("Pourcentage")
		pourcentage, _ := strconv.Atoi(pourcentageStr)

		isReduc := priceRedStr != "" || pourcentageStr != ""

		NewProduit := Produits{
			Img:              r.FormValue("image"),
			ImgAlt:           r.FormValue("imageAlt"),
			Name:             r.FormValue("name"),
			Prix:             price,
			PrixReduit:       priceRed,
			PourcentageReduc: pourcentage,
			Description:      r.FormValue("description"),
			Taille:           r.FormValue("sizes"),
			IsReduc:          isReduc,
			Id:               len(ListProduits) + 1,
		}

		ListProduits = append(ListProduits, NewProduit)

		temp.ExecuteTemplate(w, "Ajouter", ListProduits)
	})
	http.HandleFunc("/produit", func(w http.ResponseWriter, r *http.Request) {
		idProduit := r.FormValue("id")
		produitId, err := strconv.Atoi(idProduit)
		if err != nil {
			http.Error(w, "Erreur: id du produit invalide", http.StatusBadRequest)
			return
		}

		for _, product := range ListProduits {
			if product.Id == produitId {
				temp.ExecuteTemplate(w, "Consulter", product)
				return
			}
		}

		http.Error(w, "Produit non trouvé", http.StatusNotFound)
	})

	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8000", nil)
}
