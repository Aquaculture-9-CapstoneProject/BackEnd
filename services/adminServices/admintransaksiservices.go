package adminservices

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminTransaksiService interface {
	GetPaymentDetails() ([]map[string]interface{}, error)
	DeletePaymentByID(id int) error
}

type adminTransaksiService struct {
	repoadminTransaksi admin.AdminTransaksiRepo
}

func NewAdminTransaksiServices(repoadminTransaksi admin.AdminTransaksiRepo) AdminTransaksiService {
	return &adminTransaksiService{repoadminTransaksi: repoadminTransaksi}
}

func (ps *adminTransaksiService) GetPaymentDetails() ([]map[string]interface{}, error) {
	details, err := ps.repoadminTransaksi.GetPaymentDetails()
	if err != nil {
		return nil, errors.New("gagal mendapatkan detail pembayaran: " + err.Error())
	}
	return details, nil
}

func (ps *adminTransaksiService) DeletePaymentByID(id int) error {
	err := ps.repoadminTransaksi.DeletePaymentByID(id)
	if err != nil {
		return errors.New("gagal menghapus payment: " + err.Error())
	}
	return nil
}
