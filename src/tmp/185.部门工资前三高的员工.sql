--
-- @lc app=leetcode.cn id=185 lang=mysql
--
-- [185] 部门工资前三高的员工
--
# Write your MySQL query statement below


SELECT d.Name AS Department, e.Name AS Employee, e.Salary
FROM Employee e
JOIN Department d
ON e.DepartmentId = d.Id
WHERE (
	SELECT COUNT(DISTINCT Salary) FROM Employee
	WHERE Salary > e.Salary
	AND DepartmentId = d.Id
) < 3
ORDER BY d.Name, e.Salary
DESC;
