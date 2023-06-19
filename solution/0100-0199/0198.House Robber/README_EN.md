# [198. House Robber](https://leetcode.com/problems/house-robber)

[中文文档](/solution/0100-0199/0198.House%20Robber/README.md)

## Description

<p>You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and <b>it will automatically contact the police if two adjacent houses were broken into on the same night</b>.</p>

<p>Given an integer array <code>nums</code> representing the amount of money of each house, return <em>the maximum amount of money you can rob tonight <b>without alerting the police</b></em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,3,1]
<strong>Output:</strong> 4
<strong>Explanation:</strong> Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,7,9,3,1]
<strong>Output:</strong> 12
<strong>Explanation:</strong> Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 400</code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        f = [0] * (n + 1)
        f[1] = nums[0]
        for i in range(2, n + 1):
            f[i] = max(f[i - 1], f[i - 2] + nums[i - 1])
        return f[n]
```

```python
class Solution:
    def rob(self, nums: List[int]) -> int:
        f, g = 0, nums[0]
        for x in nums[1:]:
            f, g = g, max(f + x, g)
        return g
```

### **Java**

```java
class Solution {
    public int rob(int[] nums) {
        int n = nums.length;
        int[] f = new int[n + 1];
        f[1] = nums[0];
        for (int i = 2; i <= n; ++i) {
            f[i] = Math.max(f[i - 1], f[i - 2] + nums[i - 1]);
        }
        return f[n];
    }
}
```

```java
class Solution {
    public int rob(int[] nums) {
        int f = 0, g = nums[0];
        for (int i = 1; i < nums.length; ++i) {
            int t = g;
            g = Math.max(g, f + nums[i]);
            f = t;
        }
        return g;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int rob(vector<int>& nums) {
        int n = nums.size();
        int f[n + 1];
        memset(f, 0, sizeof(f));
        f[1] = nums[0];
        for (int i = 2; i <= n; ++i) {
            f[i] = max(f[i - 1], f[i - 2] + nums[i - 1]);
        }
        return f[n];
    }
};
```

```cpp
class Solution {
public:
    int rob(vector<int>& nums) {
        int f = 0, g = nums[0];
        for (int i = 1; i < nums.size(); ++i) {
            int t = g;
            g = max(g, f + nums[i]);
            f = t;
        }
        return g;
    }
};
```

### **Go**

```go
func rob(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	f[1] = nums[0]
	for i := 2; i <= n; i++ {
		f[i] = max(f[i-1], f[i-2]+nums[i-1])
	}
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

```go
func rob(nums []int) int {
	f, g := 0, nums[0]
	for _, x := range nums[1:] {
		f, g = g, max(f+x, g)
	}
	return g
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### **TypeScript**

```ts
function rob(nums: number[]): number {
    const n = nums.length;
    const f: number[] = Array(n + 1).fill(0);
    f[1] = nums[0];
    for (let i = 2; i <= n; ++i) {
        f[i] = Math.max(f[i - 1], f[i - 2] + nums[i - 1]);
    }
    return f[n];
}
```

```ts
function rob(nums: number[]): number {
    let [f, g] = [0, nums[0]];
    for (let i = 1; i < nums.length; ++i) {
        [f, g] = [g, Math.max(f + nums[i], g)];
    }
    return g;
}
```

### **Rust**

```rust
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let mut f = [0, 0];
        for x in nums {
            f = [f[1], f[1].max(f[0] + x)]
        }
        f[1]
    }
}
```

### **...**

```

```

<!-- tabs:end -->
