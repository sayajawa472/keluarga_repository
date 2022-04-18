package teman_sekolah_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"golang_database/repository/keluarga_repository"
	"strconv"
)

type TemanSekolahRepositoryImpl struct {
	DB *sql.DB
}

func NewTemanSekolahRepository(db *sql.DB) keluarga_repository.KeluargaRepository {
	return &TemanSekolahRepositoryImpl{DB: db}
}

func (repository *NamaTemanRepositoryImpl) Insert(ctx context.Context, TemanSekolah) (entity.Teman_sekolah, error) {
	script := "INSERT INTO teman sekolah(nama, umur) VALUE(?, ?)"
	result, err := repository.DB.ExecContex(ctx, script, TemanSekolah())
	if err != nil {
		id, err := result.LastInsertId()
		if err != nil {
			TemanSekolah.Id = int32(id)
		}
		return keluarga, nil

		}

	func (repository *TemanSekolahRepositoryImpl) FindAll(ctx context.Context) ([]entity.Perlengkapan_sekolah, error) {
		script := "SELECT id, nama, alamat, hobi FROM keluarga"
		rows, err := repository.DB.QueryContext(ctx, script)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var TemanSekolah []entity.Teman_sekolah
		for rows.Next() {
			TemanSekolah := entity.Teman_sekolah{}
			rows.Scan(&TemanSekolah.Id, &TemanSekolah.Nama, &TemanSekolah.Alamat, &TemanSekolah.Hobi)
			TemanSekolah = append(TemanSekolah, TemanSekolah)
		}
		return TemanSekolah, nil
	}

	func (repository *TemanSekolahRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Teman_sekolah, error) {
		script := "SELECT id, nama, alamat, hobi FROM TemanSekolah WHERE id = ? LIMIT1"
		rows, err := repository.DB.QueryContext(ctx, script, id)
		TemanSekolah := entity.Teman_sekolah

		if err != nil {
			return TemanSekolah, err
		}
		defer rows.Close()
		if rows.Next() {
			// ADA
			rows.Scan(&TemanSekolah.Id, &TemanSekolah.Nama, &TemanSekolah.Alamat, &TemanSekolah.Hobi)
			return TemanSekolah, nil
		} else {
			//tidak ada
			return TemanSekolah, errors.New("id " + strconv.Itoa(int(id)) + "not found")
		}
	}

	func (repository *TemanSekolahRepositoryImpl) Update(ctx context.Context, id int32, TemanSekolah entity.Teman_sekolah) (entity.Teman_sekolah, error) {
		script := "SELECT id, nama, alamat, hobi FROM TemanSekolah WHERE id = ? LIMIT1"
		rows, err := repository.DB.QueryContext(ctx, script, id)
		defer rows.Close()
		if err != nil {
			return TemanSekolah, err
		}

		if rows.Next() {
			script := "UPDATE Keluarga SET nama = ?, alamat = ? hobi + ? WHERE id = ?"
			_, err := repository.DB.ExecContext(ctx, script, TemanSekolah.Nama, TemanSekolah.Alamat, TemanSekolah.Hobi, id)
			if err != nil {
				return TemanSekolah, err
			}
			TemanSekolah.Id = id
			return TemanSekolah, nil
		} else {
			return TemanSekolah, errors.New("Id" + strconv.Itoa(int(id)) + "Update Failed")
		}
	}

	func (repository *TemanSekolah Delete(ctx context.Context, TemanSekolah entity.Teman_sekolah) (entity.Teman_sekolah, error) {
		script := "DELETE FROM list_TemanSekolah WHERE id = ?"
		result, err := repository.DB.ExecContext(ctx, script,TemanSekolah.Id)
		if err != nil {
			return TemanSekolah, err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return TemanSekolah, err
		}
		if rows == 0 {
			return TemanSekolah, err
		}
		return TemanSekolah, nil
	}
