package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3694955888")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_rG9OPplpdi` + "`" + ` ON ` + "`" + `flight_data` + "`" + ` (\n  ` + "`" + `time_rel` + "`" + `,\n  ` + "`" + `id_launch` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation1974573518")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(11, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1289051546",
			"hidden": false,
			"id": "relation1602282008",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "id_launch",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3694955888")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_rG9OPplpdi` + "`" + ` ON ` + "`" + `flight_data` + "`" + ` (\n  ` + "`" + `launch_id` + "`" + `,\n  ` + "`" + `time_rel` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2089973043",
			"hidden": false,
			"id": "relation1974573518",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "launch_id",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation1602282008")

		return app.Save(collection)
	})
}
