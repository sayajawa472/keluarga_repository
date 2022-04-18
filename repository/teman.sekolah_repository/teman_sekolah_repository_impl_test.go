package teman_sekolah_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	"testing"
)

func TestConnectionInsert(t *testing.T) {
	TemanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

	ctx := context.Background()
	TemanSekolah := entity.Teman_sekolah{
		Nama: "Irmawati",
		Umur: "20 tahun",
	}

	result, err := TemanSekolahRepository.Insert(ctx, TemanSekolah)
	if err != nil {

		fmt.Println(result)
	}

	func comment_repository.TestFindById(t *testing.T)
	TemanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

	TemanSekolah, err := TemanSekolahRepository.FindById(context.Background(), 37)
	if err != nil {

		fmt.Println(TemanSekolah)
	}

	func TemanSekolahRepository.FindAll(t *testing.T) {
		KeluargaRepository := 	NewTemanSekolahRepository(golang_database.GetConnection())

		NamaPersonil, err := KeluargaRepository.FindById(context.Background(), 37)
		if err != nil {

			fmt.Println(NamaPersonil)
		}

		func TestUpdate(t *testing.T) {
			TemanSekolahRepository := NewTemanSekolahRepository(golang_database.GetConnection())

			ctx := context.Background()
			TemanSekolahRepository:= entity.Teman_sekolah
				JenisKelamin: "Wanita",
				Hobi: "Berenang",
			}

			func TestDelete(t*testing.T)
			{
				TemanSekolahRepository := TemanSekolahRepository(golang_database.GetConnection())

				ctx := context.Background()
				delete("Jenis Kelamin")
				entity.Teman_sekolah{
				}