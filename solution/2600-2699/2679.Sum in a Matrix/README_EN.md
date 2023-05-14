# [2679. Sum in a Matrix](https://leetcode.com/problems/sum-in-a-matrix)

[中文文档](/solution/2600-2699/2679.Sum%20in%20a%20Matrix/README.md)

## Description

<p>You are given a <strong>0-indexed</strong> 2D integer array <code>nums</code>. Initially, your score is <code>0</code>. Perform the following operations until the matrix becomes empty:</p>

<ol>
	<li>From each row in the matrix, select the largest number and remove it. In the case of a tie, it does not matter which number is chosen.</li>
	<li>Identify the highest number amongst all those removed in step 1. Add that number to your <strong>score</strong>.</li>
</ol>

<p>Return <em>the final <strong>score</strong>.</em></p>
<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [[7,2,1],[6,4,2],[6,5,3],[3,2,1]]
<strong>Output:</strong> 15
<strong>Explanation:</strong> In the first operation, we remove 7, 6, 6, and 3. We then add 7 to our score. Next, we remove 2, 4, 5, and 2. We add 5 to our score. Lastly, we remove 1, 2, 3, and 1. We add 3 to our score. Thus, our final score is 7 + 5 + 3 = 15.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [[1]]
<strong>Output:</strong> 1
<strong>Explanation:</strong> We remove 1 and add it to the answer. We return 1.</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 300</code></li>
	<li><code>1 &lt;= nums[i].length &lt;= 500</code></li>
	<li><code>0 &lt;= nums[i][j] &lt;= 10<sup>3</sup></code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def matrixSum(self, nums: List[List[int]]) -> int:
        for row in nums:
            row.sort()
        n = len(nums[0])
        return sum(max(row[j] for row in nums) for j in range(n))
```

### **Java**

```java
class Solution {
    public int matrixSum(int[][] nums) {
        for (var row : nums) {
            Arrays.sort(row);
        }
        int ans = 0;
        for (int j = 0; j < nums[0].length; ++j) {
            int mx = 0;
            for (var row : nums) {
                mx = Math.max(mx, row[j]);
            }
            ans += mx;
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int matrixSum(vector<vector<int>>& nums) {
        for (auto& row : nums) {
            sort(row.begin(), row.end());
        }
        int ans = 0;
        for (int j = 0; j < nums[0].size(); ++j) {
            int mx = 0;
            for (auto& row : nums) {
                mx = max(mx, row[j]);
            }
            ans += mx;
        }
        return ans;
    }
};
```

### **Go**

```go
func matrixSum(nums [][]int) (ans int) {
	for _, row := range nums {
		sort.Ints(row)
	}
	for i := 0; i < len(nums[0]); i++ {
		mx := 0
		for _, row := range nums {
			mx = max(mx, row[i])
		}
		ans += mx
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### **...**

```

```

<!-- tabs:end -->
