# Write your MySQL query statement below

select t.name,t.bonus
from (select e.name, b.bonus
from Employee e
LEFT JOIN Bonus b ON e.empId = b.empId) as t
where t.bonus < 1000 OR t.bonus is null

