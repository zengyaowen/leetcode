# [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals)

[中文文档](/solution/0400-0499/0435.Non-overlapping%20Intervals/README.md)

## Description

<p>Given an array of intervals <code>intervals</code> where <code>intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]</code>, return <em>the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> intervals = [[1,2],[2,3],[3,4],[1,3]]
<strong>Output:</strong> 1
<strong>Explanation:</strong> [1,3] can be removed and the rest of the intervals are non-overlapping.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> intervals = [[1,2],[1,2],[1,2]]
<strong>Output:</strong> 2
<strong>Explanation:</strong> You need to remove two [1,2] to make the rest of the intervals non-overlapping.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> intervals = [[1,2],[2,3]]
<strong>Output:</strong> 0
<strong>Explanation:</strong> You don&#39;t need to remove any of the intervals since they&#39;re already non-overlapping.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= intervals.length &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>intervals[i].length == 2</code></li>
	<li><code>-2 * 10<sup>4</sup> &lt;= start<sub>i</sub> &lt; end<sub>i</sub> &lt;= 2 * 10<sup>4</sup></code></li>
</ul>


## Solutions

Greedy.

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        if not intervals:
            return 0
        intervals.sort(key=lambda x: x[1])
        cnt, end = 0, intervals[0][1]
        for interval in intervals[1:]:
            if interval[0] >= end:
                end = interval[1]
            else:
                cnt += 1
        return cnt
```

### **Java**

```java
class Solution {
    public int eraseOverlapIntervals(int[][] intervals) {
        if (intervals == null || intervals.length == 0) {
            return 0;
        }
        Arrays.sort(intervals, Comparator.comparingInt(a -> a[1]));
        int end = intervals[0][1], cnt = 0;
        for (int i = 1; i < intervals.length; ++i) {
            if (intervals[i][0] >= end) {
                end = intervals[i][1];
            } else {
                ++cnt;
            }
        }
        return cnt;
    }
}
```

### **Go**

```go
func eraseOverlapIntervals(intervals [][]int) int {
	if intervals == nil || len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end, cnt := intervals[0][1], 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
		} else {
			cnt++
		}
	}
	return cnt
}
```

### **...**

```

```

<!-- tabs:end -->
