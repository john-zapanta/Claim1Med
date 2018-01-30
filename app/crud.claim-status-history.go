package main

import (
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "claim-status-history",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetClaimStatusHistory",
		UpdateDataSource: "DBApp.AddClaimStatusHistory",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimStatusHistory", "GetClaimStatusHistory", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddClaimStatusHistory", "AddClaimStatusHistory", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("status_code", "string", "in", 1, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
