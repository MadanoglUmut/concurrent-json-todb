package main

import (
	"ReadProducts/models"
	"ReadProducts/pkg"
	"ReadProducts/pkg/psql"
	"ReadProducts/repositories"
	"ReadProducts/services"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Env Dosyası Okunamadi", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	var db = psql.Connect(host, user, password, name, port)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)

	files := []string{
		"../.././data/products_1mb.json",
		"../.././data/products_25mb.json",
		"../.././data/products_75mb.json",
		"../.././data/products_150mb.json",
	}

	var wg sync.WaitGroup
	sm := make(chan struct{}, 2)

	const batchSize = 500
	productChan := make(chan []models.CreateProduct, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := productService.İmportProducts(productChan); err != nil {
			log.Fatal("Batch işleme hatası:", err)
		}
	}()

	var fileWg sync.WaitGroup
	for _, file := range files {
		fileWg.Add(1)
		go func(f string) {
			defer fileWg.Done()
			sm <- struct{}{}
			log.Printf("Başladı: %s", f)
			//defer func() { <-sm }()
			err := pkg.LoadProductsFromFile(f, batchSize, productChan)
			if err != nil {
				log.Printf("Dosya okuma hatası (%s): %v", f, err)
			} else {
				log.Printf("Dosya başarıyla işlendi: %s", f)
			}
			<-sm
		}(file)
	}

	go func() {
		fileWg.Wait()
		close(productChan)
	}()

	wg.Wait()

	log.Println("Tüm ürünler başarıyla aktarıldı")

}

//ZATEN DOSYADA OLAN ÜRÜNLER İÇİN DB YE GİTMESİN
//PRODUTC'IN SADECE BELİRLİ BİR FİELDI DEĞİŞTİYSE UPDATE YAPALIM
