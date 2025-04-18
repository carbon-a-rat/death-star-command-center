package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
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
					"id": "number1653554401",
					"max": null,
					"min": null,
					"name": "acc_x",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number361237623",
					"max": null,
					"min": null,
					"name": "acc_y",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number2357288397",
					"max": null,
					"min": null,
					"name": "acc_z",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number2981183584",
					"max": null,
					"min": null,
					"name": "gyro_x",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number3333845238",
					"max": null,
					"min": null,
					"name": "gyro_y",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number1606361420",
					"max": null,
					"min": null,
					"name": "gyro_z",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number3150619954",
					"max": null,
					"min": null,
					"name": "altitude",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number1605345383",
					"max": null,
					"min": null,
					"name": "pressure",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "number3192793708",
					"max": null,
					"min": null,
					"name": "temperature",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				}
			],
			"id": "pbc_3694955888",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_rG9OPplpdi` + "`" + ` ON ` + "`" + `flight_data` + "`" + ` (\n  ` + "`" + `launch_id` + "`" + `,\n  ` + "`" + `time_rel` + "`" + `\n)"
			],
			"listRule": null,
			"name": "flight_data",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3694955888")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
