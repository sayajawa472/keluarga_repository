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

	result, err := nama_personil_repository2.NewNamaPersonilRepository.Insert(ctx, PenyanyiFavorit)
	if err != nil {

		fmt.Println(result)
	}

	func comment_repository.TestFindById(t *testing.T) {
	PenyanyiFavoritRepository := NewPenyanyiFavoritRepository(golang_database.GetConnection())

	PenyanyiFavorit, err := PenyanyiFavoritRepository.FindById(context.Background(), 37)
		if err != nil {

			fmt.Println(nama_personil_repository2.NamaPersonil)
		}