SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[v_claim_notes]
as
	select 
		N.* ,
		reference_no = isnull(I.service_no, C.claim_no),
		service_type = isnull(I.service_type, 'CLM'),

		can_edit = cast(0 as bit),
		note_type_name = M.note_type,
		note_sub_type_name = S.note_sub_type
	from claim_notes N
		left outer join services I on N.service_id = I.id
		left outer join claims C on N.claim_id = C.id
		join note_types M ON N.note_type = M.code
		join note_sub_types S ON N.note_type = S.note_type AND N.note_sub_type = S.code


GO
