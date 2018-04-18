DROP PROCEDURE [dbo].[GetAuditlogTypes]
GO

CREATE PROCEDURE [dbo].[GetAuditlogTypes]
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
	@code char(3) = NULL,
	@filter as varchar(100) = '',
	@action int = 0, -- 0:list, 1:lookup, 10:for editing, 20:for new record, 50:fetch updated data

    @page int = 1, 
	@pagesize int = 0, 
	@row_count int = 0 OUTPUT, 
	@page_count int = 0 OUTPUT, 
	@sort varchar(200) = 'code',
	@order varchar(10) = 'asc',

    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	IF @action in (10,20,50)
	BEGIN
		SELECT * FROM v_auditlog_types WHERE code = @code
	END ELSE
	BEGIN

		--DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
		DECLARE @user_id AS int = 0
		DECLARE @columns varchar(100) = '*'  

		IF @action = 1
		BEGIN
			SET @pagesize = 1000000
		END
    
		EXEC RunSimpleDynamicQuery
			@source = 'v_auditlog_types',
			@columns = @columns,

			@filter = @filter,
			@where = '[description] like @filter',

			@page = @page, 
    		@pagesize = @pagesize, 
    		@row_count = @row_count OUTPUT, 
    		@page_count = @page_count OUTPUT,
			@sort = @sort,
			@order = @order
	END
END















GO

