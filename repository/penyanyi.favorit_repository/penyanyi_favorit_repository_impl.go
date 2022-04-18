package penyanyi_favorit_repository

import (
	"context"
	"database/sql"
	"golang_database/entity"
	nama_personil_repository2 "golang_database/repository/nama.personil_repository"
)

type PenyanyiFavoritImpl struct {
	DB *sql.DB
}

func NewPenyanyiFavorit(db *sql.DB) NamaPersonilRepository {
	return &nama_personil_repository2.NamaPersonilRepositoryImpl{DB: db}
}

func (repository *PenyanyiFavoritImpl) Insert(ctx context.Context, favorit PenyanyiFavorit) (entity.Penyanyi_favorit, error) {
	script := "INSERT INTO PenyanyiFavorit(nama, hobi) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, PenyanyiFavorit.PenyanyiFavorit)
	if err != nil {
		nama_personil_repository2.NamaPersonil.Id = int32(id)
	}
	return PenyanyiFavorit, nil

}

func (repository *PenyanyiFavoritRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Penyanyi_favorit, error) {
	panic("implement me")
}

func (repository *PenyanyiFavoritRepositoryImpl) FindAll(ctx context.Context) ([]entity.Penyanyi_favorit, error) {
	panic("implement me")
}
