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
		Pid: "product",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/product/{keyid:new}/{clientid:[0-9]+\\/?}")
			// ts.Add("/app/product/{keyid\\/?}")
			// ts.Add("/app/product/{keyid}")
			// ts.Add("/app/product/{keyid:[A-Z,a-z]+\\/?}")
			ts.Add("/app/product/{keyid:[A-Z,a-z]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			vars := mux.Vars(r)
			
			code := ""
			
			if vars["keyid"] != "new" {
				code = vars["keyid"]
			}
						
			dbProduct := dbase.Connections["DBApp"].OpenDataTable("GetProducts", dbase.TParameters{"code": code, "action":10, "visit_id": vid})
			// fmt.Println("keyid", vars["keyid"])

			if code == "" {
				p.Title = "New Product"
				p.Nav.PageTitle = "New Product"
				p.Nav.WindowTitle = "New Product"
			} else {
				p.Title = "Product: " + dbProduct.Get("code").(string)
				p.Nav.PageTitle = "Product: " + dbProduct.Get("code").(string)
				p.Nav.WindowTitle = dbProduct.Get("product_name").(string)
			}
			
			if code == "" {
				dbProduct.Add(dbase.TDataTableRow{
					"code": "",
					"product_name": "",
				})
			}
			
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(code == "", 1, 0),
				"data": dbProduct.GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "product", "Product", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Title = "Details"
					s.Icon = "table-edit"
					s.Action = "admin"
					// s.Url = "app/product"
					// s.Css = "*"
					// s.Run = "ClaimDetailsView"
					s.Params["product_code"] = code
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "plans"
					s.Title = "Plans"
					s.Icon = "table"
					s.Action = "admin"				
					s.Params["product_code"] = code
				})
			
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
