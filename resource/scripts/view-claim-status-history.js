// ****************************************************************************************************
// Last modified on
// 27-SEP-2017
// ****************************************************************************************************
//==================================================================================================
// File name: view-claim-diagnosis-summary.js
//==================================================================================================
function ClaimStatusHistoryView(params) {
	var ChangeStatus = function(status, new_status) {
		var params = {
			mode: "new",
			data: JSON.stringify([{
				claim_id: desktop.dbClaim.getKey(),
				status_code: status
				// new_status_code: new_status
			}])
		};

		desktop.Ajax(self, "/app/get/update/claim-status-history", params, function(result) {
			if (result.status == 0) {
				location.reload();
			}
		})
	};
	
	return new jGrid($.extend(params, {
		paintParams: {
			css: "claim-status",
			toolbar: {theme: "svg"}
		},
		init: function(grid, callback) {			
			grid.Events.OnInit.add(function(grid) {
				grid.optionsData.url = "app/claim-status-history";
				
				grid.options.viewType = "cardview";
				grid.options.horzScroll = true;
				grid.options.allowSort = false;
				grid.options.showPager = false;
				grid.options.showMasterDetail = false;
				grid.options.editNewPage = false;
				grid.options.hideHeader = true;
				// grid.options.showSummary = true;
				
				grid.search.visible = false;
				
				grid.Events.OnInitDataRequest.add(function(grid, dataParams) {
					dataParams
						.addColumn("claim_id", params.claim_id)
				});
				
				grid.Events.OnInitData.add(function(grid, data) {
					data.Columns
						.setprops("id", {label:"#", numeric:true, key: true})
						.setprops("status_code", {label:"Code"})
						.setprops("status", {label:"Status"})
						.setprops("status_date", {label:"Date", type:"date", format:"datetime"})
						.setprops("user_full_name", {label:"User"})
				});
				
				grid.Events.OnInitCard.add(function(grid, card) {
					grid.dataset.gotoKey(parseInt(card.attr("row-id")));
					card.attr("x-status", grid.dataset.get("status_code"));
					
					// CreateElement("div", card).addClass("status").html(grid.dataset.get("status"))
					CreateElementEx("div", card, function(container) {
						CreateElement("div", container).addClass("status").html(grid.dataset.get("status"));
						
						CreateElementEx("div", container, function(container) {
							CreateElement("div", container).html("Created by:");
							CreateElement("div", container).html(grid.dataset.text("user_full_name"));
						}, "log");
						
						CreateElementEx("div", container, function(container) {
							CreateElement("div", container).html("Created on:");
							CreateElement("div", container).html(grid.dataset.formatDateTime("status_date", "MMMM d, yyyy"));
						}, "log");
					}, "card-row");
					
					// CreateElementEx("div", card, function(container) {
						// CreateElementEx("div", container, function(container) {
							// CreateElement("div", container).html("Created by")
							// CreateElement("div", container).html(grid.dataset.text("user_full_name"))
						// }, "user");
						// CreateElementEx("div", container, function(container) {
							// CreateElement("div", container).html("Created on")
							// CreateElement("div", container).html(grid.dataset.formatDateTime("status_date", "MMMM d, yyyy"))
						// }, "user");
					// }, "log");
					
					// CreateElementEx("div", card, function(container) {
						// CreateElement("span", container).html(grid.dataset.get("status"))
					// }, "status")
					
					// CreateElement("div", card)
						// .addClass("user")
						// .html(grid.dataset.get("user_full_name"))
					
					// CreateElement("div", card)
						// .addClass("date")
						// .html(grid.dataset.Columns.get("status_date").formatDateTime("dd-MMM-yyyy"))
				});
				
				grid.Events.OnInitColumns.add(function(grid) {
					// grid.NewColumn({fname: "diagnosis_code", width: 150, fixedWidth:true, allowSort: true});
				});
				
				grid.Events.OnInitToolbar.add(function(view, toolbar) {
					var status = desktop.dbClaim.get("status_code");
					if (status === "O" || status === "N") {
						toolbar.NewDropDownConfirmItem({
							id: "close",
							icon: "claim-close",
							color: "black",
							hint: "Close Claim",
							title: "Close Claim",
							subTitle: "Please confirm to close claim.",
							// dataBind: desktop.dbClaim,
							// dataEvent: function(dataset, button) {
								// button.show(!dataset.editing);
							// },
							confirm: function() {
								ChangeStatus("C");
							}
						});
					
						toolbar.NewDropDownConfirmItem({
							id: "decline",
							icon: "claim-decline",
							color: "firebrick",
							hint: "Decline Claim",
							title: "Decline Claim",
							subTitle: "Please confirm to decline claim.",
							// dataBind: desktop.dbClaim,
							// dataEvent: function(dataset, button) {
								// button.show(!dataset.editing);
							// },
							confirm: function() {
								ChangeStatus("D");
							}
						});
					}
					
					if (status != "O" && status != "N") {
						toolbar.NewDropDownConfirmItem({
							id: "open",
							icon: "claim-open",
							color: "#1AB394",
							hint: "Re-Open Claim",
							title: "Re-Open Claim",
							subTitle: "Please confirm to re-open claim.",
							// dataBind: desktop.dbClaim,
							// dataEvent: function(dataset, button) {
								// button.show(!dataset.editing);
							// },
							confirm: function() {
								ChangeStatus("O");
							}
						});
					}
				});
			});
		}
	}));
};
