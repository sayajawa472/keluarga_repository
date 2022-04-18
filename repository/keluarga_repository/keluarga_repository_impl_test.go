package keluarga_repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"golang_database/repository/comment_repository"
	"testing"
)

func TestConnectionInsert(t *testing.T)  {
	KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())

	ctx := context.Background()
	keluarga := entity.Keluarga{
		Nama: "Khansa Zahra",
		Umur: "16 tahun",
	}

	result, err := KeluargaRepository.Insert(ctx, keluarga)
	if err != nil {

		fmt.Println(result)
	}

	func KeluargaRepository.FindById(t *testing.T) {
		KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())

		keluarga, err := comment_repository.CommentRepository.FindById(context.Background(), 37)
		if err != nil {

			fmt.Println(keluarga)
		}

		func KeluargaRepository.FindAll(t *testing.T) {
			KeluargaRepository := 	NewKeluargaRepository(golang_database.GetConnection())

			Keluarga, err := KeluargaRepository.FindById(context.Background(), 37)
			if err != nil {

				fmt.Println(keluarga)
			}

			func  TestUpdate(t *testing.T) {
			KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())
			ctx := context.Background()
			keluarga := entity.Keluarga{
				Nama: "Khansa Zahra",
				Hobi: "bermain futsal",
			}
			result, err := KeluargaRepository.Update(ctx, keluarga)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			}

			func TestDelete(t *testing.T)
				KeluargaRepository := NewKeluargaRepository(golang_database.GetConnection())
				ctx := context.Background()
				result, err := KeluargaRepository.Delete(ctx, 1)
				if err != nil {
					panic(err)
				}
				fmt.Println(result)
			}