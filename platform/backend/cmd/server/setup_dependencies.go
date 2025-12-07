package main

import (
	"platform/backend/config"
	gendbtodo "platform/backend/internal/gendb/todo"
	"platform/backend/pkg/workos"
)

type dependencies struct {
	todoStore    gendbtodo.Querier
	workOsClient workos.Client
}

func setupDependencies(cfg config.Config, conn *connections) (*dependencies, error) {
	todoStore := gendbtodo.New(conn.db)

	workOsClient := workos.NewClient(workos.ClientConfig{
		APIKey:      cfg.WorkOs.APIKey,
		ClientID:    cfg.WorkOs.ClientID,
		RedirectURI: cfg.WorkOs.RedirectURI,
	})

	return &dependencies{
		todoStore:    todoStore,
		workOsClient: workOsClient,
	}, nil
}
