package admin

import (
	"errors"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type AdminTransaksiRepo interface {
	GetPaymentDetails(page, perPage int) ([]map[string]interface{}, int64, error)
	DeletePaymentByID(id int) error
}

type adminTransaksiRepo struct {
	db *gorm.DB
}

func NewAdminTransaksiRepo(db *gorm.DB) AdminTransaksiRepo {
	return &adminTransaksiRepo{db: db}
}

func (ar *adminTransaksiRepo) GetPaymentDetails(page, perPage int) ([]map[string]interface{}, int64, error) {
	var results []map[string]interface{}
	var totalItems int64

	// Hitung total items
	err := ar.db.Table("payments").
		Joins("JOIN orders ON orders.id = payments.order_id").
		Count(&totalItems).Error
	if err != nil {
		return nil, 0, errors.New("gagal menghitung total item: " + err.Error())
	}

	// Hitung offset berdasarkan page dan per_page
	offset := (page - 1) * perPage

	// Ambil data pembayaran dengan per_page dan offset
	err = ar.db.Table("payments").
		Select("payments.id, payments.invoice_id as id_pesanan, orders.metode_pembayaran, payments.status, orders.created_at").
		Joins("JOIN orders ON orders.id = payments.order_id").
		Limit(perPage).
		Offset(offset).
		Scan(&results).Error
	if err != nil {
		return nil, 0, errors.New("gagal mengambil detail pembayaran: " + err.Error())
	}

	// Format created_at
	for _, result := range results {
		if createdAt, ok := result["created_at"].(time.Time); ok {
			result["created_at"] = createdAt.Format("2006-01-02 15:04:05")
		}
	}

	return results, totalItems, nil
}

func (ar *adminTransaksiRepo) DeletePaymentByID(id int) error {
	err := ar.db.Delete(&entities.Payment{}, id).Error
	if err != nil {
		return errors.New("gagal menghapus payment: " + err.Error())
	}
	return nil
}
