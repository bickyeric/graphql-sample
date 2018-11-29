package main

import (
	"strconv"

	"github.com/bickyeric/graphql-sample/pkg/model"
)

func customerSeed() error {
	stores, err := model.Store{}.All()
	if err != nil {
		return err
	}
	for _, s := range stores {
		for i := 1; i <= 10; i++ {
			_, err := model.Customer{
				StoreID: s.ID,
				Name:    "Pelanggan ke " + strconv.Itoa(i),
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
