# Write your MySQL query statement below
select e1.name as Employee
from Employee e join Employee e1 ON e.id = e1.managerId
where e1.salary > e.salary