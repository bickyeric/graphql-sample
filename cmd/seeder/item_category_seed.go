package main

import (
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func itemCategorySeed() error {
	outlets, err := model.Outlet{}.All()
	if err != nil {
		return err
	}
	for _, o := range outlets {
		for i := 1; i <= 3; i++ {
			_, err := model.ItemCategory{
				Name:     "Category no " + strconv.Itoa(i),
				OutletID: o.ID,
				StoreID:  o.StoreID,
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
