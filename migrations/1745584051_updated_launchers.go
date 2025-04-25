package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2089973043")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"deleteRule": "@request.auth.id = owner.id && current_user.id = \"\"",
			"updateRule": "((@request.auth.id = owner.id || @request.auth.id ?= allowed_users.id) && (current_user.id = @request.auth.id || (current_user.id = \"\" && @request.body.current_user = @request.auth.id))) || @request.auth.id = id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2089973043")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"deleteRule": "@request.auth.id = owner.id",
			"updateRule": "@request.auth.id = owner.id || @request.auth.id ?= allowed_users.id || @request.auth.id = id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
