CREATE TABLE if not exists rand_numbers(number INT);

insert into rand_numbers ( number ) select rand() from (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t1, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t2, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t3, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t4, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t5, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t6, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t7, (
    select 0 as i
    union select 1 union select 2 union select 3
    union select 4 union select 5 union select 6
    union select 7 union select 8 union select 9
) as t8;
