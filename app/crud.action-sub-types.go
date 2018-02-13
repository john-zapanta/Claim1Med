package main

import (
	"net/http"
	"ibsi/crud"
	"ibsi/dbase"
)

func init() {

	crud.Handler(crud.CrudHandler{
		Name:             "action-sub-types",
		Action:           "action-sub-types",
		KeyName:          "code",
		ListDataSource:   "DBApp.GetActionSubTypes",
		UpdateDataSource: "DBApp.AddActionSubTypes",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			row["action_type"] = r.Form.Get("action_type")
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetActionSubTypes", "GetActionSubTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("action_type", "string", "in", 3, "")
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddActionSubTypes", "AddActionSubTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("action_type", "string", "inout", 3, "")
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("action_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
