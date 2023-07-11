# [1050. 合作过至少三次的演员和导演](https://leetcode.cn/problems/actors-and-directors-who-cooperated-at-least-three-times)

[English Version](/solution/1000-1099/1050.Actors%20and%20Directors%20Who%20Cooperated%20At%20Least%20Three%20Times/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p><code>ActorDirector</code>&nbsp;表：</p>

<pre>
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| actor_id    | int     |
| director_id | int     |
| timestamp   | int     |
+-------------+---------+
timestamp 是这张表的主键.
</pre>

<p>&nbsp;</p>

<p>写一条SQL查询语句获取合作过至少三次的演员和导演的 id 对&nbsp;<code>(actor_id, director_id)</code></p>

<p><strong>示例：</strong></p>

<pre>
ActorDirector 表：
+-------------+-------------+-------------+
| actor_id    | director_id | timestamp   |
+-------------+-------------+-------------+
| 1           | 1           | 0           |
| 1           | 1           | 1           |
| 1           | 1           | 2           |
| 1           | 2           | 3           |
| 1           | 2           | 4           |
| 2           | 1           | 5           |
| 2           | 1           | 6           |
+-------------+-------------+-------------+

Result 表：
+-------------+-------------+
| actor_id    | director_id |
+-------------+-------------+
| 1           | 1           |
+-------------+-------------+
唯一的 id 对是 (1, 1)，他们恰好合作了 3 次。</pre>

## 解法

<!-- 这里可写通用的实现逻辑 -->

**方法一：使用 `GROUP BY` + `HAVING`**

我们将 `ActorDirector` 表按照 `actor_id` 和 `director_id` 进行分组，然后使用 `HAVING` 过滤出合作次数大于等于 $3$ 次的组。

<!-- tabs:start -->

### **SQL**

```sql
# Write your MySQL query statement below
SELECT actor_id, director_id
FROM ActorDirector
GROUP BY 1, 2
HAVING count(1) >= 3;
```

<!-- tabs:end -->
