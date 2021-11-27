-- liquekj a;lkdjf 
-- alskdfj
-- alsdkfja
create or replace FUNCTION dbo.two(a varchar, b varchar)
return table 
$function$
begin
    select * from dbo.fn_helloworld(a,b);
    return query select now();
end
$function$