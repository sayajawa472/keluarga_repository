package nama_personil_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	"testing"
)

func TestConnectInsert(t *testing.T)  {
	NamaPersonilRepositoryImpl := NewNamaPersonilRepository(golang_database.GetConnection())

	ctx := context.Background()
	NamaPersonil  := entity.Nama_personil{
		Nama: "Harry Styles",
		Jenis_Kelamin: "Pria",
	}

	result, err := NewNamaPersonilRepository.Insert(ctx, NamaPersonil)
	if err != nil {

		fmt.Println(result)

	}

	func NamaPersonilRepositoryImpl.TestFindById(t *testing.T) {
		NewNamaPersonilRepository(golang_database.GetConnection())

		NamaPersonil, err := NewNamaPersonilRepository.FindById(context.Background(), 37)
		if err != nil {

			fmt.Println(NamaPersonil)
		}

		func NamaPersonilRepository.TestFindAll(t *testing.T) {
			NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())

			NamaPersonil, err := NamaPersonilRepository.FindById(context.Background(), 37)
			if err != nil {

				fmt.Println(NamaPersonil)
			}

			func TestUpdate(t *testing.T) {
				NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())

				ctx := context.Background()
				NamaPersonil := entity.Comment{
					Nama : "Harry Styles",
					JenisKelamin: "Pria",
				}

				func TestDelete(t*testing.T) {
					NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())

					ctx := context.Background()
					delete("Jenis Kelamin") entity.Nama_personil{
					}
