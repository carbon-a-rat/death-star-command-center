package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1889285177")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text1024124636")

		// remove field
		collection.Fields.RemoveById("date4088452399")

		// remove field
		collection.Fields.RemoveById("select2063623452")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "number3113930206",
			"max": null,
			"min": 0,
			"name": "volume",
			"onlyInt": false,
			"presentable": false,
			"required": true,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "number1812159334",
			"max": null,
			"min": 0,
			"name": "mass",
			"onlyInt": false,
			"presentable": false,
			"required": true,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"hidden": false,
			"id": "number1039061966",
			"max": null,
			"min": 0,
			"name": "nozzle_diameter",
			"onlyInt": false,
			"presentable": false,
			"required": true,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"hidden": false,
			"id": "number729234583",
			"max": null,
			"min": 0,
			"name": "rocket_diameter",
			"onlyInt": false,
			"presentable": false,
			"required": true,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"hidden": false,
			"id": "number3610066113",
			"max": null,
			"min": 0,
			"name": "drag_coefficient",
			"onlyInt": false,
			"presentable": false,
			"required": true,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"cascadeDelete": false,
			"collectionId": "_pb_users_auth_",
			"hidden": false,
			"id": "relation568698700",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "id_owner",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2103600739",
			"hidden": false,
			"id": "relation516658094",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "id_manufacturer",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1889285177")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text1024124636",
			"max": 0,
			"min": 0,
			"name": "manufacturer",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "date4088452399",
			"max": "",
			"min": "",
			"name": "last_launch",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"hidden": false,
			"id": "select2063623452",
			"maxSelect": 1,
			"name": "status",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"Available",
				"Busy",
				"Offline"
			]
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("number3113930206")

		// remove field
		collection.Fields.RemoveById("number1812159334")

		// remove field
		collection.Fields.RemoveById("number1039061966")

		// remove field
		collection.Fields.RemoveById("number729234583")

		// remove field
		collection.Fields.RemoveById("number3610066113")

		// remove field
		collection.Fields.RemoveById("relation568698700")

		// remove field
		collection.Fields.RemoveById("relation516658094")

		return app.Save(collection)
	})
}
