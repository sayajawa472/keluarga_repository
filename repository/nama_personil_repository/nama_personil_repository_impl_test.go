package nama_personil_repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang_database"
	belajar_db "golang_database"
	"golang_database/entity"
	"testing"
)

func TestConnectInsert(t *testing.T) {
	NamaPersonil := NewNamaPersonilRepository(golang_database.GetConnection())

	ctx := context.Background()
	NamaPersonil := entity.NamaPersonil{
		Nama:         "Harry Styles",
		JenisKelamin: "Pria",
	}
}

	result, err := NewNamaPersonilRepository.Insert(ctx, NamaPersonil)
	if err != nil {

		fmt.Println(result)

	}
	func TestFindById(t *testing.T) {
		NewNamaPersonilRepository(golang_database.GetConnection())

		NamaPersonil, err := NewNamaPersonilRepository.FindById(context.Background(), 37)
		if err != nil {
			panic(err)
			fmt.Println(NamaPersonil)
		}

		func TestFindAll(t *testing.T)
			NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())

			NamaPersonil, err := NamaPersonilRepository.FindById(context.Background(), 37)
			if err != nil {
				panic(err)
				fmt.Println(NamaPersonil)
			}

			func TestDelete(t *testing.T) {

				NamaPersonilRepository := NewNamaPersonilRepository(golang_database.GetConnection())
				result, err := NamaPersonilRepository.Delete(context.Background(), 1)
				if err != nil {
					panic(err)
				}
				fmt.Println(result)

			}
	}

			func TestUpdate(t*testing.T) {
					NamaPersonilRepository := NewNamaPersonilRepository(belajar_db.GetConnection())

					ctx := context.Background()
					NamaPersonil := entity.Nama_personil{
						Id: 1,
						Nama: "Harry Styles",
						jenisKelamin: "Pria",
					}

					result, err := NamaPersonilRepository.Update(ctx, NamaPersonil)
					if err != nil {
						panic(err)
					}
					fmt.Println(result)
					}
