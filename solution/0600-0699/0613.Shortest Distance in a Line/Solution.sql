# Write your MySQL query statement below
SELECT
    min(abs(p1.x - p2.x)) AS shortest
FROM
    Point AS p1
    JOIN Point AS p2 ON p1.x != p2.x;
