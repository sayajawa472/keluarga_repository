package keluarga_repository

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"golang_database/entity"
	"log"
)

type KeluargaRepositoryImpl struct {
	DB *sql.DB
}

func NewKeluargaRepository(db *sql.DB) KeluargaRepository {
	return &KeluargaRepositoryImpl{DB: db}
}

func (repository *KeluargaRepositoryImpl) Insert(ctx context.Context, keluarga entity.Keluarga) (entity.Keluarga, error) {
	script := "INSERT INTO keluarga(nama, umur) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, keluarga.Umur, keluarga.Nama)
	if err != nil {
		return keluarga, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return keluarga, err
	}
	keluarga.Id = int32(id)
	return keluarga, nil
}

func (repository *KeluargaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Keluarga, error) {
	panic("implement me")
}

func (repository *KeluargaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	panic("implement me2")
}

func (repository KeluargaRepositoryImpl) update(ctx context.Context, keluarga entity.Keluarga) {
	script := "SELECT Keluarga strock =?, WHERE Id = ?"
	rows, err := repository.DB.PrepareContext(ctx, script)
	if err != nil {
		return nil, err
	}
	_, err = rows.ExecContext(
		ctx,
		keluarga.Nama,
		keluarga.Umur,
		keluarga.Hobi,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return keluarga, nil
}

func (repository *KeluargaRepositoryImpl) delete(ctx context.Context, id int32)
	db, err := mysql.Config()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return keluarga, nil
}