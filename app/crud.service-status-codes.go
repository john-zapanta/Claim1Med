package main

import (
	"net/http"
	"ibsi/crud"
	"ibsi/dbase"
)

func init() {

	crud.Handler(crud.CrudHandler{
		Name:             "service-status-codes",
		Action:           "service-status-codes",
		KeyName:          "code",
		ListDataSource:   "DBApp.GetServiceTypes",
		UpdateDataSource: "",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			
		},
	})

	dbase.Connections["DBApp"].NewQuery("GetServiceTypes", func(cmd dbase.ICommand) {
		}, 
		func() string {
			return `
				select 
					distinct(module) as service_type,
					service_type as code
				from v_service_status_codes
			`
		},
	)
}
