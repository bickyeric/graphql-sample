package main

import (
	"math/rand"
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func outletSeed() error {
	stores, err := model.Store{}.All()
	if err != nil {
		return err
	}
	for _, s := range stores {
		for i := 1; i <= 3; i++ {
			_, err := model.Outlet{
				Name:        "Outlet " + strconv.Itoa(i) + " Toko " + strconv.Itoa(s.ID),
				Address:     "Jln.outlet No : " + strconv.Itoa(i),
				PhoneNumber: strconv.Itoa(rand.Int()),
				City:        "Kota outlet " + strconv.Itoa(i),
				State:       "Provinsi outlet " + strconv.Itoa(i),
				Status:      true,
				Zip:         "Kode Pos" + strconv.Itoa(i),
				StoreID:     s.ID,
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
