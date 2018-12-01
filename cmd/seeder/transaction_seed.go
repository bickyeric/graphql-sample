package main

import (
	"github.com/bickyeric/graphql-sample/pkg/model"
)

func transactionSeed() error {
	outlet, err := model.Outlet{}.All()
	if err != nil {
		return err
	}
	for _, o := range outlet {
		employees, err := o.Employee()
		if err != nil {
			return err
		}
		for _, e := range employees {
			_, err := model.Transaction{
				EmployeeID: e.ID,
				Status:     "PAID",
			}.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
