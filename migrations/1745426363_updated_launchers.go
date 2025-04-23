package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2089973043")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(14, []byte(`{
			"hidden": false,
			"id": "date4109299031",
			"max": "",
			"min": "",
			"name": "last_ping_at",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(15, []byte(`{
			"hidden": false,
			"id": "date3281630529",
			"max": "",
			"min": "",
			"name": "last_user_ping_at",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2089973043")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("date4109299031")

		// remove field
		collection.Fields.RemoveById("date3281630529")

		return app.Save(collection)
	})
}
