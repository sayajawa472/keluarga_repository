package perlengkapan_sekolah_repository

import (
	"context"
	"errors"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	"strconv"
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


	func (repository *PerlengkapanSekolahImpl) FindById (ctx context.Context, id int32) (entity.Perlengkapan_sekolah, error) {
		script := "SELECT id, email, comment FROM WHERE id = ? LIMIT1"
		rows, err := PerlengkapanSekolahRepository.DB.QueryContext(ctx, script, Id)
		comment := entity.Perlengkapan_sekolah{}

		if err != nil : comment, err
		defer rows.Close()
		if rows.Next() {
			// ADA
			rows.Scan(&PerlengkapanSekolah.Id, &PerlengkapanSekolah.Nama, &PerlengkapanSekolah.Bahan)
			return PerlengkapanSekolah, nil
		} else {
			//tidak ada
			return PerlengkapanSekolah, errors.New("id " + strconv.Itoa(int(id)) + "not found")
		}
	}

	func (repository *PerlengkapanSekolah) FindAll(ctx context.Context) ([]entity.Perlengkapan_sekolah, error) {
		script := "SELECT id, Nama, Bahan FROM PerlengkapanSekolah"
		rows, err := repository.DB.QueryContext(ctx, script)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var PerlengkapanSekolah []entity.Perlengkapan_sekolah
		for rows.Next() {
			PerlengkapanSekolah := entity.Perlengkapan_sekolah{}
			rows.Scan(&PerlengkapanSekolah.Id, &PerlengkapanSekolah.Nama, &PerlengkapanSekolah.PerlengkapanSekolah)
			PerlengkapanSekolah = append(PerlengkapanSekolah.PerlengkapanSekolah)
		}
		return PerlengkapanSekolah, nil{
	}

		func TestUpdate(t *testing.T) {
			PerlengkapanSekolah := NewPerlengkapanSekolah(golang_database.GetConnection())

			ctx := context.Background()
			PerlengkapanSekolah := entity.Perlengkapan_sekolah{
				Nama : "ransel",
				Bahan: "kanvas",
			}

			func TestDelete(t*testing.T) {
				PerlengkapanSekolahRepository := NewPerlengkapanSekolah(golang_database.GetConnection())

				ctx := context.Background()
				delete("fungsi") entity.Perlengkapan_sekolah{
				}
