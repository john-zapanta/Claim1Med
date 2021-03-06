DROP PROCEDURE [dbo].[GetCustomerServiceData]
GO
/****** Object:  StoredProcedure [dbo].[SP_CUSTOMER_SERVICE]    Script Date: 8/19/2017 9:18:23 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[GetCustomerServiceData]
	@visit_id as bigint = 0
as
begin
	set nocount on

	create table #CSTEMP (
		customer_service_no varchar(20),
		customer_service_type char(4),

        caller_first_name varchar(30),
        caller_middle_name varchar(30),
        caller_last_name varchar(30),
        caller_full_name varchar(100),
        caller_title varchar(20),
        caller_relationship varchar(20),

        caller_phone_no varchar(20),
        caller_mobile_no varchar(20),
        caller_fax_no varchar(20),
        caller_email varchar(60),

        caller_country_code char(3),
        caller_town varchar(30),
        caller_place varchar(30),

		member_id int,
        member_name_id int,
        first_name varchar(30),
        middle_name varchar(30),
        last_name varchar(30),
        full_name varchar(100),
        sex char(1),
        dob datetime,
        home_country_code char(3),
        nationality_code char(10),
        reference_no1 varchar(50),
        reference_no2 varchar(50),
        reference_no3 varchar(50),

        has_plan tinyint,
        certificate_no varchar(25),
        plan_code char(15),
		plan_name varchar(60),
        effective_date datetime,
        expiry_date datetime,
			
			MEM_CERT_ID int,
			MEM_RELATION char(2),
			MEM_DEP_CODE smallint,

        client_id int,
        client_name varchar(100),

        has_policy tinyint,
        policy_id int,
        policy_name_id int,
        policy_no varchar(20),
		policy_holder varchar(100),
        product_code char(10),
		product_name varchar(100),
        policy_issue_date datetime,
        policy_effective_date datetime,
        policy_expiry_date datetime,

		claim_id int,
		claim_no varchar(20),
		hcm_reference_no varchar(15),
		claim_type varchar(4),
		case_owner varchar(10),
		status_code char(1),
		claim_create_date datetime,
		claim_update_date datetime,

		service_id int,
		service_no varchar(20),
		service_type char(3),
		service_sub_type char(4),
		service_create_date datetime,
		service_update_date datetime,
		service_status_code char(1),
		service_sub_status_code char(3),

		provider_id int,
		doctor_id int,
			PROVIDER_NAME varchar(100),
			DOCTOR_NAME	varchar(100),

		invoice_no varchar(15),
		invoice_received_date datetime,
		service_date datetime,

		REC_CLM_SUB_TYPE char(4),
		REC_PROVIDER_NAME varchar(100),
		REC_STATUS_CODE char(3),

		CAS_CLM_SUB_TYPE char(4),
		CAS_STATUS_CODE char(3),

		NOTES text,


		InsertUser varchar(10),
		InsertDate datetime,
		UpdateUser varchar(10),
		UpdateDate datetime
	)

	insert into #CSTEMP(claim_id, service_id, caller_relationship, has_policy, has_plan) Values(NULL, 0, 'MEMBER', 1, 1)

	select * from #CSTEMP

end
GO
