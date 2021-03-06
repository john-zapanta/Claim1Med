SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetUserGroups] 
-- ***************************************************************************************************
-- Last modified on
-- 14-OCT-2014 ihms.0.0.1.0
-- *************************************************************************************************** 
(
	@user_name as varchar(10) = '',
	/*
	@filter as varchar(100) = '',

    @page int = 1, 
	@pagesize int = 1000000, 
	@row_count int = 0 OUTPUT, 
	@page_count int = 0 OUTPUT, 
	@sort varchar(200) = 'user_name',
	@order varchar(10) = 'asc',
	*/
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0;

	SELECT
		rtrim(id) as id,
		rtrim(group_name) as group_name
	FROM groups 
	WHERE id in (SELECT group_id FROM usergroups WHERE user_name = @user_name)

/*	WITH cte_usergroups (id, group_name, selected)
	AS (
		SELECT
			rtrim(id),
			rtrim(group_name),
			cast(1 as bit)
		FROM groups 
		WHERE id in (SELECT group_id FROM usergroups WHERE user_name = @user_name)
		UNION
		SELECT
			rtrim(id),
			rtrim(group_name),
			cast(0 as bit)
		FROM groups 
		WHERE id not in (SELECT group_id FROM usergroups WHERE user_name = @user_name)
	) SELECT * FROM cte_usergroups ORDER BY group_name; */
END


GO
