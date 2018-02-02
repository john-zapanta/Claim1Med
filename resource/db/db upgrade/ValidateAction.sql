DROP PROCEDURE [dbo].[ValidateAction]
GO

CREATE PROCEDURE [dbo].[ValidateAction]  
/****************************************************************************************************
 Update History:

 30-JAN-2018
****************************************************************************************************/
@id int = 0,
@string_value varchar(30) = 0,
@key_id int = 0,
@action_name varchar(100) = '' OUTPUT,
@visit_id as bigint = 0, 
@action_status_id as int = 0 OUTPUT, -- 1: Warning, 2: Info, 0:None
@action_msg as varchar(2048) = '' OUTPUT
AS
BEGIN
	SET NOCOUNT ON;

	DECLARE @user_id AS int = 0
	DECLARE @user_name AS varchar(10) = ''

	EXEC [dbo].[System_ValidateUser]
		@user_id = @user_id OUTPUT,
		@user_name = @user_name OUTPUT,
		@visit_id = @visit_id

	EXEC [dbo].[sp_validate_action] 
		@ACTION_ID = @id,
		@VALUE_STR = @string_value,
		@INVOICE_ID = @key_id,
		@USER_NAME = @user_name,
		@ACTION_NAME = @action_name OUTPUT,
		@ErrCode = @action_status_id OUTPUT,
		@ErrMessage = @action_msg OUTPUT

END
GO
