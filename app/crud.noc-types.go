package main

import (
	"net/http"
	"ibsi/crud"
	"ibsi/dbase"
)

func init() {

	crud.Handler(crud.CrudHandler{
		Name:             "noc-types",
		Action:           "noc-types",
		KeyName:          "code",
		ListDataSource:   "DBApp.GetNOCTypes",
		UpdateDataSource: "DBApp.AddNOCTypes",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			row["is_active"] = true
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetNOCTypes", "GetNOCTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddNOCTypes", "AddNOCTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
