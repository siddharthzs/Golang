-- liquekj a;lkdjf 
-- alskdfj
-- alsdkfja
create or replace VIEW dbo.fn_o365hello(a varchar, b varchar)
return table 
$function$
begin
    select * from dbo.fn_helloworld(a,b);
    return query select now();
end
$function$