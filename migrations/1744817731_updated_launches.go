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
			"createRule": "@request.auth.id = id_launcher.id_owner.id || @request.auth.id ?= id_launcher.allowed_users.id",
			"deleteRule": "@request.auth.id = id_launcher.id_owner || @request.auth.id ?= id_launcher.allowed_users.id",
			"listRule": "@request.auth.id != \"\"",
			"updateRule": "@request.auth.id = id_launcher.id_owner.id || @request.auth.id ?= id_launcher.allowed_users.id || @request.auth.id = id_launcher.id",
			"viewRule": "@request.auth.id != \"\""
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
			"createRule": null,
			"deleteRule": null,
			"listRule": null,
			"updateRule": null,
			"viewRule": null
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
