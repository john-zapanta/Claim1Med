DROP PROCEDURE [dbo].[GetServiceStatus]
GO

CREATE PROCEDURE [dbo].[GetServiceStatus]
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
	@service_type char(3) = NULL, 
	@lookup int = 0,
	@filter as varchar(100) = '',

    @page int = 1, 
	@pagesize int = 0, 
	@row_count int = 0 OUTPUT, 
	@page_count int = 0 OUTPUT, 
	@sort varchar(200) = 'status_code',
	@order varchar(10) = 'asc',

    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
	DECLARE @columns varchar(100) = 'status_code, status_description as main_status'  
	DECLARE @where2 varchar(200) = ''
	
	IF @lookup > 0  
	BEGIN
		SET @pagesize = 1000000
		SET @columns = 'status_code, status_description as main_status'  
	END

	IF LEN(@service_type) > 0
	   SET @where2 = '[service_type] = ' + QUOTENAME(@service_type, '''')
    
	EXEC RunSimpleDynamicQuery
        @source = 'service_status',
		@columns = @columns,

        @filter = @filter,
        @where = '[status_code] like @filter or [status_description] like @filter',
		@where2 = @where2,
        
        @page = @page, 
    	@pagesize = @pagesize, 
    	@row_count = @row_count OUTPUT, 
    	@page_count = @page_count OUTPUT,
        @sort = @sort,
        @order = @order

END







GO


