package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": "@request.auth.id = launcher.id",
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
					"hidden": false,
					"id": "date2782324286",
					"max": "",
					"min": "",
					"name": "timestamp",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "date"
				},
				{
					"cascadeDelete": true,
					"collectionId": "pbc_2089973043",
					"hidden": false,
					"id": "relation2073834773",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "launcher",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
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
					"id": "select2599078931",
					"maxSelect": 1,
					"name": "level",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "select",
					"values": [
						"info",
						"warning",
						"error",
						"debug"
					]
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1001261735",
					"max": 0,
					"min": 0,
					"name": "event",
					"pattern": "[a-zA-Z]+(?:_[a-zA-Z]+)*",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text3065852031",
					"max": 0,
					"min": 0,
					"name": "message",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "json2918445923",
					"maxSize": 0,
					"name": "data",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "json"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_329874174",
			"indexes": [],
			"listRule": "@request.auth.id != \"\"",
			"name": "launcher_logs",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": "@request.auth.id != \"\""
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_329874174")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
