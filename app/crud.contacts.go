package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
	"ibsi/utils"
)

func init() {
	crud.Handler(crud.CrudHandler {
		Name: "contacts",
		Action: "contacts",
		Path: "",
		KeyName: "id",
		ListDataSource: "DBApp.GetContacts",
		UpdateDataSource: "DBApp.AddContacts",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			row["name_id"] = utils.StrToInt(r.Form.Get("name_id"))
			// row["address_type_name"] = ""
			// row["country"] = ""
			// utils.Println(vars)
			// row["notes"] = ""
			// row["status_code"] = "A"
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetContacts", "GetContacts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddContacts", "AddContacts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("title", "string", "in", 20, "")
		cmd.NewParameter("full_name", "string", "in", 50, "")
		cmd.NewParameter("department", "string", "in", 50, "")
		cmd.NewParameter("position", "string", "in", 50, "")
		cmd.NewParameter("phone", "string", "in", 20, "")
		cmd.NewParameter("mobile", "string", "in", 20, "")
		cmd.NewParameter("fax", "string", "in", 20, "")
		cmd.NewParameter("email", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
