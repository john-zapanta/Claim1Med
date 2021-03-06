SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[v_service_status_history]
AS
	SELECT  
		c.*,
		s.status,
		ss.sub_status
	FROM dbo.service_status_history c
	JOIN services i on c.service_id = i.id
	LEFT OUTER JOIN invoice_status s on i.service_type = s.service_type and c.status_code = s.code
	LEFT OUTER JOIN sub_status_codes ss on i.service_type = ss.service_type and c.status_code = ss.status_code and c.sub_status_code = ss.sub_status_code
	--LEFT OUTER JOIN dbo.action_types ac ON c.action_type_code = ac.code
	--LEFT OUTER JOIN dbo.ACTIONS a ON C.action_type_code = a.action_type AND c.action_code = a.code
	--ACTION_CLASS AND C.ACTION_CODE = a.ACTION_CODE


GO
