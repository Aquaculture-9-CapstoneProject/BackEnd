package admin

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type AdminFilterRepo interface {
	GetPaymentsByStatus(status string) ([]entities.Payment, error)
	GetPaymentsByStatusBarang(statusBarang string) ([]entities.Payment, error)
}

type adminFilterRepo struct {
	db *gorm.DB
}

func NewAdminFilterRepo(db *gorm.DB) AdminFilterRepo {
	return &adminFilterRepo{db: db}
}

func (pr *adminFilterRepo) GetPaymentsByStatus(status string) ([]entities.Payment, error) {
	var payments []entities.Payment
	err := pr.db.Preload("Order").Preload("Order.details").Where("status = ?", status).Find(&payments).Error
	if err != nil {
		return nil, errors.New("gagal mengambil data pembayaran: " + err.Error())
	}
	return payments, nil
}

func (pr *adminFilterRepo) GetPaymentsByStatusBarang(statusBarang string) ([]entities.Payment, error) {
	var payments []entities.Payment
	err := pr.db.Where("status_barang = ?", statusBarang).Find(&payments).Error
	if err != nil {
		return nil, errors.New("gagal mengambil data pembayaran: " + err.Error())
	}
	return payments, nil
}
