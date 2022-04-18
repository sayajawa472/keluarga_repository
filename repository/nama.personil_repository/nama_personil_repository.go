package nama_personil_repository

import (
	"context"
	"golang_database/entity"
)

type NamaPersonil interface {
	Insert(ctx context.Context, nama, JenisKelamin entity.Nama_personil) (entity.Nama_personil, error)
	FindById(ctx context.Context, id int32) (entity.Nama_personil, error)
	FindAll(ctx context.Context) ([]entity.Nama_personil, error)
	Update(ctx context.Context, Nama, JenisKelamin entity.Nama_personil) (entity.Nama_personil, error)
	Delete(ctx context.Context, JenisKelamin entity.Nama_personil) (entity.Nama_personil, error)
}
