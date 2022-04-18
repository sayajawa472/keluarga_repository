package perlengkapan_sekolah_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type PerlengkapanSekolahImpl struct {
	DB *sql.DB
}

func NewPerlengkapanSekolah(db *sql.DB) PerlengkapanSekolahRepository {
	return &PerlengkapanSekolahImpl{DB: db}
}

func (repository *PerlengkapanSekolahImpl) insert(ctx context.Context, sekolah PerlengkapanSekolah) (entity.Perlengkapan_sekolah, error) {
	script := "INSERT INTO Perlengkapan sekolah(nama, bahan) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, PerlengkapanSekolah.PerlengkpanSekolah)
	if err != nil {
		id, err := result.LastInsertId()
		if err != nil {
			PerlengkapanSekolah.Id = int32(id)
		}
		return PerlengkapanSekolah, nil

	}
}

func (repository *PerlengkapanSekolah) FindAll(ctx context.Context) ([]PerlengkapanSekolah, error) {
	script := "SELECT id, nama, bahan FROM PerlengkapanSekolah"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var PerlengkapanSekolah []entity.Perlengkapan_sekolah
	for rows.Next() {
		PerlengkapanSekolah := entity.Perlengkapan_sekolah{}
		rows.Scan(&PerlengkapanSekolah.Id, &PerlengkapanSekolah.Nama, &PerlengkapanSekolah.Bahan)
		PerlengkapanSekolah = append(PerlengkapanSekolah, PerlengkapanSekolah)
	}
	return PerlengkapanSekolah, nil
}

func (repository *PerlengkapanSekolahImpl) FindById(ctx context.Context, id int32) (entity.Perlengkapan_sekolah, error) {
	script := "SELECT id, nama, bahan, fungsi, harga, FROM PerlengkpanSekolah WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	PerlengkapanSekolah := entity.Perlengkapan_sekolah{}

	if err != nil {
		return PerlengkapanSekolah, err
	}
	defer rows.Close()
	if rows.Next() {
		// ADA
		rows.Scan(&PerlengkapanSekolah.Id, &PerlengkapanSekolah.Nama, &PerlengkapanSekolah.Bahan, &PerlengkapanSekolah.Fungsi, &PerlengkapanSekolah.Harga)
		return PerlengkapanSekolah, nil
	} else {
		//tidak ada
		return PerlengkapanSekolah, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *PerlengkapanSekolahImpl) Update(ctx context.Context, id int32, PerlengkapanSekolah entity.Perlengkapan_sekolah) (entity.Perlengkapan_sekolah, error) {
	script := "SELECT id, nama, bahan, fungsi, harga, FROM PerlengkapanSekolah WHERE id = ? LIMIT1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return PerlengkapanSekolah, err
	}

	if rows.Next() {
		script := "UPDATE PerlengkapanSekolah SET nama = ?, bahan = ? fungsi + ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, PerlengkapanSekolah.Nama, PerlengkapanSekolah.Bahan, PerlengkapanSekolah.Fungsi, id)
		if err != nil {
			return PerlengkapanSekolah, err
		}
		PerlengkapanSekolah.Id = id
		return PerlengkapanSekolah, nil
	} else {
		return PerlengkapanSekolah, errors.New("Id" + strconv.Itoa(int(id)) + "Update Failed")
	}
}

func (repository *PerlengkapanSekolahImpl) Delete(ctx context.Context, PerlengkapanSekolah entity.Perlengkapan_sekolah) (entity.Perlengkapan_sekolah, error) {
	script := "DELETE FROM list_PerlengkapanSekolah WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, PerlengkapanSekolah.Id)
	if err != nil {
		return PerlengkapanSekolah, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return PerlengkapanSekolah, err
	}
	if rows == 0 {
		return PerlengkapanSekolah, err
	}
	return PerlengkapanSekolah, nil
}
