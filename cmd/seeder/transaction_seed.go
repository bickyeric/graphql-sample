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
		_, err := o.Employee()
		if err != nil {
			return err
		}
	}
	return nil
}
