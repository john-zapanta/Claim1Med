USE [MEDICS51]
GO

DROP VIEW [dbo].[medical_history_notes]
GO

CREATE VIEW [dbo].[medical_history_notes]
AS
	SELECT        
		IP_ID AS id, 
		MED_HIST AS medical_history_notes,
		UpdateDate AS update_date, 
		UpdateUser AS update_user
	FROM dbo.tb_medical_history_notes
GO
