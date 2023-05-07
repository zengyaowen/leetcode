# [2673. Make Costs of Paths Equal in a Binary Tree](https://leetcode.com/problems/make-costs-of-paths-equal-in-a-binary-tree)

[中文文档](/solution/2600-2699/2673.Make%20Costs%20of%20Paths%20Equal%20in%20a%20Binary%20Tree/README.md)

## Description

<p>You are given an integer <code>n</code> representing the number of nodes in a <strong>perfect binary tree</strong> consisting of nodes numbered from <code>1</code> to <code>n</code>. The root of the tree is node <code>1</code> and each node <code>i</code> in the tree has two children where the left child is the node <code>2 * i</code> and the right child is <code>2 * i + 1</code>.</p>

<p>Each node in the tree also has a <strong>cost</strong> represented by a given <strong>0-indexed</strong> integer array <code>cost</code> of size <code>n</code> where <code>cost[i]</code> is the cost of node <code>i + 1</code>. You are allowed to <strong>increment</strong> the cost of <strong>any</strong> node by <code>1</code> <strong>any</strong> number of times.</p>

<p>Return <em>the <strong>minimum</strong> number of increments you need to make the cost of paths from the root to each <strong>leaf</strong> node equal</em>.</p>

<p><strong>Note</strong>:</p>

<ul>
	<li>A <strong>perfect binary tree </strong>is a tree where each node, except the leaf nodes, has exactly 2 children.</li>
	<li>The <strong>cost of a path</strong> is the sum of costs of nodes in the path.</li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/2600-2699/2673.Make%20Costs%20of%20Paths%20Equal%20in%20a%20Binary%20Tree/images/binaryytreeedrawio-4.png" />
<pre>
<strong>Input:</strong> n = 7, cost = [1,5,2,2,3,3,1]
<strong>Output:</strong> 6
<strong>Explanation:</strong> We can do the following increments:
- Increase the cost of node 4 one time.
- Increase the cost of node 3 three times.
- Increase the cost of node 7 two times.
Each path from the root to a leaf will have a total cost of 9.
The total increments we did is 1 + 3 + 2 = 6.
It can be shown that this is the minimum answer we can achieve.
</pre>

<p><strong class="example">Example 2:</strong></p>
<img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/2600-2699/2673.Make%20Costs%20of%20Paths%20Equal%20in%20a%20Binary%20Tree/images/binaryytreee2drawio.png" style="width: 205px; height: 151px;" />
<pre>
<strong>Input:</strong> n = 3, cost = [5,3,3]
<strong>Output:</strong> 0
<strong>Explanation:</strong> The two paths already have equal total costs, so no increments are needed.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>3 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>n + 1</code> is a power of <code>2</code></li>
	<li><code>cost.length == n</code></li>
	<li><code>1 &lt;= cost[i] &lt;= 10<sup>4</sup></code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def minIncrements(self, n: int, cost: List[int]) -> int:
        def dfs(i: int) -> int:
            if (i << 1) > n:
                return cost[i - 1]
            l, r = dfs(i << 1), dfs(i << 1 | 1)
            nonlocal ans
            ans += max(l, r) - min(l, r)
            return cost[i - 1] + max(l, r)

        ans = 0
        dfs(1)
        return ans
```

### **Java**

```java
class Solution {
    private int[] cost;
    private int n;
    private int ans;

    public int minIncrements(int n, int[] cost) {
        this.n = n;
        this.cost = cost;
        dfs(1);
        return ans;
    }

    private int dfs(int i) {
        if ((i << 1) > n) {
            return cost[i - 1];
        }
        int l = dfs(i << 1);
        int r = dfs(i << 1 | 1);
        ans += Math.max(l, r) - Math.min(l, r);
        return cost[i - 1] + Math.max(l, r);
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int minIncrements(int n, vector<int>& cost) {
        int ans = 0;
        function<int(int)> dfs = [&](int i) -> int {
            if ((i << 1) > n) {
                return cost[i - 1];
            }
            int l = dfs(i << 1);
            int r = dfs(i << 1 | 1);
            ans += max(l, r) - min(l, r);
            return cost[i - 1] + max(l, r);
        };
        dfs(1);
        return ans;
    }
};
```

### **Go**

```go
func minIncrements(n int, cost []int) (ans int) {
	var dfs func(int) int
	dfs = func(i int) int {
		if (i << 1) > n {
			return cost[i-1]
		}
		l, r := dfs(i<<1), dfs(i<<1|1)
		ans += max(l, r) - min(l, r)
		return cost[i-1] + max(l, r)
	}
	dfs(1)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### **TypeScript**

```ts
function minIncrements(n: number, cost: number[]): number {
    let ans = 0;
    const dfs = (i: number): number => {
        if (i << 1 > n) {
            return cost[i - 1];
        }
        const [a, b] = [dfs(i << 1), dfs((i << 1) | 1)];
        ans += Math.max(a, b) - Math.min(a, b);
        return cost[i - 1] + Math.max(a, b);
    };
    dfs(1);
    return ans;
}
```

### **...**

```

```

<!-- tabs:end -->
