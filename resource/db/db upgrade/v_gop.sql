DROP VIEW [dbo].[v_gop]
GO

CREATE VIEW [dbo].[v_gop]
AS
	SELECT
		g.id,
		g.gop_name,
		s.start_date,
		s.end_date,
		s.provider_id,
		provider.name as provider_name,
		g.hospital_medical_record,
		s.provider_contact_person,
		s.provider_fax_no,
		d.discount_type_id,
		CASE WHEN d.discount_type_id in (1,3,2) then s.discount_percent else s.discount_amount  end as discount,
		s.doctor_id,
		doctor.name as doctor_name,
		s.diagnosis_notes,
		g.notes,
		g.misc_expense,
		g.room_expense,
		s.length_of_stay,
		g.admission_first_call,
		ISNULL(g.link_invoice_id, 0) AS link_invoice_id
	FROM gop_medical_services g
	JOIN services s on g.id = s.id
	LEFT OUTER JOIN names provider ON s.provider_id = provider.id
	LEFT OUTER JOIN names doctor ON s.doctor_id = doctor.id
	LEFT OUTER JOIN provider_discount d on s.provider_id = d.name_id
GO


