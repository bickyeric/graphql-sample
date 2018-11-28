package main

import (
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func storeSeed() error {
	storeCategories, err := model.StoreCategory{}.All()
	if err != nil {
		return err
	}
	i := 0
	for _, c := range storeCategories {
		for index := 0; index < 3; index++ {
			i++
			_, err := model.Store{
				CategoryID:  c.ID,
				TypeID:      c.TypeID,
				Name:        "Toko nomor " + strconv.Itoa(i),
				Address:     "Alamat toko " + strconv.Itoa(i),
				City:        "Kota toko " + strconv.Itoa(i),
				State:       "Provinsi toko " + strconv.Itoa(i),
				Zip:         "Kode Pos toko " + strconv.Itoa(i),
				Email:       "toko" + strconv.Itoa(i) + "@example.com",
				Description: "Deskripsi toko " + strconv.Itoa(i),
				Website:     "https://www.toko" + strconv.Itoa(i) + ".id",
				Twitter:     "@toko" + strconv.Itoa(i),
				Facebook:    "toko" + strconv.Itoa(i),
				Instagram:   "@toko" + strconv.Itoa(i),
				Image:       "/toko/toko" + strconv.Itoa(i) + ".jpg",
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
