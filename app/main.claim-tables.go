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
			
			p.Nav.CustomData = map[string]interface{}{
				// "flagtypes": dbase.Connections["DBApp"].OpenDataTable("GetFlagTypes", dbase.TParameters{"lookup":1, "visit_id":vid}).GetRows(),
				// "servicestatus": dbase.Connections["DBApp"].OpenDataTable("GetServiceStatus", dbase.TParameters{"lookup":1, "visit_id":vid}).GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "tables1", "Claim & Service", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "claim-types"
					s.Action = "admin"				
					s.Title = "Claim Types"
					s.Icon = "hospital"
					s.Url = "app/claim-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "service-types"
					s.Action = "admin"
					s.Title = "Service Types"
					s.Icon = "pill"
					s.Url = "app/service-types"
				})
			})
			
			utils.NewNavigatorItem(p.Nav, "tables2", "General", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "case-fees-types"
					s.Action = "admin"				
					s.Title = "Case Fees Types"
					s.Icon = "currency-usd"
					s.Url = "app/case-fees-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "customer-service-types"
					s.Action = "admin"				
					s.Title = "Customer Service Types"
					s.Icon = "deskphone"
					s.Url = "app/customer-service-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "gop-types"
					s.Action = "admin"				
					s.Title = "Guarantee of Payment Types"
					s.Icon = "currency-usd"
					s.Url = "app/gop-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "noc-types"
					s.Action = "admin"				
					s.Title = "Notification of Claims Types"
					s.Icon = "alert-circle-outline"
					s.Url = "app/noc-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "recovery-types"
					s.Action = "admin"				
					s.Title = "Recovery Types"
					s.Icon = "autorenew"
					s.Url = "app/recovery-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "cost-containment-types"
					s.Action = "admin"				
					s.Title = "Cost Containment Types"
					s.Icon = "inbox-arrow-down"
					s.Url = "app/cost-containment-types"
				})
			})
			
			utils.NewNavigatorItem(p.Nav, "flags", "Flags", func(item *utils.NavigatorItem) {
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "flag-types"
					s.Action = "admin"				
					s.Title = "Flag Types"
					s.Icon = "flag"
					s.Url = "app/flag-types"
				})
			})
			
			utils.NewNavigatorItem(p.Nav, "codes-notes", "Codes & Notes", func(item *utils.NavigatorItem) {
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "service-status-codes"
					s.Action = "admin"				
					s.Title = "Service Status Codes"
					s.Icon = "checkbox-multiple-marked-some"
					s.Url = "app/service-status-codes"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "claim-notes-types"
					s.Action = "admin"				
					s.Title = "Claim Notes"
					s.Icon = "pencil-box-outline"
					s.Url = "app/claim-notes-types"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "action-types"
					s.Action = "admin"				
					s.Title = "Actions"
					s.Icon = "thumb-up-outline"
					s.Url = "app/action-types"
				})
			})
			
			utils.NewNavigatorItem(p.Nav, "log", "Log", func(item *utils.NavigatorItem) {
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "auditlog-types"
					s.Action = "admin"				
					s.Title = "Audit Log"
					s.Icon = "table-edit"
					s.Url = "app/auditlog-types"
				})
			})
		},
	})
}
