package repositories

import (
	"errors"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/xendit/xendit-go/invoice"
	"gorm.io/gorm"
)

type PaymentsRepo interface {
	SavePaymentStatus(invoiceID, status string, orderID int, amount float64, paymentURL string) (*entities.Payment, error)
	GetOrderTotalAmount(orderID int) (float64, error)
	UpdatePaymentStatus(invoiceID, status string) error
	CancelOrder(invoiceID string) error
	UpdateBarangStatusAsync(invoiceID string) error
	GetPaymentByInvoiceID(invoiceID string) (*entities.Payment, error)
	// GetPaidOrders() ([]entities.Payment, error)
	GetPaymentStatus(invoiceID string) (string, error)
	GetAllPayments() ([]entities.Payment, error)
	GetPaymentsByUserID(userID int) ([]entities.Payment, error)
}

type paymentsRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentsRepo {
	return &paymentsRepo{db: db}
}

func (r *paymentsRepo) SavePaymentStatus(invoiceID, status string, orderID int, amount float64, paymentURL string) (*entities.Payment, error) {
	payment := entities.Payment{
		InvoiceID:  invoiceID,
		Status:     status,
		OrderID:    orderID,
		Jumlah:     amount,
		PaymentUrl: paymentURL,
	}
	if err := r.db.Create(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentsRepo) GetOrderTotalAmount(orderID int) (float64, error) {
	var order entities.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		return 0, errors.New("order tidak ditemukan")
	}
	totalAmount := order.Total
	return totalAmount, nil
}

func (r *paymentsRepo) UpdatePaymentStatus(invoiceID, status string) error {
	var payment entities.Payment
	if err := r.db.Where("invoice_id = ?", invoiceID).First(&payment).Error; err != nil {
		return errors.New("pembayaran tidak ditemukan")
	}
	payment.Status = status
	if err := r.db.Save(&payment).Error; err != nil {
		return errors.New("gagal memperbarui status pembayaran")
	}
	return nil
}

func (r *paymentsRepo) CancelOrder(invoiceID string) error {
	params := &invoice.ExpireParams{
		ID: invoiceID,
	}
	_, err := invoice.Expire(params)
	if err != nil {
		return errors.New("gagal membatalkan invoice di Xendit: " + err.Error())
	}
	var payment entities.Payment
	if err := r.db.Where("invoice_id = ?", invoiceID).First(&payment).Error; err != nil {
		return errors.New("invoice dengan ID yang diberikan tidak ditemukan")
	}
	payment.Status = "CANCEL"
	if err := r.db.Save(&payment).Error; err != nil {
		return errors.New("gagal membatalkan pembayaran")
	}
	return nil
}

func (r *paymentsRepo) UpdateBarangStatusAsync(invoiceID string) error {
	var payment entities.Payment
	if err := r.db.Where("invoice_id = ?", invoiceID).First(&payment).Error; err != nil {
		return errors.New("pembayaran dengan ID yang diberikan tidak ditemukan")
	}
	if payment.Status != "PAID" && payment.Status != "SETTLED" {
		return errors.New("status pembayaran tidak PAID atau SETTLED")
	}

	payment.StatusBarang = "DIKIRIM"
	if err := r.db.Save(&payment).Error; err != nil {
		return errors.New("gagal memperbarui status barang menjadi DIKIRIM")
	}

	go func() {
		time.Sleep(2 * time.Minute)
		payment.StatusBarang = "SELESAI"
		r.db.Save(&payment)
	}()

	return nil
}

func (r *paymentsRepo) GetPaymentByInvoiceID(invoiceID string) (*entities.Payment, error) {
	var payment entities.Payment
	err := r.db.Preload("Order.Details.Product").
		Preload("Order.Details.User").
		Where("invoice_id = ?", invoiceID).
		First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invoice not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *paymentsRepo) GetPaymentStatus(invoiceID string) (string, error) {
	var payment entities.Payment
	if err := r.db.Where("invoice_id = ?", invoiceID).First(&payment).Error; err != nil {
		return "", errors.New("pembayaran tidak ditemukan")
	}
	return payment.Status, nil
}

// func (r *paymentsRepo) GetPaidOrders() ([]entities.Payment, error) {
// 	var payments []entities.Payment
// 	err := r.db.
// 		Preload("Order.Details.Product").
// 		Preload("Order.User").
// 		Where("status = ?", "PAID", "").
// 		Find(&payments).Error

// 	if err != nil {
// 		return nil, err
// 	}
// 	return payments, nil
// }

// perbaikan dari line 129
func (r *paymentsRepo) GetAllPayments() ([]entities.Payment, error) {
	var payments []entities.Payment
	if err := r.db.Preload("Order").
		Preload("Order.Details").
		Preload("Order.Details.Product").
		Preload("Order.Details.User").
		Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *paymentsRepo) GetPaymentsByUserID(userID int) ([]entities.Payment, error) {
	var payments []entities.Payment
	if err := r.db.Preload("Order").
		Preload("Order.Details").
		Preload("Order.Details.Product").
		Preload("Order.Details.User").Joins("JOIN orders o ON payments.order_id = o.id").Where("o.user_id = ?", userID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
