DROP PROCEDURE [dbo].[AddFlagTypes]
GO

CREATE PROCEDURE [dbo].[AddFlagTypes]  
(
	@code char(4) = '' OUTPUT,
    @service_description varchar(60) = '',
	@display_name varchar(100) = '',
	@is_active bit = 1,
	
    @action as tinyint = 10, /* 20 = insert; 10 = update; 0 = delete;*/
    @visit_id as bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
    @action_status_id as int = 0 OUTPUT,
    @action_msg as varchar(200) = '' OUTPUT
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

		SELECT @cnt = count(*) from service_types where code = @code and service_type = 'FLG'

		IF @cnt > 0
		BEGIN
		   SET @action_msg = 'Code <B>' + @code + '</B> already exist.'
		   SET @action_status_id = -1
		END

		SELECT @cnt = count(*) from service_types where UPPER(service_description) = UPPER(@service_description) and service_type = 'FLG'

		IF @cnt > 0
		BEGIN
		   IF @action_msg = ''
		      SET @action_msg = 'Flag Type Description<B>' + @service_description + '</B> already exist.'
		   ELSE
		      SET @action_msg = @action_msg + ' </BR> Flag Type Description <B>' + @service_description + '</B> already exist.' 

		   SET @action_status_id = -1
		END

		IF NOT @action_msg = '' GOTO DONE

		INSERT INTO service_types(code, service_type, service_description, display_name, is_active, create_user, create_date)
		VALUES(@code, 'FLG', @service_description, @display_name, @is_active, @update_user, @update_date)
    END ELSE IF @action = 10 /* update */
    BEGIN
		UPDATE service_types set
			service_description = @service_description,
			display_name = @display_name,
			is_active = @is_active,
			update_user = @update_user,
			update_date = @update_date
        WHERE code = @code and service_type = 'FLG'
	END ELSE IF @action = 0 /* delete */
	BEGIN
		DELETE service_types WHERE code = @code and service_type = 'FLG'
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


