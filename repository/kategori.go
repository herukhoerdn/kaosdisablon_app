package repository

import (
	"context"
	"fmt"

	"github.com/kaosdisablon/entity"
)

func (r *repository) InsertKategori(ctx context.Context, kategori entity.Kategori) (int64, error) {
	query := "INSERT INTO kategori(nama) VALUES($1) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, kategori.Nama).Scan(&kategori.Id)
	if err != nil {
		return 0, err
	}
	return kategori.Id, nil
}

func (r *repository) GetKategori(ctx context.Context) ([]entity.Kategori, error) {
	var kategori []entity.Kategori
	query := "SELECT id,nama FROM kategori"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return kategori, err
	}
	defer rows.Close()

	for rows.Next() {
		var kat entity.Kategori

		if err := rows.Scan(&kat.Id, &kat.Nama); err != nil {
			return kategori, err
		}

		kategori = append(kategori, kat)
	}
	return kategori, nil
}

func (r *repository) Updatekategori(ctx context.Context, kategori entity.Kategori) (int64, error) {
	query := "UPDATE kategori SET nama=$2 WHERE id=$1 RETURNING id"
	err := r.db.QueryRowContext(ctx, query, kategori.Id, kategori.Nama).Scan(&kategori.Id)
	if err != nil {
		return 0, err
	}
	return kategori.Id, nil
}

func (r *repository) DeleteKategori(ctx context.Context, id int64) error {
	query := "DELETE FROM kategori WHERE id =$1"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
