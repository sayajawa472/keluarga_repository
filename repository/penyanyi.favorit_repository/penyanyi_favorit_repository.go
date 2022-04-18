package penyanyi_favorit_repository

import (
	"context"
	"golang_database/entity"
)

type PenyanyiFavorit interface {
	Insert(ctx context.Context, NamaPersonil entity.Nama_personil) (entity.Nama_personil, error)
	FindById(ctx context.Context, Id int32) (entity.Penyanyi_favorit, error)
	FindAll(ctx context.Context) ([]entity.Penyanyi_favorit, error)
	Update(ctx context.Context, Nama, Hobi entity.Penyanyi_favorit) (entity.Penyanyi_favorit, error)
	Delete(ctx context.Context, Umur, JenisKelamin entity.Penyanyi_favorit) (entity.Penyanyi_favorit, error)
}
