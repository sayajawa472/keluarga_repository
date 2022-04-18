package perlengkapan_sekolah_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	"testing"
)

func TestConnextionInsert(t *testing.T) {
	PerlengkapanSekolahRepository := NewPerlengkapanSekolah(golang_database.GetConnection())

	ctx := context.Background()
	PerlengkapanSekolah := entity.Perlengkapan_sekolah{
		Nama: "ransel",
		Bahan: "Kanvas",
	}

	resuly, err := PerlengkapanSekolahRepository.Insert(ctx, PerlengkapanSekolah)
	if err != nil {

		fmt.Println(result)
	}

	func comment_repository.TestFindById(t *testing.T)
	PerlengkapanSekolahRepository := NewPerlengkapanSekolah(golang_database.GetConnection())

	PerlengkapanSekolah, err := PerlengkapanSekolahRepository.FindById(context.Background(), 37)
	if err != nil {

		fmt.Println(PerlengkapanSekolah)
	}