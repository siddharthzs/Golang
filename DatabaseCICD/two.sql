-- liquekj a;lkdjf 
-- alskdfj
-- alsdkfja
CREATE OR REPLACE PROCEDURE dbo.fn_o365two(a varchar, b varchar)
return table 
$function$
begin
    select * from dbo.fn_helloworld(a,b);
    return query select now();
end
$function$