package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": "@request.auth.id = launch.launcher.id",
			"deleteRule": "@request.auth.id = launch.launcher.id",
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"cascadeDelete": true,
					"collectionId": "pbc_1289051546",
					"hidden": false,
					"id": "relation2042058741",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "launch",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"hidden": false,
					"id": "number775666903",
					"max": null,
					"min": null,
					"name": "time_rel",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number3113930206",
					"max": null,
					"min": null,
					"name": "volume",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				}
			],
			"id": "pbc_1758497698",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_XKWCbW1Aco` + "`" + ` ON ` + "`" + `water_loading_data` + "`" + ` (\n  ` + "`" + `launch` + "`" + `,\n  ` + "`" + `time_rel` + "`" + `\n)"
			],
			"listRule": "@request.auth.id != \"\"",
			"name": "water_loading_data",
			"system": false,
			"type": "base",
			"updateRule": "@request.auth.id = launch.launcher.id",
			"viewRule": "@request.auth.id != \"\""
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1758497698")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
