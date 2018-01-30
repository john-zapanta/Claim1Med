package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {
	crud.Handler(crud.CrudHandler {
		Name: "member-medical-notes",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetMemberMedicalNotes",
		UpdateDataSource: "DBApp.AddMemberMedicalNotes",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["status_code_id"] = 10
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetMemberMedicalNotes", "GetMemberMedicalNotes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 1)
		cmd.NewParameter("claim_id", "int", "in", 0, 1)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddMemberMedicalNotes", "AddMemberMedicalNotes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("medical_history_notes", "string", "in", -1, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
