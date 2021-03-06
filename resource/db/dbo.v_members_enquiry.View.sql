SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[v_members_enquiry]
AS
	SELECT        
		IP_ID AS id,		COMP_ID AS name_id,		CERT_ID AS certificate_id,		CERT_NO AS certificate_no,		DEP_CODE AS dependent_code,		CLT_CERT_NO AS alpha_id,		FIRST_NAME AS first_name,		MIDDLE_NAME AS middle_name,		LAST_NAME AS last_name,		MEMBER_NAME AS full_name,		DOB AS dob,		SEX AS sex,		NAT_CODE AS nationality_code,		HOME_CTRY AS home_country_code,		RELATION AS relationship_code,		REF_NO1 AS reference_no1,		REF_NO2 AS reference_no2,		REF_NO3 AS reference_no3,		STATUS AS status_code,		ISSUE_DATE AS issue_date,		EFF_DATE AS start_date,		EXP_DATE AS end_date,		CANCEL_DATE AS cancelation_date,		REINSTATE_DATE AS reinstatement_date,		HIST_ID AS history_id,		HIST_TYPE AS history_type,		POLICY_ID AS policy_id,		POLICY_NO AS policy_no,		POL_ISSUE_DATE AS policy_issue_date,		POL_EFF_DATE AS policy_start_date,		POL_EXP_DATE AS policy_end_date,		PLAN_CODE AS plan_code,		PlanName AS plan_name,		PROD_CODE AS product_code,		PROD_NAME AS product_name,		POLICY_HOLDER AS policy_holder,		CLIENT_ID AS client_id,		CLIENT_NAME AS client_name,		IS_POLICY AS has_policy,		IS_PLAN AS has_plan
	FROM dbo.vw_member_search


GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[40] 4[40] 2[13] 3) )"
      End
      Begin PaneConfiguration = 1
         NumPanes = 3
         Configuration = "(H (1 [50] 4 [25] 3))"
      End
      Begin PaneConfiguration = 2
         NumPanes = 3
         Configuration = "(H (1 [50] 2 [25] 3))"
      End
      Begin PaneConfiguration = 3
         NumPanes = 3
         Configuration = "(H (4 [30] 2 [40] 3))"
      End
      Begin PaneConfiguration = 4
         NumPanes = 2
         Configuration = "(H (1 [56] 3))"
      End
      Begin PaneConfiguration = 5
         NumPanes = 2
         Configuration = "(H (2 [66] 3))"
      End
      Begin PaneConfiguration = 6
         NumPanes = 2
         Configuration = "(H (4 [50] 3))"
      End
      Begin PaneConfiguration = 7
         NumPanes = 1
         Configuration = "(V (3))"
      End
      Begin PaneConfiguration = 8
         NumPanes = 3
         Configuration = "(H (1[56] 4[18] 2) )"
      End
      Begin PaneConfiguration = 9
         NumPanes = 2
         Configuration = "(H (1 [75] 4))"
      End
      Begin PaneConfiguration = 10
         NumPanes = 2
         Configuration = "(H (1[66] 2) )"
      End
      Begin PaneConfiguration = 11
         NumPanes = 2
         Configuration = "(H (4 [60] 2))"
      End
      Begin PaneConfiguration = 12
         NumPanes = 1
         Configuration = "(H (1) )"
      End
      Begin PaneConfiguration = 13
         NumPanes = 1
         Configuration = "(V (4))"
      End
      Begin PaneConfiguration = 14
         NumPanes = 1
         Configuration = "(V (2))"
      End
      ActivePaneConfig = 0
   End
   Begin DiagramPane = 
      Begin Origin = 
         Top = 0
         Left = 0
      End
      Begin Tables = 
         Begin Table = "vw_member_search"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 400
               Right = 242
            End
            DisplayFlags = 280
            TopColumn = 0
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 2025
         Alias = 2400
         Table = 1170
         Output = 1560
         Append = 1400
         NewValue = 1170
         SortType = 1350
         SortOrder = 1410
         GroupBy = 1350
         Filter = 1350
         Or = 1350
         Or = 1350
         Or = 1350
      End
   End
End
' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'VIEW',@level1name=N'v_members_enquiry'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=1 , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'VIEW',@level1name=N'v_members_enquiry'
GO
