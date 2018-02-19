package main

import (
	"net/http"
	"ibsi/crud"
	"ibsi/dbase"
)

func init() {

	crud.Handler(crud.CrudHandler{
		Name:             "service-status-details",
		Action:           "service-status-details",
		KeyName:          "code",
		ListDataSource:   "DBApp.GetServiceStatusCodes",
		UpdateDataSource: "DBApp.AddServiceStatusCodes",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			row["service_type"] = r.Form.Get("service_type")
			row["is_active"] = true
			row["is_system"] = true
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetServiceStatusCodes", "GetServiceStatusCodes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("service_type", "string", "in", 3, "")
		cmd.NewParameter("code", "string", "in", 7, "")
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "module")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
