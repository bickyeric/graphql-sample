package main

import (
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func variantSeed() error {
	items, err := model.Item{}.All()
	if err != nil {
		return err
	}
	for _, item := range items {
		for i := 1; i <= 3; i++ {
			_, err := model.Variant{
				ItemID:     item.ID,
				Price:      (i + item.ID) * 100,
				SKU:        "S" + strconv.Itoa(item.StoreID) + "O" + strconv.Itoa(item.OutletID) + "I" + strconv.Itoa(item.ID),
				TrackStock: false,
				Alert:      false,
				TrackCOGS:  false,
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
