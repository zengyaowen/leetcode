# [2886. Change Data Type](https://leetcode.com/problems/change-data-type)

[中文文档](/solution/2800-2899/2886.Change%20Data%20Type/README.md)

## Description

<pre>
DataFrame <code>students</code>
+-------------+--------+
| Column Name | Type   |
+-------------+--------+
| student_id  | int    |
| name        | object |
| age         | int    |
| grade       | float  |
+-------------+--------+
</pre>

<p>Write a solution to correct the errors:</p>

<p>The <code>grade</code> column is stored as floats,&nbsp;convert it to integers.</p>

<p>The result format is in the following example.</p>

<p>&nbsp;</p>
<pre>
<strong class="example">Example 1:</strong>
<strong>Input:
</strong>DataFrame students:
+------------+------+-----+-------+
| student_id | name | age | grade |
+------------+------+-----+-------+
| 1          | Ava  | 6   | 73.0  |
| 2          | Kate | 15  | 87.0  |
+------------+------+-----+-------+
<strong>Output:
</strong>+------------+------+-----+-------+
| student_id | name | age | grade |
+------------+------+-----+-------+
| 1          | Ava  | 6   | 73    |
| 2          | Kate | 15  | 87    |
+------------+------+-----+-------+
<strong>Explanation:</strong> 
The data types of the column grade is converted to int.</pre>

## Solutions

<!-- tabs:start -->

### **Pandas**

```python
import pandas as pd


def changeDatatype(students: pd.DataFrame) -> pd.DataFrame:
    students['grade'] = students['grade'].astype(int)
    return students
```

### **...**

```

```

<!-- tabs:end -->
