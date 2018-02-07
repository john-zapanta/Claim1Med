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
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "addresses"
					s.Title = "Addresses"
					s.Icon = "addresses"
					s.Action = "admin"				
					s.Url = "app/addresses"
					// s.Run = "AddressesView"
					s.Params["name_id"] = id
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "contacts"
					s.Title = "Contacts"
					s.Icon = "contacts"
					s.Action = "admin"		
					s.Url = "app/contacts"	
					s.Params["name_id"] = id
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "products"
					s.Title = "Products"
					s.Icon = "products"
					s.Action = "product"		
					s.Url = "app/products"	
					s.Params["client_id"] = id
				})
			})
			
			utils.NewNavigatorItem(p.Nav, "financial", "Financial", func(item *utils.NavigatorItem) {
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "banks"
					s.Title = "Banks"
					s.Icon = "banks"
					s.Action = "admin"
					s.Params["name_id"] = id
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "floats"
					s.Title = "Floats"
					s.Icon = "floats"
					s.Action = "admin"				
					s.Params["client_id"] = id
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "case-fees"
					s.Title = "Case Fees"
					s.Icon = "case-fees"
					s.Action = "admin"				
					s.Params["client_id"] = id
				})
			})
			
			utils.NewNavigatorItem(p.Nav, "security", "Security", func(item *utils.NavigatorItem) {
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "authorisation"
					s.Title = "Authorisation"
					s.Icon = "authorisation"
					s.Action = "admin"
					s.Params["client_id"] = id
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "providers"
					s.Title = "Providers"
					s.Icon = "providers"
					s.Action = "admin"				
					s.Params["client_id"] = id
				})
			})
		},
	})
	
	
}
