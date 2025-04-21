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
			"createRule": "@request.auth.id = launcher.owner.id || @request.auth.id ?= launcher.allowed_users.id",
			"deleteRule": "@request.auth.id = launcher.owner || @request.auth.id ?= launcher.allowed_users.id",
			"updateRule": "@request.auth.id = launcher.owner.id || @request.auth.id ?= launcher.allowed_users.id || @request.auth.id = launcher.id"
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1889285177",
			"hidden": false,
			"id": "relation3901624051",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "rocket",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2089973043",
			"hidden": false,
			"id": "relation726129024",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "launcher",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
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
			"createRule": "@request.auth.id = id_launcher.id_owner.id || @request.auth.id ?= id_launcher.allowed_users.id",
			"deleteRule": "@request.auth.id = id_launcher.id_owner || @request.auth.id ?= id_launcher.allowed_users.id",
			"updateRule": "@request.auth.id = id_launcher.id_owner.id || @request.auth.id ?= id_launcher.allowed_users.id || @request.auth.id = id_launcher.id"
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1889285177",
			"hidden": false,
			"id": "relation3901624051",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "id_rocket",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2089973043",
			"hidden": false,
			"id": "relation726129024",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "id_launcher",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
