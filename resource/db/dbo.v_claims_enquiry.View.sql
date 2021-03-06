SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO



CREATE VIEW [dbo].[v_claims_enquiry]
AS
	SELECT
		ISNULL(i.id, c.id) AS id,
		c.id as claim_id,
		i.id as service_id,
		c.claim_no,
		i.service_no,
		i.claim_type,
		i.service_type,
		i.service_sub_type,
		st.sub_type as service_name,
		i.service_date,
		c.hcm_reference,
		c.case_owner,
		c.member_id,
		c.client_id,
		c.product_code,
		c.plan_code,

		prov.full_name as provider_name,
		doc.full_name as doctor_name,

		isnull(i.invoice_no, '') as invoice_no,
		i.invoice_received_date,
		i.start_date,
		i.end_date,
		i.claim_currency_code,
		i.settlement_currency_code,
		isnull(i.actual_amount, 0) as actual_amount,
		isnull(i.paid_amount, 0) as paid_amount,

		isnull(i.diagnosis_code, c.diagnosis_code) as diagnosis_code,
		icd.diagnosis,

		CD.status as claim_status,
		i.status_code,
		i.sub_status_code,
		INS.status AS status,
		RES.sub_status,

		p.policy_no,
		PH.full_name AS policy_holder,

		o.full_name as client_name,
		r.product_name,
		pd.plan_name,

		pt.full_name as patient_name,
		pt.first_name,
		pt.last_name,
		pt.middle_name,
		pt.gender,
		datediff(year, pt.dob, getdate()) as age,

		m.certificate_no,
		m.reference_no1,
		m.reference_no2,
		m.reference_no3,

		-- 'Condition ' + ISNULL(CAST(C.CONDITION AS varchar(2048)), '') AS CONDITION,
		isnull(i.update_date, c.update_date) as update_date,
		isnull(i.create_date, c.create_date) as create_date
	FROM dbo.claims c
		LEFT OUTER JOIN	dbo.services i ON c.id = i.claim_id and i.is_deleted = 0
		--LEFT OUTER JOIN	dbo.services i ON c.id = i.claim_id and i.service_type = 'INV' and i.is_deleted = 0
		LEFT OUTER JOIN dbo.names pt ON c.name_id = pt.id
		LEFT OUTER JOIN dbo.names o ON c.client_id = o.id
		LEFT OUTER JOIN dbo.names prov ON i.provider_id = prov.id
		LEFT OUTER JOIN dbo.names doc ON i.doctor_id = doc.id
		LEFT OUTER JOIN dbo.members m ON C.member_id = M.id
		LEFT OUTER JOIN dbo.products r ON c.product_code = r.code
		LEFT OUTER JOIN	dbo.plans pd on c.product_code = pd.product_code and c.plan_code = pd.code
		LEFT OUTER JOIN	dbo.sub_status_codes res ON I.service_type = RES.service_type AND I.sub_status_code = RES.sub_status_code
		LEFT OUTER JOIN	dbo.service_sub_types st on I.service_type = st.service_type and I.service_sub_type = st.code
		LEFT OUTER JOIN	dbo.claim_status CD ON C.status_code = CD.code
		LEFT OUTER JOIN	dbo.invoice_status ins on i.service_type = ins.service_type and i.status_code = ins.code
		LEFT OUTER JOIN	dbo.policies p on c.policy_id = p.id
		LEFT OUTER JOIN	dbo.names ph on p.name_id = ph.id
		LEFT OUTER JOIN dbo.icd icd on isnull(i.diagnosis_code, c.diagnosis_code) = icd.code
	--WHERE i.service_type in ('INV', 'NOC', 'GOP', 'CAS', 'FLG', 'REC', 'COS') and c.is_deleted = 0 -- and I.IsDeleted = 0 -- not (C.CLM_TYPE = 'CSV')
	--WHERE i.service_type in ('INV', 'GOP') and c.is_deleted = 0 -- and I.IsDeleted = 0 -- not (C.CLM_TYPE = 'CSV')
	--WHERE i.service_type in ('INV') and c.is_deleted = 0 -- and I.IsDeleted = 0 -- not (C.CLM_TYPE = 'CSV')
	WHERE c.is_deleted = 0 and i.service_type in ('INV','GOP') --not C.CLM_TYPE = 'CSV'
	--WHERE c.is_deleted = 0




GO
