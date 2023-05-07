# [2672. Number of Adjacent Elements With the Same Color](https://leetcode.com/problems/number-of-adjacent-elements-with-the-same-color)

[中文文档](/solution/2600-2699/2672.Number%20of%20Adjacent%20Elements%20With%20the%20Same%20Color/README.md)

## Description

<p>There is a <strong>0-indexed</strong> array <code>nums</code> of length <code>n</code>. Initially, all elements are <strong>uncolored </strong>(has a value of <code>0</code>).</p>

<p>You are given a 2D integer array <code>queries</code> where <code>queries[i] = [index<sub>i</sub>, color<sub>i</sub>]</code>.</p>

<p>For each query, you color the index <code>index<sub>i</sub></code> with the color <code>color<sub>i</sub></code> in the array <code>nums</code>.</p>

<p>Return <em>an array </em><code>answer</code><em> of the same length as </em><code>queries</code><em> where </em><code>answer[i]</code><em> is the number of adjacent elements with the same color <strong>after</strong> the </em><code>i<sup>th</sup></code><em> query</em>.</p>

<p>More formally, <code>answer[i]</code> is the number of indices <code>j</code>, such that <code>0 &lt;= j &lt; n - 1</code> and <code>nums[j] == nums[j + 1]</code> and <code>nums[j] != 0</code> after the <code>i<sup>th</sup></code> query.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 4, queries = [[0,2],[1,2],[3,1],[1,1],[2,1]]
<strong>Output:</strong> [0,1,1,0,2]
<strong>Explanation:</strong> Initially array nums = [0,0,0,0], where 0 denotes uncolored elements of the array.
- After the 1<sup>st</sup> query nums = [2,0,0,0]. The count of adjacent elements with the same color is 0.
- After the 2<sup>nd</sup> query nums = [2,2,0,0]. The count of adjacent elements with the same color is 1.
- After the 3<sup>rd</sup>&nbsp;query nums = [2,2,0,1]. The count of adjacent elements with the same color is 1.
- After the 4<sup>th</sup>&nbsp;query nums = [2,1,0,1]. The count of adjacent elements with the same color is 0.
- After the 5<sup>th</sup>&nbsp;query nums = [2,1,1,1]. The count of adjacent elements with the same color is 2.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 1, queries = [[0,100000]]
<strong>Output:</strong> [0]
<strong>Explanation:</strong> Initially array nums = [0], where 0 denotes uncolored elements of the array.
- After the 1<sup>st</sup> query nums = [100000]. The count of adjacent elements with the same color is 0.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length&nbsp;== 2</code></li>
	<li><code>0 &lt;= index<sub>i</sub>&nbsp;&lt;= n - 1</code></li>
	<li><code>1 &lt;=&nbsp; color<sub>i</sub>&nbsp;&lt;= 10<sup>5</sup></code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def colorTheArray(self, n: int, queries: List[List[int]]) -> List[int]:
        nums = [0] * n
        ans = [0] * len(queries)
        x = 0
        for k, (i, c) in enumerate(queries):
            if i > 0 and nums[i] and nums[i - 1] == nums[i]:
                x -= 1
            if i < n - 1 and nums[i] and nums[i + 1] == nums[i]:
                x -= 1
            if i > 0 and nums[i - 1] == c:
                x += 1
            if i < n - 1 and nums[i + 1] == c:
                x += 1
            ans[k] = x
            nums[i] = c
        return ans
```

### **Java**

```java
class Solution {
    public int[] colorTheArray(int n, int[][] queries) {
        int m = queries.length;
        int[] nums = new int[n];
        int[] ans = new int[m];
        for (int k = 0, x = 0; k < m; ++k) {
            int i = queries[k][0], c = queries[k][1];
            if (i > 0 && nums[i] > 0 && nums[i - 1] == nums[i]) {
                --x;
            }
            if (i < n - 1 && nums[i] > 0 && nums[i + 1] == nums[i]) {
                --x;
            }
            if (i > 0 && nums[i - 1] == c) {
                ++x;
            }
            if (i < n - 1 && nums[i + 1] == c) {
                ++x;
            }
            ans[k] = x;
            nums[i] = c;
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<int> colorTheArray(int n, vector<vector<int>>& queries) {
        vector<int> nums(n);
        vector<int> ans;
        int x = 0;
        for (auto& q : queries) {
            int i = q[0], c = q[1];
            if (i > 0 && nums[i] > 0 && nums[i - 1] == nums[i]) {
                --x;
            }
            if (i < n - 1 && nums[i] > 0 && nums[i + 1] == nums[i]) {
                --x;
            }
            if (i > 0 && nums[i - 1] == c) {
                ++x;
            }
            if (i < n - 1 && nums[i + 1] == c) {
                ++x;
            }
            ans.push_back(x);
            nums[i] = c;
        }
        return ans;
    }
};
```

### **Go**

```go
func colorTheArray(n int, queries [][]int) (ans []int) {
	nums := make([]int, n)
	x := 0
	for _, q := range queries {
		i, c := q[0], q[1]
		if i > 0 && nums[i] > 0 && nums[i-1] == nums[i] {
			x--
		}
		if i < n-1 && nums[i] > 0 && nums[i+1] == nums[i] {
			x--
		}
		if i > 0 && nums[i-1] == c {
			x++
		}
		if i < n-1 && nums[i+1] == c {
			x++
		}
		ans = append(ans, x)
		nums[i] = c
	}
	return
}
```

### **TypeScript**

```ts
function colorTheArray(n: number, queries: number[][]): number[] {
    const nums: number[] = new Array(n).fill(0);
    const ans: number[] = [];
    let x = 0;
    for (const [i, c] of queries) {
        if (i > 0 && nums[i] > 0 && nums[i - 1] == nums[i]) {
            --x;
        }
        if (i < n - 1 && nums[i] > 0 && nums[i + 1] == nums[i]) {
            --x;
        }
        if (i > 0 && nums[i - 1] == c) {
            ++x;
        }
        if (i < n - 1 && nums[i + 1] == c) {
            ++x;
        }
        ans.push(x);
        nums[i] = c;
    }
    return ans;
}
```

### **...**

```

```

<!-- tabs:end -->
