SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetClaimDiagnosis] 
-- ***************************************************************************************************
-- Last modified on
-- 14-OCT-2014 ihms.0.0.1.0
-- *************************************************************************************************** 
(
    @id int = 0, 
	@sort varchar(200) = '',
	@order varchar(10) = '',

    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
    DECLARE @where AS varchar(100)
	DECLARE @sql nvarchar(MAX)
	DECLARE @params nvarchar(max)
	DECLARE @param_id int

	IF LEN(@sort) = 0 SET @sort = 'diagnosis_group,diagnosis_code'
	IF LEN(@order) = 0 SET @order = 'asc'

	CREATE TABLE #diagnosis (
		temp_id int iDENTITY(1,1),
		service_id int,
		parent_id int,
		id int,
		diagnosis_group varchar(10),
		diagnosis_code varchar(10),
		sort varchar(10)
	)

	CREATE TABLE #diagnosisgroups (
		temp_id int iDENTITY(1,1),
		parent_id int DEFAULT 0,
		service_id int,
		id int,
		diagnosis_group varchar(10) DEFAULT '',
		diagnosis_code varchar(10),
		sort varchar(10)
	)

	IF @id > 0
	BEGIN
		SET @param_id = @id
		SET @params = '@id int'
		SET @where = 'service_id = @id'
	END

	SET @sql = '
		INSERT INTO #diagnosis(
			service_id,
			id,
			diagnosis_group,
			diagnosis_code
		)
		SELECT
			@id,
			id,
			diagnosis_group,
			diagnosis_code
		FROM claim_diagnosis
		WHERE
	' + @where --+ ' ORDER BY ' + @sort +' '+ @order

	EXEC sp_executesql @sql, @params, @param_id

	INSERT INTO #diagnosisgroups(
		service_id,
		diagnosis_group,
		diagnosis_code
	) SELECT
		@id,
		diagnosis_group,
		diagnosis_group
	FROM #diagnosis GROUP BY diagnosis_group

	DECLARE @last_id int
	SELECT @last_id = MAX(id) FROM #diagnosis

	UPDATE #diagnosisgroups SET id = (@last_id + temp_id), sort = replace(STR(temp_id, 2), ' ', '0')
	UPDATE #diagnosis SET parent_id = g.id, sort = g.sort+'.'+replace(STR(d.temp_id, 2), ' ', '0') FROM #diagnosis d JOIN #diagnosisgroups g ON d.diagnosis_group = g.diagnosis_code

/*
	SELECT 
		parent_id,
		id,
		diagnosis_group,
		diagnosis_code
	FROM #diagnosisgroups
*/

	INSERT INTO #diagnosis(
		service_id,
		id,
		parent_id,
		diagnosis_group,
		diagnosis_code,
		sort
	) SELECT
		service_id,
		id,
		parent_id,
		diagnosis_group,
		diagnosis_code,
		sort
	FROM #diagnosisgroups

	SELECT 
		d.service_id,
		--s.service_no,
		--s.service_type,
		d.id,
		d.parent_id,
		d.diagnosis_group,
		d.diagnosis_code,
		icd.diagnosis,
		c.condition,
		CAST(CASE WHEN d.diagnosis_group = s.diagnosis_code THEN 1 ELSE 0 END as bit) as is_default
	FROM #diagnosis d
	JOIN services s on d.service_id = s.id
	JOIN icd on d.diagnosis_code = icd.code
	LEFT OUTER JOIN claim_diagnosis c on d.id = c.id
	ORDER BY d.sort --diagnosis_group, diagnosis_code

	DROP TABLE #diagnosis
END


GO
