--
-- @lc app=leetcode.cn id=184 lang=mysql
--
-- [184] 部门工资最高的员工
--
# Write your MySQL query statement below

-- 只筛选出每个部门的最高工资
select d.Name as Department, e.Name as Employee, e.Salary as Salary
from Employee e
join Department d
on e.DepartmentId = d.Id
where e.Salary in (
	select max(Salary) from Employee e2
	where e.DepartmentId = e2.DepartmentId
);
