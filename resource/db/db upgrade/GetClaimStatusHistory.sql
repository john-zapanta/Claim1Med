DROP PROCEDURE [dbo].[GetClaimStatusHistory]
GO
/****** Object:  StoredProcedure [dbo].[GetClaimDiagnosis]    Script Date: 9/27/2017 1:42:53 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetClaimStatusHistory] 
-- ***************************************************************************************************
-- Last modified on
-- 27-SEP-2017
-- *************************************************************************************************** 
(
    @claim_id int = 0, 
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	SELECT 
		s.id,
		s.claim_id,
		s.status_code,
		cs.status,
		s.create_date AS status_date,
		s.create_user AS user_name,
		u.name AS user_full_name
	FROM claim_status_history s
	JOIN claim_status cs
		ON s.status_code = cs.code
	LEFT OUTER JOIN users u
		ON s.create_user = u.user_name
	WHERE s.claim_id = @claim_id
	ORDER BY s.create_date DESC
END
GO
