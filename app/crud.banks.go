package main

import (
	"net/http"
	// "github.com/gorilla/mux"
	"ibsi/crud"
	"ibsi/dbase"
	"ibsi/utils"
)

func init() {

	crud.Handler(crud.CrudHandler{
		Name:             "banks",
		Action:           "banks",
		Path:             "",
		KeyName:          "id",
		ListDataSource:   "DBApp.GetBanks",
		UpdateDataSource: "DBApp.AddBanks",
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

	dbase.Connections["DBApp"].NewCommand("GetBanks", "GetBanks", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddBanks", "AddBanks", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("bank_name", "string", "in", 34, "")
		cmd.NewParameter("sort_code", "string", "in", 11, "")
		cmd.NewParameter("swift_code", "string", "in", 11, "")
		cmd.NewParameter("bank_address1", "string", "in", 34, "")
		cmd.NewParameter("bank_address2", "string", "in", 34, "")
		cmd.NewParameter("bank_address3", "string", "in", 34, "")
		cmd.NewParameter("bank_country_code", "string", "in", 2, "")
		cmd.NewParameter("beneficiary_name", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_bank_account", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_address1", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_address2", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_address3", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_country_code", "string", "in", 2, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
