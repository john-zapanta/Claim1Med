package main

import (
	"fmt"
	// "strings"
	"net/http"
	"github.com/gorilla/mux"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"ibsi/dbase"
	// "ibsi/system"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	
	type PageData struct {
		CustomerService dbase.TDataTableRows `json:"data"`
	}
	
	template.NewController(template.Controller {
		Pid: "plan",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/plan/{keyid:new}/{productcode:[0-9,A-Z,a-z]+\\/?}")
			ts.Add("/app/plan/{keyid:[0-9,A-Z,a-z]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			vars := mux.Vars(r)
			
			code := ""
			
			if vars["keyid"] != "new" {
				code = vars["keyid"]
			}
						
			dbPlan := dbase.Connections["DBApp"].OpenDataTable("GetPlans", dbase.TParameters{"code": code, "action":10, "visit_id": vid})
			fmt.Println("keyid", vars["keyid"])

			if code == "" {
				p.Title = "New Plan"
				p.Nav.PageTitle = "New Plan"
				p.Nav.WindowTitle = "New Plan"
			} else {
				p.Title = "Plan: " + dbPlan.Get("code").(string)
				p.Nav.PageTitle = "Plan: " + dbPlan.Get("code").(string)
				p.Nav.WindowTitle = dbPlan.Get("plan_name").(string)
			}
			
			if code == "" {
				dbPlan.Add(dbase.TDataTableRow{
					"code": "",
					"product_code": "",
					"plan_name": "",
				})
			}
			
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(code == "", 1, 0),
				"data": dbPlan.GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "plan", "Plan", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Title = "Details"
					s.Icon = "table-edit"
					s.Action = "admin"
					// s.Url = "app/product"
					// s.Css = "*"
					// s.Run = "ClaimDetailsView"
					s.Params["plan_code"] = code
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "history"
					s.Title = "Plan History"
					s.Icon = "timetable"
					s.Action = "admin"				
					s.Params["plan_code"] = code
				})
			})
		},
	})
	
	
}
