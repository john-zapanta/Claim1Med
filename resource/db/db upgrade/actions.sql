DROP VIEW [dbo].[actions]
GO


CREATE VIEW [dbo].[actions]
AS
	SELECT        
		ACTION_CLASS AS action_type,
		ACTION_CODE AS code,
		ACTION_NAME AS action_name,
		IS_ACTIVE AS is_active,
		InsertUser AS create_user,
		InsertDate AS create_date,
		UpdateUser AS update_user,
		UpdateDate AS update_date
	FROM dbo.tb_actions



GO
