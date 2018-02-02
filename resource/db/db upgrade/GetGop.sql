DROP PROCEDURE [dbo].[GetGop]
GO

CREATE PROCEDURE [dbo].[GetGop]
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
	@id int = 0,
	@sub_type varchar(4) = '',
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	SELECT
		s.*,
		g.id,
		g.gop_name,
		g.start_date,
		g.end_date,
		g.provider_id,
		g.provider_name,
		g.hospital_medical_record,
		g.provider_contact_person,
		g.provider_fax_no,
		g.discount_type_id,
		g.discount,
		g.doctor_id,
		g.doctor_name,
		g.diagnosis_notes,
		g.notes,
		g.misc_expense,
		g.room_expense,
		g.length_of_stay,
		g.admission_first_call,
		g.link_invoice_id
	FROM v_service s
	JOIN v_gop g ON s.id = g.id
	WHERE s.id = @id
END
GO
