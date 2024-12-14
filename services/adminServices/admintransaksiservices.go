package adminservices

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminTransaksiService interface {
	GetPaymentDetails(page, perPage int) (map[string]interface{}, error)
	DeletePaymentByID(id int) error
}

type adminTransaksiService struct {
	repoadminTransaksi admin.AdminTransaksiRepo
}

func NewAdminTransaksiServices(repoadminTransaksi admin.AdminTransaksiRepo) AdminTransaksiService {
	return &adminTransaksiService{repoadminTransaksi: repoadminTransaksi}
}

func (ps *adminTransaksiService) GetPaymentDetails(page, perPage int) (map[string]interface{}, error) {
	details, totalItems, err := ps.repoadminTransaksi.GetPaymentDetails(page, perPage)
	if err != nil {
		return nil, errors.New("gagal mendapatkan detail pembayaran: " + err.Error())
	}

	// Hitung pagination
	currentPage := page
	totalPages := int((totalItems + int64(perPage) - 1) / int64(perPage))

	// Struktur hasil dengan pagination
	response := map[string]interface{}{
		"data": details,
		"pagination": map[string]interface{}{
			"CurrentPage": currentPage,
			"TotalPages":  totalPages,
			"TotalItems":  totalItems,
		},
	}

	return response, nil
}

func (ps *adminTransaksiService) DeletePaymentByID(id int) error {
	err := ps.repoadminTransaksi.DeletePaymentByID(id)
	if err != nil {
		return errors.New("gagal menghapus payment: " + err.Error())
	}
	return nil
}
