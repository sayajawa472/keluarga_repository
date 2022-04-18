package penyanyi_favorit_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	nama_personil_repository2 "golang_database/repository/nama.personil_repository"
	"testing"
)

func TestConnectionInsert(t *testing.T)  {
	PenyanyiFavoritRepository := NewPenyanyiFavorit(golang_database.GetConnection())

	ctx := context.Background()
	PenyanyiFavorit:= entity.Penyanyi_favorit{
		Nama: "Taylor Swift",
		Hobi: "Menulis Lagu",
	}

	result, err := nama_personil_repository2.NewPenyanyiFavoritRepository.Insert(ctx, PenyanyiFavorit)
	if err != nil {

		fmt.Println(result)
	}

	func PenyanyiFavoritRepository.FindById(t *testing.T) {
	PenyanyiFavoritRepository := NewPenyanyiFavorit(golang_database.GetConnection())

	PenyanyiFavorit, err := PenyanyiFavoritRepository.FindById(context.Background(), 37)
		if err != nil {

			fmt.Println(NewPenyanyiFavorit)
		}

		func comment_repository.FindAll(t *testing.T) {
			PenyanyiFavorit := NewPenyanyiFavorit(golang_database.GetConnection())

		PenyanyiFavorit, err := NewPenyanyiFavorit()(context.Background(), 37)
			if err != nil {

				fmt.Println(PenyanyiFavorit)
			}

			func TestUpdate(t *testing.T) {
				PenyanyiFavorit := NewPenyanyiFavorit(golang_database.GetConnection())

				ctx := context.Background()
				PenyanyiFavorit := entity.Penyanyi_favorit{
					Nama : "Taylor swift",
					JenisKelamin: "wanita",
				}

				func TestDelete(t*testing.T) {
					PenyanyiFavoritRepository := NewPenyanyiFavorit(golang_database.GetConnection())

					ctx := context.Background()
					delete("Jenis Kelamin") entity.Penyanyi_favorit{
					}
