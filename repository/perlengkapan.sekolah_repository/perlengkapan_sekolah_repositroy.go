package perlengkapan_sekolah_repository

import (
	"context"
	"golang_database/entity"
)

type PerlengkapanSekolah interface {
	Insert(ctx context.Context, PerlengkapanSekolah entity.Perlengkapan_sekolah) (entity.Perlengkapan_sekolah, error)
	FindById(ctx context.Context, Id int32) (entity.Perlengkapan_sekolah, error)
	FindAll(ctx context.Context) ([]entity.Perlengkapan_sekolah, error)
	Update(ctx context.Context, sekolah entity.Perlengkapan_sekolah) (entity.Perlengkapan_sekolah, error)
	Delete(ctx context.Context, id int32) (entity.Perlengkapan_sekolah, error)
}
