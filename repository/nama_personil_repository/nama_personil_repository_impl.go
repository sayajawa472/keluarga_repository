package nama_personil_repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type NamaPersonilRepositoryImpl struct {
	DB *sql.DB
}

func NewNamaPersonilRepository(db *sql.DB) *NamaPersonilRepositoryImpl{
	return &NamaPersonilRepositoryImpl{DB: db}
}

func (repository *NamaPersonilRepositoryImpl) Insert(ctx context.Context, NamaPersonil entity.Nama_personil) (entity.Nama_personil) {
	script := "INSERT INTO NamaPersonil(nama, NamaPersonil) VALUE (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, NamaPersonil.Nama)
	if err != nil {
		result NamaPersonil, err
	}
	id, err := result.LastInsertId()
		if err != nil {
			return NamaPersonil, err

		}
		keluarga.id + int32(id)
		return NamaPersonil)
	}
}
		func(repository *NamaPersonilRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Nama_personil, error) {
			script := "SELECT id, nama, JenisKelamin FROM NamaPersonil WHERE id = ? LIMIT1"
			rows, err := repository.DB.QueryContext(ctx, script, id)
			NamaPersonil := entity.Nama_personil{}

			if err != nil {
				return NamaPersonil, err
			}
			defer rows.Close()
			if rows.Next() {
				rows.Scan(&NamaPersonil.id, &NamaPersonil.Nama)
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
				rows.Scan(&NamaPersonil.Id, &NamaPersonil.Nama, )
				NamaPersonil = append(NamaPersonil, NamaPersonil)
			}
			return NamaPersonil, nil
		}

		func (repository *NamaPersonilRepositoryImpl) Update(ctx context.Context, id int32, entity.Nama_personil) (entity.Nama_personil, error) {
			script := "UPDATE NamaPersonil SET id = ?, Nama = ?, JenisKelamin = ? WHERE id = ?"
			result, err := repository.DB.ExecContext(ctx, script, NamaPersonil.id, NamaPersonil.Nama, NamaPersonil)
			if err != nil {
				return NamaPersonil err
			}

			rowCnt, err := result.RowsAffected()
			if err != nil {
				return NamaPersonil, err

			}
			if rowCnt == 0 {
				return NamaPersonil, err

			}
				return NamaPersonil, err

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
