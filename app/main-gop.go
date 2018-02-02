package main

import (
	// "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"ibsi/dbase"
)

func init() {
	template.NewController(template.Controller {
		Pid: "gop",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/gop/{keyid:new}/{type}/{claimid:[0-9]+}")
			ts.Add("/app/gop/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			vars := mux.Vars(r)

			var id int64 = 0
			var serviceType, serviceSubType string = "GOP", ""
			
			if vars["keyid"] != "new" {
				id = utils.StrToInt64(vars["keyid"])
			} else {
				serviceSubType = vars["type"]
			}

			// dbService := dbase.Connections["DBApp"].OpenDataTable("GetService", dbase.TParameters{"id":id, "service_type":serviceType, "visit_id":vid})
			dbService := dbase.Connections["DBApp"].OpenDataTable("GetGop", dbase.TParameters{"id":id, "sub_type":serviceSubType, "visit_id":vid})

			// dbServiceSubData := dbase.Connections["DBApp"].OpenDataTable("GetServiceSubData", dbase.TParameters{"id":id, "service_type":serviceType, "sub_type":serviceSubType, "visit_id":vid})
			dbGopCalculationDates := dbase.Connections["DBApp"].OpenDataTable("GetGopCalculationDates", dbase.TParameters{"id":id, "visit_id":vid})
			if dbGopCalculationDates.RowCount() == 0 {
				// ...
			}

			dbGopEstimates := dbase.Connections["DBApp"].OpenDataTable("GetGopEstimates", dbase.TParameters{"id":id, "visit_id":vid})
			if dbGopEstimates.RowCount() == 0 {
				dbGopEstimates.Add(dbase.TDataTableRow{
					"id": id,
					"average_cost": 0,
					"average_los": 0,
					"estimated_cost": 0,
					"estimated_los": 0,
					"estimated_provider_cost": 0,
					"estimated_provider_los": 0,
				})
			}

			if id == 0 {
				p.Title = "New GOP"
				p.Nav.PageTitle = "New GOP"
				p.Nav.WindowTitle = "New GOP"
			} else {
				p.Title = dbService.Get("service_no").(string)
				p.Nav.PageTitle = dbService.Get("service_no").(string)
				p.Nav.WindowTitle = dbService.Get("service_no").(string)
			}

			// dbService.GetRows()
			// dbServiceSubType.GetRows()
			// dbGopEstimates.GetRows()
			// dbGopCalculationDates.GetRows()
			// dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows()
			// fmt.Println(dbService.Get("service_display_name").(string))
			// return
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(id == 0, 1, 0),
				"service_id": id,
				"module_type": serviceType,
				"service_sub_type": serviceSubType,
				// "service_type_name": dbServiceSubType.Get("display_name").(string),
				"service_type_name": dbService.Get("service_display_name").(string),
				"data": dbService.GetRows(),
				// "sub_type_data": dbServiceSubType.GetRows(),
				// "sub_type_data": dbServiceSubData.GetRows(),
				"estimates": dbGopEstimates.GetRows(),
				"calculation_dates": dbGopCalculationDates.GetRows(),
				"countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
			}
			
			description := dbService.Get("service_display_name").(string)

			utils.NewNavigatorItem(p.Nav, "service", description, func(item *utils.NavigatorItem) {

				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Action = "admin"
					s.Title = "Details"
					s.Description = description
					s.Icon = "table-edit"
					s.Css = "*"
					s.Run = "ServiceDetailsView"
					s.Params["service_id"] = id
				})

				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "breakdown"
					s.Action = "admin"
					s.Title = "Breakdown"
					s.Icon = "table-edit"
					s.Url = "app/service-breakdown"
					s.Params["service_id"] = id
				})

				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "documents"
					s.Title = "Documents"
					s.Icon = "documents"
					s.Action = "admin"
					s.Url = "app/claim-documents"
					s.Params["claim_id"] = 0
					s.Params["service_id"] = id
				})

				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "notes"
					s.Title = "Notes"
					s.Icon = "notes"
					s.Action = "admin"
					s.Url = "app/claim-notes"
					s.Params["type"] = "S"
					s.Params["claim_id"] = 0
					s.Params["service_id"] = id
				})

				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "audit"
					s.Action = "admin"				
					s.Title = "Audit Log"
					s.Icon = "timetable"
					s.Action = "admin"				
					s.Url = "app/claim-audit-logs"
					s.Params["claim_id"] = 0
					s.Params["service_id"] = id
				})
			})
		},
	})


}
