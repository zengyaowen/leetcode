# [2669. 统计 Spotify 排行榜上艺术家出现次数](https://leetcode.cn/problems/count-artist-occurrences-on-spotify-ranking-list)

[English Version](/solution/2600-2699/2669.Count%20Artist%20Occurrences%20On%20Spotify%20Ranking%20List/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>表：&nbsp;<code><font face="monospace">Spotify</font></code></p>

<pre>
+-------------+---------+ 
| 列名        | 类型    | 
+-------------+---------+ 
| id          | int     | 
| track_name  | varchar |
| artist      | varchar |
+-------------+---------+
id是该表的主键。
每行包含 id、track_name 和 artist。
</pre>

<p>编写一个SQL查询来查找每个艺术家在Spotify排行榜上出现的次数。</p>

<p>返回结果表，其中包含艺术家的名称以及相应的出现次数，按出现次数<strong>降序</strong>排列。如果出现次数相等，则按艺术家名称<strong>升序</strong>排列。</p>

<p>查询结果格式如下所示：</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：
</strong>Spotify 表: 
+---------+--------------------+------------+ 
| id      | track_name         | artist     |  
+---------+--------------------+------------+
| 303651  | Heart Won't Forget | Sia        |
| 1046089 | Shape of you       | Ed Sheeran |
| 33445   | I'm the one        | DJ Khalid  |
| 811266  | Young Dumb &amp; Broke | DJ Khalid  | 
| 505727  | Happier            | Ed Sheeran |
+---------+--------------------+------------+ 
<strong>输出：
</strong>+------------+-------------+
| artist     | occurrences | 
+------------+-------------+
| DJ Khalid  | 2           |
| Ed Sheeran | 2           |
| Sia        | 1           | 
+------------+-------------+ 

<strong>解释：</strong>"occurrences" 列下按降序列出了出现次数的计数。如果出现次数相同，则艺术家名称按升序排序。
</pre>

## 解法

<!-- 这里可写通用的实现逻辑 -->

<!-- tabs:start -->

### **SQL**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```sql
# Write your MySQL query statement below
SELECT
	artist,
	count( 1 ) AS occurrences
FROM
	Spotify
GROUP BY
	artist
ORDER BY
	occurrences DESC,
	artist;
```

<!-- tabs:end -->
