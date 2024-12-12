package admin

import (
	"errors"

	"gorm.io/gorm"
)

type AdminPesananRepo interface {
	GetDetailedOrders() ([]map[string]interface{}, error)
}

type adminPesananRepo struct {
	db *gorm.DB
}

func NewPesananRepo(db *gorm.DB) AdminPesananRepo {
	return &adminPesananRepo{db: db}
}

func (pr *adminPesananRepo) GetDetailedOrders() ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := pr.db.Table("payments").
		Select(`
			orders.id AS order_id, 
			users.nama_lengkap AS namapengguna, 
			products.name AS produk, 
			orders.created_at AS tanggaldanwaktu, 
			users.alamat AS alamat, 
			payments.amount AS nominal, 
			payments.status_barang AS status`).
		Joins("JOIN orders ON orders.id = payments.order_id").
		Joins("JOIN users ON users.id = orders.user_id").
		Joins("JOIN order_details ON order_details.order_id = orders.id").
		Joins("JOIN products ON products.id = order_details.product_id").
		Scan(&results).Error

	if err != nil {
		return nil, errors.New("gagal mengambil data pesanan: " + err.Error())
	}
	return results, nil
}
