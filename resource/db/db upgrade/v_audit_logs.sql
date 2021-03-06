DROP VIEW [dbo].[v_audit_logs]
GO

CREATE VIEW [dbo].[v_audit_logs] 
-- ***************************************************************************************************
-- Last modified on
-- 06-OCT-2017
-- *************************************************************************************************** 
AS
	SELECT
		l.*,
		m.module_name,
		t.description as log_name,
		ISNULL(s.service_no, c.claim_no) as reference_no,
		ISNULL(u.name, l.create_user) as create_user_name
	FROM auditlogs l
	JOIN auditlog_types t on l.code = t.code
	JOIN f_modules() m on l.module = m.module_id
	LEFT OUTER JOIN users u on l.create_user = u.user_name
	LEFT OUTER JOIN claims c ON l.claim_id = c.id
	LEFT OUTER JOIN services s ON l.service_id = s.id
GO


