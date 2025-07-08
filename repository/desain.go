package repository

import (
	"context"
	"fmt"

	"github.com/kaosdisablon/entity"
)

func (r *repository) InsertDesain(ctx context.Context, desain entity.Desain) (int64, error) {
	query := "INSERT INTO desain (user_id,file_desain,catatan,status,tanggal_upload) VALUES($1,$2,$3,$4,$5)RETURNING id"
	err := r.db.QueryRowContext(ctx, query, desain.UserId, desain.FileDesain, desain.Catatan, desain.Status, desain.TanggalUpload).Scan(&desain.Id)
	if err != nil {
		return 0, err
	}
	return desain.Id, nil
}

func (r *repository) GetDesain(ctx context.Context) ([]entity.Desain, error) {
	var desain []entity.Desain
	query := "SELECT id,user_id ,file_desain,catatan,status,tanggal_upload FROM desain"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return desain, err
	}
	defer rows.Close()

	for rows.Next() {
		var desn entity.Desain

		if err := rows.Scan(&desn.Id, &desn.UserId, &desn.FileDesain, &desn.Catatan, &desn.Status, &desn.TanggalUpload); err != nil {
			return desain, err
		}
		desain = append(desain, desn)

	}
	return desain, nil
}

func (r *repository) UpdateDesain(ctx context.Context, desain entity.Desain) (int64, error) {
	query := "UPDATE desain SET user_id=$2,file_desain=$3,catatan=$4,status=$5,tanggal_upload=$6 WHERE id =$1 RETURNING id"
	err := r.db.QueryRowContext(ctx, query, desain.Id, desain.UserId, desain.FileDesain, desain.Catatan, desain.Status, desain.TanggalUpload).Scan(&desain.Id)

	if err != nil {
		return 0, err
	}
	return desain.Id, nil
}
func (r *repository) UpdateStatusOnly(ctx context.Context, id int64, status string) error {
	query := "UPDATE desain SET status=$1 WHERE id=$2"
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *repository) DeleteDesain(ctx context.Context, id int64) error {
	query := "DELETE FROM desain WHERE id =$1"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	return nil
}
//Detail pesananan
func (r *repository) GetDesainDetail(ctx context.Context) ([]entity.DesainDetail, error) {
	query := `SELECT d.id,d.user_id,u.username,p.nama,d.file_desain,d.catatan,d.status,d.tanggal_upload FROM desain d LEFT JOIN users u ON d.user_id = u.id LEFT JOIN checkout c ON d.id = c.desain_id
		LEFT JOIN produk p ON c.produk_id = p.id`

	var result []entity.DesainDetail
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dd entity.DesainDetail
		err := rows.Scan(
			&dd.Id, &dd.UserId, &dd.Username, &dd.NamaProduk,
			&dd.FileDesain, &dd.Catatan, &dd.Status, &dd.TanggalUpload,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, dd)
	}

	return result, nil
}
