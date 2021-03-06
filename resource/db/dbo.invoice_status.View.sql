SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[invoice_status]
as
	SELECT
		MODULE_ID as service_type,
		STATUS_CODE as code,
		STATUS_DESC as status,
		INTERNAL as is_internal
	FROM MEDICS40JKT.dbo.F_INVOICE_STATUS()
GO
