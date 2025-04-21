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
			"createRule": "@request.auth.id = launch.launcher.id",
			"deleteRule": "@request.auth.id = launch.launcher.id",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_rG9OPplpdi` + "`" + ` ON ` + "`" + `flight_data` + "`" + ` (\n  ` + "`" + `time_rel` + "`" + `,\n  ` + "`" + `launch` + "`" + `\n)"
			],
			"updateRule": "@request.auth.id = launch.launcher.id"
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(11, []byte(`{
			"cascadeDelete": true,
			"collectionId": "pbc_1289051546",
			"hidden": false,
			"id": "relation1602282008",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "launch",
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
			"createRule": "@request.auth.id = id_launch.id_launcher.id",
			"deleteRule": "@request.auth.id = id_launch.id_launcher.id",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_rG9OPplpdi` + "`" + ` ON ` + "`" + `flight_data` + "`" + ` (\n  ` + "`" + `time_rel` + "`" + `,\n  ` + "`" + `id_launch` + "`" + `\n)"
			],
			"updateRule": "@request.auth.id = id_launch.id_launcher.id"
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(11, []byte(`{
			"cascadeDelete": true,
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
	})
}
