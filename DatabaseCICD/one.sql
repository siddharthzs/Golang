-- liquekj a;lkdjf 
-- alskdfj
-- alsdkfja
create or replace FUNCTION dbo.two(a varchar, b varchar)
return table 
$function$
begin
    perform dbo.fn_helloworld(a,b)
    select * from dbo.fn_helloworld(a,b)
end
$function$