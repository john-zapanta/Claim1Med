package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "clients",
		Action: "clients",
		KeyName: "id",
		ListDataSource: "DBApp.GetClients",
		UpdateDataSource: "DBApp.AddClient",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			row["country"] = ""
			row["specialisation"] = ""
			row["notes"] = ""
			row["status_code"] = "A"
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClients", "GetClients", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
