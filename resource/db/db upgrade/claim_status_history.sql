USE [MEDICS51]
GO

DROP VIEW [dbo].[claim_status_history]
GO

CREATE VIEW [dbo].[claim_status_history]
AS
	SELECT        
		STATUS_ID AS id, 
		CLAIM_NO AS claim_id,
		STATUS AS status_code,
		InsertDate AS create_date, 
		InsertUser AS create_user,
		UpdateDate AS update_date, 
		UpdateUser AS update_user
	FROM dbo.tb_claim_status_history
GO
