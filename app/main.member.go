package main

import (
	"net/http"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/dbase"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	
	type PageData struct {
		CustomerService dbase.TDataTableRows `json:"data"`
	}
	
	template.NewController(template.Controller {
		Pid: "member",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/member/new/{clientid:new\\/?}")
			ts.Add("/app/member/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData2: func(p *template.Page, vid int64, vars map[string]string, r *http.Request) {
			var id int64 = 0
			
			if vars["keyid"] != "new" {
				id = utils.StrToInt64(vars["keyid"])
			}
						
			dbMember := dbase.Connections["DBApp"].OpenDataTable("GetMember", dbase.TParameters{"id": id, "visit_id": vid})
			
			if id == 0 {
				p.Title = "New Member"
				p.Nav.PageTitle = "New Member"
				p.Nav.WindowTitle = "New Member"
			} else {
				p.Title = "Member: " + dbMember.Get("name").(string)
				p.Nav.PageTitle = "Member: " + dbMember.Get("name").(string)
				p.Nav.WindowTitle = dbMember.Get("name").(string)
			}
			
			if id == 0 {
				dbMember.Add(dbase.TDataTableRow{
					"id": 0,
					"name": "",
				})
			}
			
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(id == 0, 1, 0),
				"data": dbMember.GetRows(),
				"countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "member", "Member", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Title = "Details"
					s.Icon = "table-edit"
					s.Action = "admin"
					// s.Url = "app/claim-details"
					// s.Css = "*"
					// s.Run = "ClaimDetailsView"
					s.Params["member_id"] = id
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
