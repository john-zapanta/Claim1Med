DROP PROCEDURE [dbo].[AddClaimStatusHistory]
GO

CREATE PROCEDURE [dbo].[AddClaimStatusHistory]  
/****************************************************************************************************
 Update History:

 30-JAN-2018
****************************************************************************************************/
(
	@claim_id int = 0,
	@status_code char(1) = '',
	-- @new_status_code char(1) = '',
    @action as tinyint = 10, /* 20 = insert; 10 = update; 0 = delete;*/
    @visit_id as bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
    @action_status_id as int = 0 OUTPUT,
    @action_msg as varchar(200) = '' OUTPUT
) AS
BEGIN
	SET NOCOUNT ON;

	DECLARE @user_id AS int = 0
	DECLARE @update_user AS varchar(10) = ''
	DECLARE @update_date datetime = GETDATE()    

	EXEC [dbo].[System_ValidateUser]
		@user_id = @user_id OUTPUT,
		@user_name = @update_user OUTPUT,
		@action = @action,
		@visit_id = @visit_id,
		@action_status_id = @action_status_id OUTPUT,
		@action_msg = @action_msg OUTPUT

	--SET @action_msg = 'TEST: ' + @status_code
	--RETURN
	IF @action_status_id = 0
	BEGIN
		DECLARE @status_id int

		INSERT claim_status_history (
			claim_id,
			status_code,
			create_date,
			create_user,
			update_date,
			update_user
		)
		VALUES (
			@claim_id,
			@status_code,
			getdate(),
			@update_user,
			getdate(),
			@update_user
		)

		SET @status_id = CAST(SCOPE_IDENTITY() as int)

		UPDATE claims SET
			status_code = @status_code,
			status_id = @status_id
		WHERE id = @claim_id

		SET @action_status_id = 0

		IF @status_code = 'O' 
			SET @action_msg = 'Claim was successfully Re-Opened.'
		ELSE IF @status_code = 'C' 
			SET @action_msg = 'Claim status was successfully Closed.'
		ELSE IF @status_code = 'D' 
			SET @action_msg = 'Claim status was successfully Declined.'
			
	END
END
GO
