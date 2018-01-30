DROP PROCEDURE [dbo].[AddMemberMedicalNotes]
GO

CREATE PROCEDURE [dbo].[AddMemberMedicalNotes]  
/****************************************************************************************************
 Update History:
 ===============

 02-OCT-2017
****************************************************************************************************/
(
	@id int = 0,
	@claim_id int = 0,
	@medical_history_notes varchar(max) = '',
    @action as tinyint = 10, /* 20 = insert; 10 = update; 0 = delete;*/
    @visit_id as bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
    @action_status_id as int = 0 OUTPUT,
    @action_msg as varchar(200) = '' OUTPUT
) AS
BEGIN
	SET NOCOUNT ON;

	IF @action = 10
	BEGIN
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

		DECLARE @old_notes varchar(max)

		SELECT 
			@old_notes = medical_history_notes
		FROM members
		WHERE id = @id

		INSERT INTO [dbo].[medical_history_notes] (
			id,
			medical_history_notes,
			update_date,
			update_user
		) VALUES (
			@id,
			@old_notes,
			getdate(),
			@update_user
		)

		UPDATE members SET
			medical_history_notes = @medical_history_notes
		WHERE id = @id

		INSERT INTO dbo.auditlogs (
			claim_id,
			service_id,
			code,
			module,
			notes,
			create_date,
			create_user
		) VALUES (
			@claim_id,
			0,
			'B03',
			'CLM',
			@medical_history_notes,
			getdate(),
			@update_user
		)
	END
END
GO
