package admin

import (
	"errors"

	"gorm.io/gorm"
)

type AdminPesananRepo interface {
	GetDetailedOrders(page, perPage int) ([]map[string]interface{}, int64, error)
}

type adminPesananRepo struct {
	db *gorm.DB
}

func NewPesananRepo(db *gorm.DB) AdminPesananRepo {
	return &adminPesananRepo{db: db}
}

func (pr *adminPesananRepo) GetDetailedOrders(page, perPage int) ([]map[string]interface{}, int64, error) {
	var results []map[string]interface{}
	var totalItems int64

	// Menghitung total item
	pr.db.Table("payments").
		Joins("JOIN orders ON orders.id = payments.order_id").
		Joins("JOIN users ON users.id = orders.user_id").
		Joins("JOIN order_details ON order_details.order_id = orders.id").
		Joins("JOIN products ON products.id = order_details.product_id").
		Count(&totalItems)

	// Membatasi data yang diambil sesuai pagination
	offset := (page - 1) * perPage
	err := pr.db.Table("payments").
		Select(`
			payments.id AS payment_id,      
			orders.id AS order_id, 
			users.nama_lengkap AS namapengguna, 
			products.variasi AS variasi,
			products.nama AS produk, 
			order_details.kuantitas AS kuantitas,   
			orders.created_at AS tanggaldanwaktu, 
			users.alamat AS alamat, 
			payments.jumlah AS nominal, 
			payments.status_barang AS status`).
		Joins("JOIN orders ON orders.id = payments.order_id").
		Joins("JOIN users ON users.id = orders.user_id").
		Joins("JOIN order_details ON order_details.order_id = orders.id").
		Joins("JOIN products ON products.id = order_details.product_id").
		Offset(offset).Limit(perPage).
		Scan(&results).Error

	if err != nil {
		return nil, 0, errors.New("gagal mengambil data pesanan: " + err.Error())
	}
	return results, totalItems, nil
}
