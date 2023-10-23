# [178. Rank Scores](https://leetcode.com/problems/rank-scores)

[中文文档](/solution/0100-0199/0178.Rank%20Scores/README.md)

## Description

<p>Table: <code>Scores</code></p>

<pre>
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| score       | decimal |
+-------------+---------+
id is the primary key (column with unique values) for this table.
Each row of this table contains the score of a game. Score is a floating point value with two decimal places.
</pre>

<p>&nbsp;</p>

<p>Write a solution to find the rank of the scores. The ranking should be calculated according to the following rules:</p>

<ul>
	<li>The scores should be ranked from the highest to the lowest.</li>
	<li>If there is a tie between two scores, both should have the same ranking.</li>
	<li>After a tie, the next ranking number should be the next consecutive integer value. In other words, there should be no holes between ranks.</li>
</ul>

<p>Return the result table ordered by <code>score</code> in descending order.</p>

<p>The result format is in the following example.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> 
Scores table:
+----+-------+
| id | score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
<strong>Output:</strong> 
+-------+------+
| score | rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
</pre>

## Solutions

<!-- tabs:start -->

**Solution 1: Use `DENSE_RANK()`**

Use `DENSE_RANK()` to solve this problem.

```sql
DENSE_RANK() OVER (
    PARTITION BY <expression>[{,<expression>...}]
    ORDER BY <expression> [ASC|DESC], [{,<expression>...}]
)
```

**Solution 2: Use variables**

MySQL only provides [window function](https://dev.mysql.com/doc/refman/8.0/en/window-function-descriptions.html) after version 8. In previous versions, variables can be used to achieve similar functions.

<!-- tabs:start -->

### **SQL**

```sql
# Write your MySQL query statement below
SELECT
    score,
    DENSE_RANK() OVER (ORDER BY score DESC) AS 'rank'
FROM Scores;
```

```sql
SELECT
    Score,
    CONVERT(rk, SIGNED) `Rank`
FROM
    (
        SELECT
            Score,
            IF(@latest = Score, @rank, @rank := @rank + 1) rk,
            @latest := Score
        FROM
            Scores,
            (
                SELECT
                    @rank := 0,
                    @latest := NULL
            ) tmp
        ORDER BY
            Score DESC
    ) s;
```

### **Pandas**

```python
import pandas as pd


def order_scores(scores: pd.DataFrame) -> pd.DataFrame:
    # Use the rank method to assign ranks to the scores in descending order with no gaps
    scores["rank"] = scores["score"].rank(method="dense", ascending=False)

    # Drop id column & Sort the DataFrame by score in descending order
    result_df = scores.drop("id", axis=1).sort_values(by="score", ascending=False)

    return result_df
```

<!-- tabs:end -->
