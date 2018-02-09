package main

import (
	"net/http"
	"ibsi/utils"
	"ibsi/template"
	//"ibsi/session"
	"ibsi/dbase"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	
	type PageData struct {
		CustomerService dbase.TDataTableRows `json:"data"`
	}
	
	template.NewController(template.Controller {
		Pid: "claim-tables",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/{pid:claim-tables\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			//vid := session.GetVisitorId(r)
			
			p.Title = "Claim Tables"
			p.Nav.WindowTitle = "Claim Tables"
			
			// p.Nav.CustomData = map[string]interface{}{
				// "countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
				// "specialisation": dbase.Connections["DBApp"].OpenDataTable("GetDoctorSpecialisation", dbase.TParameters{"lookup":1, "visit_id":vid}).GetRows(),
			// }
			
			utils.NewNavigatorItem(p.Nav, "tables1", "Claim & Service", func(item *utils.NavigatorItem) {
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "doctors"
					// s.Action = "admin"				
					// s.Title = "Doctors"
					// s.Icon = "doctor"
					// s.Url = "app/providers"
					// s.Run = "ProviderMasterDetailView"
					// s.Css = "*"
					// s.Params["provider_type"] = "D"
				// })
				
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "members"
					// s.Action = "admin"
					// s.Title = "Claims Processing"
					// s.Icon = "users"
					// s.Url = "app/claims-entry"
				// })
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "enquiries"
					// s.Action = "admin"
					// s.Title = "Enquiries"
					// s.Icon = "search"
				// })
				
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "reports"
					// s.Action = "admin"
					// s.Title = "Reports"
					// s.Icon = "table"
					// s.Url = "engine/sys-reports"
				// })
			})
			
			utils.NewNavigatorItem(p.Nav, "tables2", "General", func(item *utils.NavigatorItem) {
				
			})
			
			utils.NewNavigatorItem(p.Nav, "flags", "Flags", func(item *utils.NavigatorItem) {
				
			})
			
			utils.NewNavigatorItem(p.Nav, "codes-notes", "Codes & Notes", func(item *utils.NavigatorItem) {
				
			})
			
			utils.NewNavigatorItem(p.Nav, "log", "Log", func(item *utils.NavigatorItem) {
				
			})
		},
	})
}
