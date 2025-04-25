package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1289051546")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id = launcher.current_user.id",
			"deleteRule": "@request.auth.id = launcher.current_user.id",
			"updateRule": "@request.auth.id = launcher.current_user.id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1289051546")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id = launcher.owner.id || @request.auth.id ?= launcher.allowed_users.id",
			"deleteRule": "@request.auth.id = launcher.owner || @request.auth.id ?= launcher.allowed_users.id",
			"updateRule": "@request.auth.id = launcher.owner.id || @request.auth.id ?= launcher.allowed_users.id || @request.auth.id = launcher.id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
