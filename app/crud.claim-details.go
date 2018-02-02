package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.CommandHandle(crud.CommandHandler {
		Name: "can-delete-claim",
		Action: "claim",
		DataSource: "DBApp.CanDeleteClaim",
		OnInitCommand: func(cmd crud.CommandHandler, params dbase.TParameters, w http.ResponseWriter, r *http.Request) {
			params["id"] = r.Form.Get("id")
			params["change_plan"] = r.Form.Get("change_plan")
		},
	});

	crud.CommandHandle(crud.CommandHandler {
		Name: "validate-action",
		Action: "claim",
		DataSource: "DBApp.ValidateAction",
		OnInitCommand: func(cmd crud.CommandHandler, params dbase.TParameters, w http.ResponseWriter, r *http.Request) {
			params["id"] = r.Form.Get("id")
			params["string_value"] = r.Form.Get("string_value")
			params["key_id"] = r.Form.Get("key_id")
			params["action_name"] = r.Form.Get("action_name")
		},
	});

	crud.Handler(crud.CrudHandler {
		Name: "claim-details",
		Action: "claim-details",
		KeyName: "id",
		ListDataSource: "DBApp.GetClaim",
		UpdateDataSource: "DBApp.AddClaim",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["status_code_id"] = 10
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClaim", "GetClaim", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
	
	dbase.Connections["DBApp"].NewCommand("AddClaim", "AddClaim", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("claim_no" , "string", "inout", 200, 0)
		cmd.NewParameter("member_id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("policy_id", "int", "in", 0, 0)
		cmd.NewParameter("client_id", "int", "in", 0, 0)
		cmd.NewParameter("product_code", "string", "in", 10, "")
		cmd.NewParameter("plan_code", "string", "in", 15, "")
		cmd.NewParameter("plan_code2", "string", "in", 15, "")
		cmd.NewParameter("sub_product", "string", "in", 15, "")
		cmd.NewParameter("claim_type", "string", "in", 4, "")
		cmd.NewParameter("base_currency_code", "string", "in", 3, "")
		cmd.NewParameter("client_currency_code" , "string", "in", 3, "")
		cmd.NewParameter("eligibility_currency_code" , "string", "in", 3, "")
		cmd.NewParameter("notification_date", "datetime", "in", 0, nil)
		cmd.NewParameter("case_owner" , "string", "in", 10, "")
		cmd.NewParameter("status_code" , "string", "inout", 1, "")
		cmd.NewParameter("status" , "string", "inout", 20, "")
		cmd.NewParameter("hcm_reference" , "string", "in", 15, "")
		cmd.NewParameter("reference_no1" , "string", "in", 15, "")
		cmd.NewParameter("reference_no2" , "string", "in", 15, "")
		cmd.NewParameter("reference_no3" , "string", "in", 15, "")
		cmd.NewParameter("third_party", "int", "in", 0, 0)

		cmd.NewParameter("is_accident", "int", "in", 0, 0)
		cmd.NewParameter("accident_date", "datetime", "in", 0, nil)
		cmd.NewParameter("accident_code", "string", "in", 10, "")
		
		cmd.NewParameter("is_preexisting", "int", "in", 0, 0)
		cmd.NewParameter("first_symptom_date", "datetime", "in", 0, nil)
		cmd.NewParameter("travel_return_date", "datetime", "in", 0, nil)

		cmd.NewParameter("travel_departure_date", "datetime", "in", 0, nil)
		cmd.NewParameter("first_consultation_date", "datetime", "in", 0, nil)
		
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
	
	dbase.Connections["DBApp"].NewCommand("CanDeleteClaim", "CanDeleteClaim", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("change_plan", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
	
	dbase.Connections["DBApp"].NewCommand("ValidateAction", "ValidateAction", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("string_value", "string", "in", 30, "")
		cmd.NewParameter("key_id", "int", "in", 0, 0)
		cmd.NewParameter("action_name", "string", "inout", 100, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
