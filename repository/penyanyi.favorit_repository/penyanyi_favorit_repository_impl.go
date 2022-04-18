package penyanyi_favorit_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang_database/entity"
	nama_personil_repository2 "golang_database/repository/nama.personil_repository"
	"strconv"
)

type PenyanyiFavoritImpl struct {
	DB *sql.DB
}

func NewPenyanyiFavorit(db *sql.DB) PenyanyiFavorit {
	return &PenyanyiFavorit

func (repository *PenyanyiFavoritImpl) Insert(ctx context.Context, favorit PenyanyiFavorit) (entity.Penyanyi_favorit, error) {
	script := "INSERT INTO PenyanyiFavorit(nama, hobi) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, PenyanyiFavorit.PenyanyiFavorit)
	if err != nil {
		PenyanyiFavorit.
	}
	return PenyanyiFavorit, nil

}

	func (repository *PenyanyiFavoritImpl) FindById(ctx context.Context, id int32) (entity.Penyanyi_favorit, error) {
		script := "SELECT id, nama, jeniskelamin, hobi FROM penyanyifavorit WHERE id = ? LIMIT1"
		rows, err := repository.DB.QueryContext(ctx, script, id)
		keluarga := entity.Penyanyi_favorit{}

		if err != nil {
			return PenyanyiFavorit err
		}
		}
		defer rows.Close()
		if rows.Next() {
			// ADA
			rows.Scan(&PenyanyiFavorit().Id, &PenyanyiFavorit.nama,  &PenyanyiFavorit.jeniskelamin, &keluarga.Hobi)
			return jeniskelamin, nil
		} else {
			//tidak ada
			return jeniskelamin, errors.New("id " + strconv.Itoa(int(id)) + "not found")
		}
	}

	func (repository *PenyanyiFavoritImpl) Update(ctx context.Context, id int32, penyanyifavorit entity.Penyanyi_favorit) (entity.Penyanyi_favorit, error) {
		script := "SELECT id, nama, jenisKelamin, hobi FROM PenyanyiFavorit WHERE id = ? LIMIT1"
		rows, err := repository.DB.QueryContext(ctx, script, id)
		defer rows.Close()
		if err != nil {
			return penyanyifavorit, err
		}

		if rows.Next() {
			script := "UPDATE PenyanyiFavorit SET nama = ?, jeniskelamin = ? hobi + ? WHERE id = ?"
			_, err := repository.DB.ExecContext(ctx, script, PenyanyiFavorit.Nama, PenyanyiFavorit.JenisKelamin, PenyanyiFavorit.Hobi, id)
			if err != nil {
				return penyanyifavorit, err
			}
			PenyanyiFavorit.id
			return penyanyifavorit, nil
		} else {
			return penyanyifavorit, errors.New("Id" + strconv.Itoa(int(id)) + "Update Failed")
		}
	}

	func (repository *PenyanyiFavoritImpl) Delete(ctx context.Context, PenyanyiFavorit entity.Penyanyi_favorit) (entity.Penyanyi_favorit, error) {
		script := "DELETE FROM list_keluarga WHERE id = ?"
		result, err := repository.DB.ExecContext(ctx, script, PenyanyiFavorit.Id)
		if err != nil {
			return PenyanyiFavorit, err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return PenyanyiFavorit, err
		}
		if rows == 0 {
			return PenyanyiFavorit, err
		}
		return PenyanyiFavorit := nil
	}
