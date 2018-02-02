DROP PROCEDURE [dbo].[SetDefaultAddress]
GO

CREATE PROCEDURE [dbo].[SetDefaultAddress] 
@name_id int = 0,
@address_id int = 0,
@visit_id AS bigint = 0, /* if needed, the user will be got from here ... and subsequently, their rights*/
@action_status_id AS int = 0 OUTPUT,
@action_msg AS varchar(200) = '' OUTPUT
AS
BEGIN
    SET NOCOUNT ON

    UPDATE names SET 
		address_id = @address_id
    WHERE id = @name_id

    IF @@rowcount = 0
    BEGIN
        SET @action_status_id = -10
        SET @action_msg = 'Cannot find address to update.'
    END
END
GO