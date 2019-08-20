--
-- @lc app=leetcode.cn id=180 lang=mysql
--
-- [180] 连续出现的数字
--
# Write your MySQL query statement below

select distinct l1.Num as ConsecutiveNums
from Logs l1
join Logs l2 on l1.Id = l2.Id-1
join Logs l3 on l1.Id = l3.Id-2
where l1.Num = l2.Num and l1.Num = l3.Num;

-- select distinct l1.Num as ConsecutiveNums
-- from Logs l1, Logs l2, Logs l3
-- where l1.Id = l2.Id-1
-- and l1.Id = l3.Id-2
-- and l1.Num = l2.Num
-- and l1.Num = l3.Num;
