package services

import (
	"errors"
	"os"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type PaymentServices interface {
	GetTotalAmount(orderID int) (float64, error)
	CreateInvoice(orderID int) (map[string]interface{}, error)
	CheckPaymentStatus(invoiceID string) (string, error)
	CancelOrder(invoiceID string) error
	GetPaymentByInvoiceID(invoiceID string) (*entities.Payment, error)
	// GetPaidOrders() ([]entities.Payment, error)
	GetPaymentsByUserID(userID int) ([]entities.Payment, error)
}

type paymentServices struct {
	paymentRepo repositories.PaymentsRepo
}

func NewPaymentServices(paymentRepo repositories.PaymentsRepo) PaymentServices {
	return &paymentServices{paymentRepo: paymentRepo}
}

func (s *paymentServices) GetTotalAmount(orderID int) (float64, error) {
	jumlah, err := s.paymentRepo.GetOrderTotalAmount(orderID)
	if err != nil {
		return 0, err
	}
	return jumlah, err
}

func (s *paymentServices) CreateInvoice(orderID int) (map[string]interface{}, error) {
	if orderID < 0 {
		return nil, errors.New("orderID tidak valid")
	}
	jumlah, err := s.paymentRepo.GetOrderTotalAmount(orderID)
	if err != nil {
		return nil, errors.New("gagal menghitung total amount: " + err.Error())
	}
	orderid := strconv.Itoa(orderID)

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")
	if xendit.Opt.SecretKey == "" {
		return nil, errors.New("xendit Secret Key tidak ditemukan")
	}
	inv, err := invoice.Create(&invoice.CreateParams{
		ExternalID:  "order_" + orderid,
		Amount:      jumlah,
		Description: "Pemabayaran untuk order " + orderid,
	})
	_, Saveerr := s.paymentRepo.SavePaymentStatus(inv.ID, "PENDING", orderID, jumlah, inv.InvoiceURL)
	if Saveerr != nil {
		return nil, errors.New("gagal menyimpan status pembayaran: " + Saveerr.Error())
	}
	return map[string]interface{}{
		"invoice_id":  inv.ID,
		"external_id": inv.ExternalID,
		"order_id":    orderID,
		"status":      inv.Status,
		"jumlah":      jumlah,
		"invoice_url": inv.InvoiceURL,
	}, nil
}

func (s *paymentServices) CheckPaymentStatus(invoiceID string) (string, error) {
	if invoiceID == "" {
		return "", errors.New("invoiceID tidak boleh kosong")
	}
	params := &invoice.GetParams{ID: invoiceID}
	inv, err := invoice.Get(params)
	if err != nil {
		return "", errors.New("gagal mendapatkan status pembayaran dari Xendit: " + err.Error())
	}
	currentstatus, _ := s.paymentRepo.GetPaymentStatus(invoiceID)
	if currentstatus == "CANCEL" {
		return currentstatus, nil
	}
	s.paymentRepo.UpdatePaymentStatus(invoiceID, inv.Status)
	if inv.Status == "PAID" || inv.Status == "SETTLED" {
		err := s.paymentRepo.UpdateBarangStatusAsync(invoiceID)
		if err != nil {
			return inv.Status, errors.New("gagal memperbarui status barang: " + err.Error())
		}
	}
	return inv.Status, nil
}

func (s *paymentServices) CancelOrder(invoiceID string) error {
	params := &invoice.ExpireParams{
		ID: invoiceID,
	}
	_, err := invoice.Expire(params)
	if err != nil {
		return errors.New("gagal membatalkan invoice di Xendit: " + err.Error())
	}
	s.paymentRepo.CancelOrder(invoiceID)
	return nil
}

func (s *paymentServices) GetPaymentByInvoiceID(invoiceID string) (*entities.Payment, error) {
	payment, err := s.paymentRepo.GetPaymentByInvoiceID(invoiceID)
	if err != nil {
		return nil, errors.New("payment with the given invoice ID not found")
	}
	return payment, nil

}

// func (s *paymentServices) GetPaidOrders() ([]entities.Payment, error) {
// 	payments, err := s.paymentRepo.GetPaidOrders()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return payments, nil
// }

func (s *paymentServices) GetPaymentsByUserID(userID int) ([]entities.Payment, error) {
	return s.paymentRepo.GetPaymentsByUserID(userID)
}
