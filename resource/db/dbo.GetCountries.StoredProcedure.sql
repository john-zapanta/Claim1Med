SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetCountries]
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
	@code char(3) = NULL,
	@lookup int = 0,
	@filter as varchar(100) = '',

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

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
	DECLARE @columns varchar(100) = '*'  
	IF @lookup > 0  
	BEGIN
		SET @pagesize = 1000000
		SET @columns = 'code,country'
	END
    
	EXEC RunSimpleDynamicQuery
        @source = 'countries',
		@columns = @columns,

        @filter = @filter,
        @where = '[code] like @filter or [country] like @filter',
        
        @page = @page, 
    	@pagesize = @pagesize, 
    	@row_count = @row_count OUTPUT, 
    	@page_count = @page_count OUTPUT,
        @sort = @sort,
        @order = @order

END


GO
