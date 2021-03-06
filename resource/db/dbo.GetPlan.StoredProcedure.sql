SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetPlan]
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
    @code varchar(10) = '', 
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
    
	SELECT
		p.*,
		r.product_name,
		r.client_id,
		n.full_name as client_name
	FROM plans p
	join products r on p.product_code = r.code
	join names n on r.client_id = n.id
	WHERE p.code = @code
END



GO
