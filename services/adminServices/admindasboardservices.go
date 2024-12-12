package adminservices

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminPaymentService interface {
	HitungAdminPendapatanBulanIni() (float64, error)
	GetAdminJumlahPesananBulanIni(status []string) (int64, error)
	GetTotalProduk() (int, error)
	GetJumlahPesananByStatus(status string) (int64, error)
	GetJumlahPesananDikirim() (int64, error)
	GetJumlahPesananDiterima() (int64, error)
	GetTotalPendapatBulan() ([]entities.TotalPendapatan, error)
	GetJumlahArtikel() (int64, error)
	GetProdukDenganKategoriStokTerbanyak() ([]entities.Product, error)
}

type adminPaymentService struct {
	adminPaymentRepo admin.AdminPaymentRepository
}

func NewAdminPaymentService(adminPaymentRepo admin.AdminPaymentRepository) AdminPaymentService {
	return &adminPaymentService{adminPaymentRepo: adminPaymentRepo}
}

func (s *adminPaymentService) HitungAdminPendapatanBulanIni() (float64, error) {
	total, err := s.adminPaymentRepo.GetAdminTotalPendapatanBulanIni()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *adminPaymentService) GetAdminJumlahPesananBulanIni(status []string) (int64, error) {
	return s.adminPaymentRepo.GetAdminJumlahPesananByStatus(status)
}

func (s *adminPaymentService) GetTotalProduk() (int, error) {
	totalProduk, err := s.adminPaymentRepo.GetTotalProdukInStock()
	if err != nil {
		return 0, err
	}
	return totalProduk, nil
}

func (s *adminPaymentService) GetJumlahPesananByStatus(status string) (int64, error) {
	count, err := s.adminPaymentRepo.GetJumlahPesananByStatus(status)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *adminPaymentService) GetJumlahPesananDikirim() (int64, error) {
	count, err := s.adminPaymentRepo.GetJumlahPesananDikirim()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *adminPaymentService) GetJumlahPesananDiterima() (int64, error) {
	count, err := s.adminPaymentRepo.GetJumlahPesananDiterima()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *adminPaymentService) GetTotalPendapatBulan() ([]entities.TotalPendapatan, error) {
	return s.adminPaymentRepo.GetTotalPendapatan()
}

func (s *adminPaymentService) GetJumlahArtikel() (int64, error) {
	return s.adminPaymentRepo.GetJumlahArtikel()
}

func (s *adminPaymentService) GetProdukDenganKategoriStokTerbanyak() ([]entities.Product, error) {
	return s.adminPaymentRepo.GetProdukDenganKategoriStokTerbanyak()
}
