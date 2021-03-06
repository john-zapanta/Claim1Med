SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetUser] 
-- ***************************************************************************************************
-- Last modified on
-- 14-OCT-2014 ihms.0.0.1.0
-- *************************************************************************************************** 
(
    @id int = 0, 
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
    
	SELECT
		id,
		user_name,
		name,
		designation,
		phone_no,
		is_active,
		is_supervisor,
		cast(dbo.F_UserGroups(user_name, 1) as varchar(255)) as groups
	FROM users
	WHERE id = @id    
END


GO
