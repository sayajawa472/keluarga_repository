package perlengkapan_sekolah_repository

import (
	"context"
	"database/sql"
	"golang_database/entity"
	"golang_database/repository/nama.personil_repository"
)

type PerlengkapanSekolahImpl struct {
	DB *sql.DB
}

func NewPerlengkapanSekolah(db *sql.DB) PerlengkapanSekolahRepository {
	return &PerlengkapanSekolahImpl{DB: db,}
	}

func (repository *PerlengkapanSekolahImpl) insert(ctx context.Context, sekolah PerlengkapanSekolah) (entity.Perlengkapan_sekolah, error)  {
	script := "INSERT INTO Perlengkapan sekolah(nama, bahan) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, PerlengkapanSekolah.PerlengkpanSekolah)
		if err != nil {
			id , err := result.LastInsertId()
			if err != nil {
				PerlengkapanSekolah.Id = int32(id)
			}
			return PerlengkapanSekolah, nil

			}

			func(repository *PerlengkapanSekolahImpl) FindById(ctx context.Context, id int32) (entity.Perlengkapan_sekolah, err) {
		panic("implement me")
	}
	func (repository *nama_personil_repository.NamaPersonilRepositoryImpl) FindAll (ctx context.Context) ([]entity.Perlengkapan_sekolah, error) {
		panic("implement me")
	}