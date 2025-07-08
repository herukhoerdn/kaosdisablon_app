package repository

import (
	"context"
	"fmt"

	"github.com/kaosdisablon/entity"
)

func (r *repository) InsertCheckout(ctx context.Context, checkout entity.Checkout) (int64, error) {
	query := "INSERT INTO checkout (user_id,produk_id,kuantiti,total_harga,metode_pembayaran,tanggal_order,status,desain_id)VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, checkout.UserId, checkout.ProdukId, checkout.Kuantiti, checkout.TotalHarga, checkout.MetodePembayaran, checkout.TanggalOrder, checkout.Status, checkout.DesainId).Scan(&checkout.Id)
	if err != nil {
		return 0, err
	}
	return checkout.Id, nil
}
func (r *repository) GetCheckout(ctx context.Context) ([]entity.Checkout, error) {
	var checkout []entity.Checkout
	query := "SELECT id,user_id,produk_id,kuantiti,total_harga,metode_pembayaran,tanggal_order,status,desain_id FROM checkout"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return checkout, err
	}
	defer rows.Close()

	for rows.Next() {
		var chec entity.Checkout

		if err := rows.Scan(&chec.Id, &chec.UserId, &chec.ProdukId, &chec.Kuantiti, &chec.TotalHarga, &chec.MetodePembayaran, &chec.TanggalOrder, &chec.Status, &chec.DesainId); err != nil {
			return checkout, err
		}
		checkout = append(checkout, chec)
	}
	return checkout, nil
}
func (r *repository) UpdateCheckout(ctx context.Context, checkout entity.Checkout) (int64, error) {
	query := "UPDATE checkout SET user_id=$2,produk_id=$3,kuantiti=$4,total_harga=$5,metode_pembayaran=$6,tanggal_order=$7,status=$8,desain_id=$9 WHERE id=$1 RETURNING id "
	err := r.db.QueryRowContext(ctx, query, checkout.Id, checkout.UserId, checkout.ProdukId, checkout.Kuantiti, checkout.TotalHarga, checkout.MetodePembayaran, checkout.TanggalOrder, checkout.Status, checkout.DesainId).Scan(&checkout.Id)

	if err != nil {
		return 0, err
	}
	return checkout.Id, nil
}
func (r *repository) DeleteCheckout(ctx context.Context, id int64) error {
	query := "DELETE FROM checkout WHERE id =$1"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
func (r *repository) FindCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error) {
	query := `SELECT id, user_id, produk_id, kuantiti, total_harga, metode_pembayaran, tanggal_order, status, desain_id
			  FROM checkout WHERE desain_id = $1`
	var c entity.Checkout
	err := r.db.QueryRowContext(ctx, query, desainId).Scan(
		&c.Id, &c.UserId, &c.ProdukId, &c.Kuantiti, &c.TotalHarga,
		&c.MetodePembayaran, &c.TanggalOrder, &c.Status, &c.DesainId,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
