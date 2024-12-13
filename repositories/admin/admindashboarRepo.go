package admin

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type AdminPaymentRepository interface {
	GetAdminTotalPendapatanBulanIni() (float64, error)
	GetAdminJumlahPesananByStatus(status []string) (int64, error)
	GetTotalProdukInStock() (int, error)
	GetJumlahPesananByStatus(status string) (int64, error)
	GetJumlahPesananDikirim() (int64, error)
	GetJumlahPesananDiterima() (int64, error)
	GetTotalPendapatan() ([]entities.TotalPendapatan, error)
}

type adminPaymentRepository struct {
	db *gorm.DB
}

func NewAdminPaymentRepository(db *gorm.DB) AdminPaymentRepository {
	return &adminPaymentRepository{db: db}
}

func (r *adminPaymentRepository) GetAdminTotalPendapatanBulanIni() (float64, error) {
	var totalPendapatan float64
	err := r.db.Model(&entities.Payment{}).Where("status = ? ", "PAID").Select("SUM(jumlah)").Scan(&totalPendapatan).Error
	if err != nil {
		return 0, err
	}
	return totalPendapatan, nil
}

func (r *adminPaymentRepository) GetAdminJumlahPesananByStatus(status []string) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Order{}).
		Joins("JOIN payments ON orders.id = payments.order_id").Where("payments.status IN (?)", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *adminPaymentRepository) GetTotalProdukInStock() (int, error) {
	var totalProduk int
	err := r.db.Model(&entities.Product{}).Select("SUM(stok)").Scan(&totalProduk).Error
	if err != nil {
		return 0, err
	}
	return totalProduk, nil
}

func (r *adminPaymentRepository) GetJumlahPesananByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Payment{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *adminPaymentRepository) GetJumlahPesananDikirim() (int64, error) {
	var count int64
	err := r.db.Model(&entities.Payment{}).Where("status_barang = ?", "DIKIRIM").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *adminPaymentRepository) GetJumlahPesananDiterima() (int64, error) {
	var count int64
	err := r.db.Model(&entities.Payment{}).
		Where("status_barang = ?", "SELESAI").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *adminPaymentRepository) GetTotalPendapatan() ([]entities.TotalPendapatan, error) {
	var pendapatan []entities.TotalPendapatan
	err := r.db.Find(&pendapatan).Error
	return pendapatan, err
}

//artikel
