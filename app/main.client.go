package main

import (
	// "fmt"
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
		Pid: "client",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/client/{keyid:new\\/?}")
			ts.Add("/app/client/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			vars := mux.Vars(r)
			
			var id int64 = 0
			
			if vars["keyid"] != "new" {
				id = utils.StrToInt64(vars["keyid"])
			}
						
			dbClient := dbase.Connections["DBApp"].OpenDataTable("GetClients", dbase.TParameters{"id": id, "action":10, "visit_id": vid})

			// fmt.Println(vars["keyid"])
			// return
			
			if id == 0 {
				p.Title = "New Client"
				p.Nav.PageTitle = "New Client"
				p.Nav.WindowTitle = "New Client"
			} else {
				p.Title = "Client: " + dbClient.Get("name").(string)
				p.Nav.PageTitle = "Claim: " + dbClient.Get("name").(string)
				p.Nav.WindowTitle = dbClient.Get("name").(string)
			}
			
			if id == 0 {
				dbClient.Add(dbase.TDataTableRow{
					"id": 0,
					"name": "",
				})
			}
			
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(id == 0, 1, 0),
				"data": dbClient.GetRows(),
				"countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "client", "Client", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Title = "Details"
					s.Icon = "table-edit"
					s.Action = "admin"
					// s.Url = "app/claim-details"
					// s.Css = "*"
					// s.Run = "ClaimDetailsView"
					s.Params["client_id"] = id
				})
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "call-log"
					// s.Title = "Call Logs"
					// s.Icon = "phone"
					// s.Action = "admin"				
					// s.Params["claim_id"] = id
				// })
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "history"
					// s.Title = "Case History"
					// s.Icon = "history"
					// s.Action = "admin"				
					// s.Url = "app/member-case-history"
					// s.Params["claim_id"] = id
					// s.Params["member_id"] = memberId
				// })
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "documents"
					// s.Title = "Documents"
					// s.Icon = "documents"
					// s.Action = "admin"
					// s.Url = "app/claim-documents"
					// s.Params["claim_id"] = id
					// s.Params["service_id"] = 0
				// })
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "notes"
					// s.Title = "Notes"
					// s.Icon = "notes"
					// s.Action = "admin"				
					// s.Url = "app/claim-notes"
					// s.Params["type"] = "C"
					// s.Params["claim_id"] = id
					// s.Params["service_id"] = 0
				// })
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "log"
					// s.Title = "Audit Log"
					// s.Icon = "timetable"
					// s.Action = "admin"				
					// s.Url = "app/claim-audit-logs"
					// s.Params["claim_id"] = id
					// s.Params["service_id"] = 0
				// })
			})
		},
	})
	
	
}
