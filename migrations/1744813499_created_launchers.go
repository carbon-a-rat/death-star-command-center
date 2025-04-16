package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"authAlert": {
				"emailTemplate": {
					"body": "<p>Hello,</p>\n<p>We noticed a login to your {APP_NAME} account from a new location.</p>\n<p>If this was you, you may disregard this email.</p>\n<p><strong>If this wasn't you, you should immediately change your {APP_NAME} account password to revoke access from all other locations.</strong></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Login from a new location"
				},
				"enabled": true
			},
			"authRule": "",
			"authToken": {
				"duration": 604800
			},
			"confirmEmailChangeTemplate": {
				"body": "<p>Hello,</p>\n<p>Click on the button below to confirm your new email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-email-change/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Confirm new email</a>\n</p>\n<p><i>If you didn't ask to change your email address, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
				"subject": "Confirm your {APP_NAME} new email address"
			},
			"createRule": null,
			"deleteRule": null,
			"emailChangeToken": {
				"duration": 1800
			},
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
					"cost": 0,
					"hidden": true,
					"id": "password901924565",
					"max": 0,
					"min": 8,
					"name": "password",
					"pattern": "",
					"presentable": false,
					"required": true,
					"system": true,
					"type": "password"
				},
				{
					"autogeneratePattern": "[a-zA-Z0-9]{50}",
					"hidden": true,
					"id": "text2504183744",
					"max": 60,
					"min": 30,
					"name": "tokenKey",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"exceptDomains": null,
					"hidden": false,
					"id": "email3885137012",
					"name": "email",
					"onlyDomains": null,
					"presentable": false,
					"required": false,
					"system": true,
					"type": "email"
				},
				{
					"hidden": false,
					"id": "bool1547992806",
					"name": "emailVisibility",
					"presentable": false,
					"required": false,
					"system": true,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool256245529",
					"name": "verified",
					"presentable": false,
					"required": false,
					"system": true,
					"type": "bool"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text4217917526",
					"max": 0,
					"min": 3,
					"name": "codename",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1579384326",
					"max": 0,
					"min": 0,
					"name": "name",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "select2063623452",
					"maxSelect": 1,
					"name": "status",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "select",
					"values": [
						"available",
						"busy",
						"offline"
					]
				},
				{
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
				},
				{
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
				},
				{
					"cascadeDelete": false,
					"collectionId": "_pb_users_auth_",
					"hidden": false,
					"id": "relation1109442351",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "current_user",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "_pb_users_auth_",
					"hidden": false,
					"id": "relation3609378609",
					"maxSelect": 999,
					"minSelect": 0,
					"name": "allowed_users",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_1889285177",
					"hidden": false,
					"id": "relation1216748164",
					"maxSelect": 999,
					"minSelect": 0,
					"name": "rockets_loaded",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
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
			"fileToken": {
				"duration": 180
			},
			"id": "pbc_2089973043",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_QSTGS5eF42` + "`" + ` ON ` + "`" + `launchers` + "`" + ` (` + "`" + `codename` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_tokenKey_pbc_2089973043` + "`" + ` ON ` + "`" + `launchers` + "`" + ` (` + "`" + `tokenKey` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_email_pbc_2089973043` + "`" + ` ON ` + "`" + `launchers` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''"
			],
			"listRule": null,
			"manageRule": null,
			"mfa": {
				"duration": 1800,
				"enabled": false,
				"rule": ""
			},
			"name": "launchers",
			"oauth2": {
				"enabled": false,
				"mappedFields": {
					"avatarURL": "",
					"id": "",
					"name": "",
					"username": ""
				}
			},
			"otp": {
				"duration": 180,
				"emailTemplate": {
					"body": "<p>Hello,</p>\n<p>Your one-time password is: <strong>{OTP}</strong></p>\n<p><i>If you didn't ask for the one-time password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "OTP for {APP_NAME}"
				},
				"enabled": false,
				"length": 8
			},
			"passwordAuth": {
				"enabled": true,
				"identityFields": [
					"codename"
				]
			},
			"passwordResetToken": {
				"duration": 1800
			},
			"resetPasswordTemplate": {
				"body": "<p>Hello,</p>\n<p>Click on the button below to reset your password.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-password-reset/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Reset password</a>\n</p>\n<p><i>If you didn't ask to reset your password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
				"subject": "Reset your {APP_NAME} password"
			},
			"system": false,
			"type": "auth",
			"updateRule": null,
			"verificationTemplate": {
				"body": "<p>Hello,</p>\n<p>Thank you for joining us at {APP_NAME}.</p>\n<p>Click on the button below to verify your email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-verification/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Verify</a>\n</p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
				"subject": "Verify your {APP_NAME} email"
			},
			"verificationToken": {
				"duration": 259200
			},
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2089973043")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
