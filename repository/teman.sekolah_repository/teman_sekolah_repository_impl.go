package teman_sekolah_repository

import (
	"context"
	"database/sql"
	"golang_database/entity"
	"golang_database/repository/keluarga_repository"
)

type TemanSekolahRepositoryImpl struct {
	DB *sql.DB
}

func NewTemanSekolahRepository(db *sql.DB) keluarga_repository.KeluargaRepository {
	return &TemanSekolahRepositoryImpl{DB: db}
}

func (repository *NamaTemanRepositoryImpl) Insert(ctx context.Context, TemanSekolahz) (entity.Teman_sekolah, error) {
	script := "INSERT INTO teman sekolah(nama, umur) VALUE(?, ?)"
	result, err := repository.DB.ExecContex(ctx, script, TemanSekolah.TemanSekolah)
	if err != nil {
		id, err := result.LastInsertId()
		if err != nil {
			TemanSekolah.Id = int32(id)
		}
		return keluarga, nil

		}

		func (repository *TemanSekolahRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Teman_sekolah, error()
	{
		panic("implement me")
	}

	func (repository *TemanSekolahRepositoryImpl) FindAll(ctx context.Context) ([]entity.Teman_sekolah, error)
		panic("implement me2")
}
