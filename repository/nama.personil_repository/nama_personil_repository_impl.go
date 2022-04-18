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

		func (repository *NamaPersonilRepositoryImpl) FindById(ctx context.Context) ([]entity.Nama_personil, error) {
			panic("implement me")
		}

		func (repository *NamaPersonilRepositoryImpl) FindAll(ctx context.Context) ([]entity.Nama_personil, error) {
			panic("implement me")
		}