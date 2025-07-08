package repository

import (
	"context"
	"fmt"

	"github.com/kaosdisablon/entity"
)

func (r *repository) CreateInvoice(ctx context.Context, invoice entity.Invoice) (int64, error) {
	fmt.Printf("[DEBUG] Trying insert invoice with checkout_id = %d\n", invoice.CheckoutId)
	query := `INSERT INTO invoice (checkout_id, user_id, total_harga, tanggal_buat, file_invoice)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64

	// fmt.Printf(">> INSERT invoice debug: checkout_id=%d user_id=%d total=%.2f\n", invoice.CheckoutId, invoice.UserId, invoice.TotalHarga)
	err := r.db.QueryRowContext(ctx, query,
		invoice.CheckoutId,
		invoice.UserId,
		invoice.TotalHarga,
		invoice.TanggalBuat,
		invoice.FileInvoice,
	).Scan(&id)
	return id, err
}

func (r *repository) GetInvoicesByUser(ctx context.Context, userId int64) ([]entity.Invoice, error) {
	query := `SELECT id, checkout_id, user_id, total_harga, tanggal_buat, file_invoice FROM invoice WHERE user_id = $1 AND file_invoice IS NOT NULL AND file_invoice <> ''
			  AND total_harga > 0 AND checkout_id > 0`
	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []entity.Invoice
	for rows.Next() {
		var inv entity.Invoice
		err := rows.Scan(&inv.Id, &inv.CheckoutId, &inv.UserId, &inv.TotalHarga, &inv.TanggalBuat, &inv.FileInvoice)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, inv)
	}

	return invoices, nil
}
func (r *repository) GetCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error) {
	query := `SELECT id, user_id, total_harga FROM checkout WHERE desain_id = $1`
	var c entity.Checkout
	err := r.db.QueryRowContext(ctx, query, desainId).Scan(&c.Id, &c.UserId, &c.TotalHarga)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *repository) UpdateInvoiceFilePath(ctx context.Context, invoiceID int64, fileName string) error {
	query := `UPDATE invoice SET file_invoice=$1 WHERE id=$2`
	_, err := r.db.ExecContext(ctx, query, fileName, invoiceID)
	return err
}

