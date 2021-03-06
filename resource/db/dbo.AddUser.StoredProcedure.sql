SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

/* =============================================*/
/* Author:		JZ*/
/* Create date: 2014-05-24*/
/* Description:	*/
/* =============================================*/
CREATE PROCEDURE [dbo].[AddUser]  
(
    @id as int = 0 OUTPUT,
    @user_name as varchar(10) = '' ,
    @name as varchar(60) = '' ,
	@designation as varchar(60) = '' ,
	@phone_no as varchar(20) = '' ,
    @is_active as bit = 0,
    @is_supervisor as bit = 0,
    
    @action as tinyint = 10, /* 20 = insert; 10 = update; 0 = delete;*/
    @visit_id as bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
    @action_status_id as int = 0 OUTPUT,
    @action_msg as varchar(200) = '' OUTPUT
) AS
BEGIN
	SET NOCOUNT ON;

    DECLARE @user_id AS int = 0--dbo.F_VisitUserID(@visit_id)
    DECLARE @update_user AS varchar(10) = 'john'--LEFT(dbo.F_VisitUserName(@visit_id), 10)

    SET @action_msg = ''
    SET @action_status_id = 0

    IF NOT @action IN (0, 10, 20)
        SET @action_status_id = -250 /* invalid request*/

    -- IF @action_status_id >= 0 AND (@user_id < 1 or not exists (select * from IHMS01.dbo.users where [id] = @user_id) ) and exists (select * from IHMS01.dbo.users where [id] > 0) /* skip during DB initialisation */
        -- SET @action_status_id = -15 /* not logged in*/
  	
    IF @action_status_id < 0 GOTO DONE
    
    DECLARE @upd_mode [tinyint] 
    -- @upd_mode = 1 New
    -- @upd_mode = 0 Update 
    -- @upd_mode = 2 Delete

    IF @id = 0 
        SET @upd_mode = 1 -- New record
    ELSE IF @action = 0
        SET @upd_mode = 2
    ELSE
        SET @upd_mode = 0

    DECLARE @update_date [datetime] = GETDATE()    

	IF @action = 20 /* insert */
	BEGIN
		SET NOCOUNT ON;
    END ELSE IF @action = 10 /* update */
    BEGIN
		UPDATE users set
			name = @name,
			user_name = @user_name,
			designation = @designation,
			phone_no = @phone_no,
			is_active = @is_active,
			is_supervisor = @is_supervisor,
			update_user = @update_user,
			update_date = getdate()
        WHERE id = @id
	END ELSE IF @action = 0 /* update */
	BEGIN
		DELETE users WHERE id = @id
    END
    
DONE:
    IF @action_status_id < 0
    BEGIN
        DECLARE @error_log_id AS int
        --EXEC dbo.LogError @error_log_id OUTPUT, @action_status_id, '', 0, @visit_id 
        --SET @action_msg = dbo.F_GetErrorLog(@error_log_id)
    END
    
END


GO
