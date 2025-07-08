package repository

import (
	"context"

	"github.com/kaosdisablon/entity"
)

func (r *repository) GetPembayaran(ctx context.Context) ([]entity.Pembayaran, error) {
	var metode_pembayaran []entity.Pembayaran
	query := "SELECT id,nama FROM metode_pembayaran"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return metode_pembayaran, err
	}
	defer rows.Close()

	for rows.Next() {
		var pembayaran entity.Pembayaran

		if err := rows.Scan(&pembayaran.Id, &pembayaran.Nama); err != nil {
			return metode_pembayaran, err
		}
		metode_pembayaran = append(metode_pembayaran, pembayaran)
	}
	return metode_pembayaran, nil
}
