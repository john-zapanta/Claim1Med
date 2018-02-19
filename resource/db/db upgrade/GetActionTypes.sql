DROP PROCEDURE [dbo].[GetActionTypes] 
GO

CREATE PROCEDURE [dbo].[GetActionTypes] 
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
		SELECT * FROM action_types WHERE code = @code
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
			@source = 'action_types',
			@columns = @columns,

			@filter = @filter,
			@where = '[code] like @filter or [action_type] like @filter',
        
			@page = @page, 
    		@pagesize = @pagesize, 
    		@row_count = @row_count OUTPUT, 
    		@page_count = @page_count OUTPUT,
			@sort = @sort,
			@order = @order
	END
END




GO

