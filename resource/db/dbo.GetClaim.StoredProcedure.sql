SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetClaim] 
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
    @id int = 0,
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;
/*
	SELECT
		*
	FROM claims where id = @id
*/

	select  
		c.[id], 
		c.[claim_no],
		c.[member_id], 
		m.[main_name_id],
		c.[name_id], 
		c.[policy_id], 
		c.[client_id], 
		c.[product_code], 
		c.[plan_code], 
		c.[claim_type], 
		c.[base_currency_code], 
		c.[client_currency_code], 
		c.[eligibility_currency_code], 
		c.[notification_date], 
		c.[case_owner], 
		c.[status_code], 
		S.[status],
		c.[hcm_reference], 
		c.[reference_no1], 
		c.[reference_no2], 
		c.[reference_no3], 
        c.[sub_product], -- 10-SEP-2014: PANIN
		--c.IS_LARGE_LOSS,
		csv_service_id = cast(isnull(I.id, 0) as int),
		c.[is_prexisting],
		c.[is_accident],
		c.[first_symptom_date],
		c.[first_consultation_date],
		c.[incident_date],
		c.[country_of_incident],
		c.[accident_date],
		c.[accident_code],
		a.[accident],
		--c.[diagnosis_code],
		--c.[diagnosis_notes],
		--icd.[diagnosis],
		--c.[final_diagnosis_code]
		c.[create_date], 
		c.[create_user], 
		c.[update_date],  
		c.[update_user]
	from claims c
		left outer join members m ON c.member_id = m.id
		left outer join claim_status S on c.status_code = S.code
		left outer join services i on c.id = I.claim_id and I.claim_type = 'CSV'
		left outer join accident_types a on c.accident_code = a.code
	where c.id = @id

/*
	select  
		c.CLAIM_NO,	
		c.IS_PREEXISTING,
		c.IS_ACCIDENT,
		c.SYMP1_DATE,
		c.CONSULT1_DATE,
		c.INC_DATE,
		c.INC_CTRY,
		c.ACCD_DATE,
		c.ACCD_CODE,
		c.ICD_CODE,
		c.CONDITION,
		ICDDescription = cast(ICD.DESCRPTION as text),
		C.FINAL_ICD,
		FINAL_CONDITION = cast(C.FINAL_CONDITION as text),
		FINAL_Description = cast(icd2.descrption as text),
		AccidentType = A.ACCD_DESC
	from CLAIMMAIN c
		left outer join ICD9 ICD on C.ICD_CODE = ICD.ICD_CODE
		left outer join ICD9 icd2 on C.final_icd = icd2.icd_code
		left outer join ACCIDENTTYPES A on C.ACCD_CODE = A.ACCD_CODE
	where CLAIM_NO = @CLAIM_NO
*/
END

GO
