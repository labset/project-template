package main

import (
	"platform/backend/config"
	gendbtodo "platform/backend/internal/gendb/todo"
)

type dependencies struct {
	todoStore gendbtodo.Querier
}

func setupDependencies(_ config.Config, conn *connections) (*dependencies, error) {
	todoStore := gendbtodo.New(conn.db)

	return &dependencies{
		todoStore: todoStore,
	}, nil
}
