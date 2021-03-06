SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE function [dbo].[f_document_categories] (@document_id int)
returns varchar(100) as begin

	declare @categories varchar(100)

    select
      @categories = COALESCE(@categories + ',', '') + rtrim(code)
    from document_categories
    where document_id = @document_id
    order by code

	return @categories
      
end

GO
