package adminservices

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductExportService interface {
	ExportToCSV() error
}

type productExportService struct {
	repo repositories.ProdukIkanRepo
}

func NewProducExportService(repo repositories.ProdukIkanRepo) ProductExportService {
	return &productExportService{repo: repo}
}

func (s *productExportService) ExportToCSV() error {
	products, err := s.repo.GetSemuaProduk()
	if err != nil {
		return err
	}
	file, err := os.Create("products.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Gambar", "Nama", "Deskripsi", "Keunggulan", "Harga", "Variasi", "Kategori", "Kota Asal", "Rating", "Stok", "Total Review", "Status", "Barang Terjual", "Tips Penyimpanan"})

	for _, product := range products {
		record := []string{
			strconv.Itoa(int(product.ID)),
			product.Gambar,
			product.Nama,
			product.Deskripsi,
			product.Keunggulan,
			strconv.FormatFloat(product.Harga, 'f', 3, 64),
			product.Variasi,
			product.Kategori,
			product.KotaAsal,
			strconv.FormatFloat(product.Rating, 'f', 2, 64),
			strconv.Itoa(product.Stok),
			strconv.Itoa(product.TotalReview),
			product.Status,
			strconv.Itoa(product.Terjual),
			product.TipsPenyimpanan,
		}
		writer.Write(record)
	}
	return nil
}
