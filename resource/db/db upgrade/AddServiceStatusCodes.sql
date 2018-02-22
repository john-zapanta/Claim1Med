DROP PROCEDURE [dbo].[AddServiceStatusCodes]
GO

CREATE PROCEDURE [dbo].[AddServiceStatusCodes]  
(
	@status_code as char(1) = '',
	@sub_status_code as char(3) = '',
    @service_type as char(3) = '',
	@sub_status as varchar(100) = '',
	@is_system as bit = 1,
	@is_active as bit = 1,
    @visit_id as bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
    @action as tinyint = 10, /* 20 = insert; 10 = update; 0 = delete;*/
    @action_status_id as int = 0 OUTPUT,
    @action_msg as varchar(2048) = '' OUTPUT
) AS
BEGIN
	SET NOCOUNT ON;

	DECLARE @user_id AS int = 0
    DECLARE @update_user AS varchar(10) = ''
	DECLARE @update_date [datetime] = GETDATE()   

	EXEC [dbo].[System_ValidateUser]
		@user_id = @user_id OUTPUT,
		@user_name = @update_user OUTPUT,
		--@update_date datetime = NULL OUTPUT,
		@action = @action,
		@visit_id = @visit_id,
		@action_status_id = @action_status_id OUTPUT,
		@action_msg = @action_msg OUTPUT

	DECLARE @cnt AS int = 0

	IF @action = 20 /* insert */
	BEGIN
		SET NOCOUNT ON;

		SELECT @cnt = count(*) from service_status_codes where sub_status_code = @sub_status_code and service_type = @service_type

		IF @cnt > 0
		BEGIN
		   SET @action_msg = 'Code <B>' + @sub_status_code + '</B> already exist.'
		   SET @action_status_id = -1
		END

		SELECT @cnt = count(*) from service_status_codes where UPPER(sub_status) = UPPER(@sub_status) and service_type = @service_type

		IF @cnt > 0
		BEGIN
		   IF @action_msg = ''
		      SET @action_msg = 'Sub-Status Description <B>' + @sub_status + '</B> already exist.'
		   ELSE
		      SET @action_msg = @action_msg + ' </BR> Sub-Status Description <B>' + @sub_status + '</B> already exist.' 

		   SET @action_status_id = -1
		END

		IF NOT @action_msg = '' GOTO DONE

		INSERT INTO service_status_codes(status_code, sub_status_code, service_type, sub_status, is_system, is_active, create_user, create_date)
		VALUES(@status_code, @sub_status_code, @service_type, @sub_status, @is_system, @is_active, @update_user, @update_date)
    END ELSE IF @action = 10 /* update */
    BEGIN
		UPDATE service_status_codes set
			sub_status = @sub_status,
			is_system = @is_system,
			is_active = @is_active,
			update_user = @update_user,
			update_date = @update_date
        WHERE sub_status_code = @sub_status_code and service_type = @service_type
	END ELSE IF @action = 0 /* delete */
	BEGIN
		DELETE service_status_codes WHERE sub_status_code = @sub_status_code and service_type = @service_type
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


