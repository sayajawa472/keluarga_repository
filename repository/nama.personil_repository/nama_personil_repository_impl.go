package nama_personil_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	"strconv"
)

type NamaPersonilRepositoryImpl struct {
	DB *sql.DB
}

func NewNamaPersonilRepository(db *sql.DB) NamaPersonilRepository {
	return &NamaPersonilRepositoryImpl{DB: db}
}

func (repository *NamaPersonilRepositoryImpl) Insert(ctx context.Context, NamaPersonil entity.Nama_personil) (entity.Nama_personil) {
	script := "INSERT INTO NamaPersonil(nama, NamaPersonil) VALUE (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, NamaPersonil.NamaPersonil)
	if err != nil {
		id, err := result.LastInsertId()
		if err != nil {
			NamaPersonil.Id = int32(id)
			return NamaPersonil, nil

		}

func (repository *NamaPersonilRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Perlengkapan_sekolah, error) {
			script := "SELECT id, nama, NamaPersonil FROM WHERE id = ? LIMIT1"
			rows, err := repository.DB.QueryContext(ctx, script, id)
			NamaPersonil := entity.Nama_personil{}

			if err != nil {
				return NamaPersonil, err
			}
			defer rows.Close()
			if rows.Next() {
				rows.Scan(&NamaPersonil.Id, &NamaPersonil.Nama, &NamaPersonil.jeniskelamin)
				return NamaPersonil, nil
			} else {
				return NamaPersonil, errors.New("id " + strconv.Itoa(int(id)) + "not found")
			}
		}

		func (repository *NamaPersonilRepositoryImpl) FindAll(ctx context.Context) ([]entity.Nama_personil, error) {
			script := "SELECT id, nama, JenisKelamin FROM NamaPersonil"
			rows, err := repository.DB.QueryContext(ctx, script)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var NamaPersonil []entity.Nama_personil
			for rows.Next() {
				NamaPersonil := entity.Nama_personil{}
				rows.Scan(&NamaPersonil.Id, &NamaPersonil.Nama, &NamaPersonil.JenisKelamin id)
				NamaPersonil = append(NamaPersonil, NamaPersonil)
			}
			return NamaPersonil, nil {
		}

		func (repository *NamaPersonilRepositoryImpl) Update(ctx context.Context, id int32, entity.Nama_personil) (entity.Nama_personil, error) {
			script := "SELECT id, nama, jenis kelamin, FROM NamaPersonil WHERE id = ? LIMIT1"
			rows, err := repository.DB.QueryContext(ctx, script, id)
			defer rows.Close()
			if err != nil {
				return NamaPersonil err
			}

			if rows.Next() {
				script := "UPDATE NamaPersonil SET id = ?, nama = ? JenisKelamin + ? WHERE id = ?"
				_, err := repository.DB.ExecContext(ctx, script, NamaPersonil.Nama, NamaPersonil.JenisKelamin, id)
				if err != nil {
					return NamaPersonil, err
				}
				NamaPersonil.Id = id
				return NamaPersonil, nil
			} else {
				return NamaPersonil, errors.New("Id" + strconv.Itoa(int(id)) + "Update Failed")
			}
		}

		func (repository *NamaPersonilRepositoryImpl) Delete(ctx context.Context, NamaPersonil entity.Nama_personil) (entity.Nama_personil, error) {
			script := "DELETE FROM list_NamaPersonil WHERE id = ?"
			result, err := repository.DB.ExecContext(ctx, script, NamaPersonil.Id)
			if err != nil {
				return NamaPersonil, err
			}
			rows, err := result.RowsAffected()
			if err != nil {
				return NamaPersonil, err
			}
			if rows == 0 {
				return NamaPersonil, err
			}
			return NamaPersonil, nil
		}
