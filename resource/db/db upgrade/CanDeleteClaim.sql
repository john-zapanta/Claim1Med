DROP PROCEDURE [dbo].[CanDeleteClaim]
GO

CREATE PROCEDURE [dbo].[CanDeleteClaim] 
@id int = NULL,
@change_plan int = 0,
@visit_id AS bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
@action_status_id AS int = 0 OUTPUT,
@action_msg AS varchar(200) = '' OUTPUT
AS
BEGIN
    SET NOCOUNT ON

    SET @action_status_id = 0
    SET @action_msg = ''

    IF EXISTS (SELECT
            *
        FROM services
        WHERE claim_id = @id
        AND is_deleted = 0)
	BEGIN
        IF @change_plan = 0
            SET @action_msg = 'Cannot delete claim. There are already services attached to it.'
        ELSE
            SET @action_msg = 'Cannot change plan, there are already services attached to this claim. Please create a new claim.'

		IF LEN(@action_msg) != 0
			SET @action_status_id = -1
	END
END
GO
