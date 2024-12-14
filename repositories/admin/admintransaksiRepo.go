package admin

import (
	"errors"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type AdminTransaksiRepo interface {
	GetPaymentDetails(limit, offset int) ([]map[string]interface{}, error)
	DeletePaymentByID(id int) error
}

type adminTransaksiRepo struct {
	db *gorm.DB
}

func NewAdminTransaksiRepo(db *gorm.DB) AdminTransaksiRepo {
	return &adminTransaksiRepo{db: db}
}

func (ar *adminTransaksiRepo) GetPaymentDetails(limit, offset int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := ar.db.Table("payments").
		Select("payments.id, payments.invoice_id as id_pesanan, orders.metode_pembayaran, payments.status, orders.created_at").
		Joins("JOIN orders ON orders.id = payments.order_id").
		Limit(limit).
		Offset(offset).
		Scan(&results).Error

	if err != nil {
		return nil, errors.New("gagal mengambil detail pembayaran: " + err.Error())
	}

	for _, result := range results {
		if createdAt, ok := result["created_at"].(time.Time); ok {
			result["created_at"] = createdAt.Format("2006-01-02 15:04:05")
		}
	}

	return results, nil
}

func (ar *adminTransaksiRepo) DeletePaymentByID(id int) error {
	err := ar.db.Delete(&entities.Payment{}, id).Error
	if err != nil {
		return errors.New("gagal menghapus payment: " + err.Error())
	}
	return nil
}
