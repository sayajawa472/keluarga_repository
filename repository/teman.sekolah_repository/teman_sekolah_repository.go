package teman_sekolah_repository

import (
	"context"
	"golang_database/entity"
)

type TemanSekolah interface {
	Insert(ctx context.Context, TemanSekolah entity.Teman_sekolah) (entity.Perlengkapan_sekolah, error)
	FindById(ctx context.Context, Id int32) (entity.Teman_sekolah, error)
	FindByAll(ctx context.Context) ([]entity.Teman_sekolah, error)
	Update(ctx context.Context, Nama, Umur entity.Teman_sekolah) (entity.Teman_sekolah, error)
	Delete(ctx context.Context, JenisKelamin, Alamat entity.Teman_sekolah) (entity.Perlengkapan_sekolah, error)
}
