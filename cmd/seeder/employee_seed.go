package main

import (
	"math/rand"
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func employeeSeed() error {
	outlets, err := model.Outlet{}.All()
	if err != nil {
		return err
	}
	for _, o := range outlets {
		for i := 1; i <= 5; i++ {
			_, err := model.Employee{
				OutletID:    o.ID,
				StoreID:     o.StoreID,
				FirstName:   "depan" + strconv.Itoa(i),
				LastName:    "belakang" + strconv.Itoa(i),
				PhoneNumber: strconv.Itoa(rand.Int())[:8],
				Email:       "karyawan" + strconv.Itoa(i) + "@example.co.id",
				Password:    "JanganBilangSiapaSiapa",
				Confirmed:   false,
				Active:      false,
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
