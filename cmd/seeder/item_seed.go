package main

import (
	"math/rand"
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func itemSeed() error {
	outlets, err := model.Outlet{}.All()
	if err != nil {
		return err
	}
	for _, o := range outlets {
		itemCategory, err := o.ItemCategory()
		if err != nil {
			return err
		}
		for i := 1; i <= 15; i++ {
			ic := itemCategory[rand.Intn(len(itemCategory))]
			mustNull := rand.Intn(2)
			item := model.Item{
				OutletID:    o.ID,
				StoreID:     o.StoreID,
				Name:        "Barang jualan no " + strconv.Itoa(i),
				Description: "Deskripsi untuk barang no " + strconv.Itoa(i),
				CategoryID:  ic.ID,
			}
			if mustNull == 0 {
				item.CategoryID = 0
			}
			_, err := item.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
