CREATE DEFINER=`pgadmin`@`%` FUNCTION `SPLIT_STR`(
  str VARCHAR(255)
) RETURNS varchar(255) CHARSET utf8
begin
/* =================================================
Author:lzx
Create date:2020-07-06
description:分隔字符串，并在每个值加上单引号
modify record:

TEST CASE
select split_str('B2B,B3B,B4B'); return 'B2B','B3B','B4B' 
select split_str('B2B,B3B');  return 'B2B','B3B'
SELECT split_str('B2B');    return 'B2B'
==================*/
declare i int default 1;
declare count int;
declare out_str varchar(4000) default '';
-- SET str='1,2,3,4';
IF str is null then
	return null;
end if;
SET count=LENGTH(str) - LENGTH(REPLACE( str,',','' ));

IF count=0 then
	return concat('''',str,'''');
end if;

IF count=1 then
   set out_str=concat('''',substring_index(str,',',1),''',');
   set out_str=concat(out_str,'''',substring_index(str,',',-1),'''');
   return out_str;
end if;

   while i<=count+1 DO
   if i<count+1 then
	set out_str=concat(out_str,'''',substring_index(substring_index(str,',',i),',','-1'),''',');

    elseif i=count+1 then	

    set out_str=concat(out_str,'''',substring_index(substring_index(str,',',i),',','-1'),'''');
   END IF;
	SET i=i+1;
   end while;
   

RETURN out_str;
end