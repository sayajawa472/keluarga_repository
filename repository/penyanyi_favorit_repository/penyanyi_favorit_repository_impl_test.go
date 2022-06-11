package penyanyi_favorit_repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang_database"
	belajar_db "golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnectionInsert(t *testing.T)  {
	PenyanyiFavoritRepository := NewPenyanyiFavorit(golang_database.GetConnection())

	ctx := context.Background()
	PenyanyiFavorit:= entity.Penyanyi_favorit{
		Nama: "Taylor Swift",
		Hobi: "Menulis Lagu",
	}

	result, err := PenyanyiFavorit.Insert(ctx, PenyanyiFavorit)
	if err != nil {
		panic(err)
		fmt.Println(result)
	}

	func TestFindById(t *testing.T) {
	PenyanyiFavoritRepository := NewPenyanyiFavorit(golang_database.GetConnection())

	PenyanyiFavorit, err := PenyanyiFavoritRepository.FindById(context.Background(), 37)
		if err != nil {
			panic(err)
			fmt.Println(NewPenyanyiFavorit)
		}

		func TestFindAll(t *testing.T)
		{
			PenyanyiFavorit := NewPenyanyiFavorit(golang_database.GetConnection())

			PenyanyiFavorit, err := PenyanyiFavorit.FindById(context.Background(), 37)
			if err != nil {
				panic(err)
				fmt.Println(PenyanyiFavorit)
			}

			func TestDelete(t *testing.T) {
				PenyanyiFavorit := NewPenyanyiFavorit(golang_database.GetConnection())
				result, err := PenyanyiFavoritRepository.Delete(context.Background(), 1)
				if err != nil {
					panic(err)
				}
				fmt.Println(result)
		}

			}


				func TestUpdate(t*testing.T)
		{

			PenyanyiFavoritRepository := NewPenyanyiFavorit(golang_database.GetConnection())

			ctx := context.Background()
			PenyanyiFavorit := entity.PenyanyiFavorit{
				Id:           1,
				Nama:         "Taylor swift",
				Jenis_Kelamin: "Wanita",
				Hobi:         "Menulis Lagu",
			}

			result, err := PenyanyiFavoritRepository.Update(ctx, PenyanyiFavorit)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
		}