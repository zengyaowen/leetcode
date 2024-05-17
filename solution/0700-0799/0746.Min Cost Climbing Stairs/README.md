---
comments: true
difficulty: 简单
edit_url: https://github.com/doocs/leetcode/edit/main/solution/0700-0799/0746.Min%20Cost%20Climbing%20Stairs/README.md
tags:
    - 数组
    - 动态规划
---

<!-- problem:start -->

# [746. 使用最小花费爬楼梯](https://leetcode.cn/problems/min-cost-climbing-stairs)

[English Version](/solution/0700-0799/0746.Min%20Cost%20Climbing%20Stairs/README_EN.md)

## 题目描述

<!-- description:start -->

<p>给你一个整数数组 <code>cost</code> ，其中 <code>cost[i]</code> 是从楼梯第 <code>i</code> 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。</p>

<p>你可以选择从下标为 <code>0</code> 或下标为 <code>1</code> 的台阶开始爬楼梯。</p>

<p>请你计算并返回达到楼梯顶部的最低花费。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>cost = [10,<em><strong>15</strong></em>,20]
<strong>输出：</strong>15
<strong>解释：</strong>你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>cost = [<em><strong>1</strong></em>,100,<em><strong>1</strong></em>,1,<em><strong>1</strong></em>,100,<em><strong>1</strong></em>,<em><strong>1</strong></em>,100,<em><strong>1</strong></em>]
<strong>输出：</strong>6
<strong>解释：</strong>你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= cost.length &lt;= 1000</code></li>
	<li><code>0 &lt;= cost[i] &lt;= 999</code></li>
</ul>

<!-- description:end -->

## 解法

<!-- solution:start -->

### 方法一：动态规划

我们定义 $f[i]$ 表示到达第 $i$ 个阶梯所需要的最小花费，初始时 $f[0] = f[1] = 0$，答案即为 $f[n]$。

当 $i \ge 2$ 时，我们可以从第 $i - 1$ 个阶梯使用 $1$ 步直接到达第 $i$ 个阶梯，或者从第 $i - 2$ 个阶梯使用 $2$ 步到达第 $i$ 个阶梯，因此我们有状态转移方程：

$$
f[i] = \min(f[i - 1] + cost[i - 1], f[i - 2] + cost[i - 2])
$$

最终的答案即为 $f[n]$。

时间复杂度 $O(n)$，空间复杂度 $O(n)$。其中 $n$ 是数组 `cost` 的长度。

我们注意到，状态转移方程中的 $f[i]$ 只和 $f[i - 1]$ 与 $f[i - 2]$ 有关，因此我们可以使用两个变量 $f$ 和 $g$ 交替地记录 $f[i - 2]$ 和 $f[i - 1]$ 的值，这样空间复杂度可以优化到 $O(1)$。

<!-- tabs:start -->

#### Python3

```python
class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        n = len(cost)
        f = [0] * (n + 1)
        for i in range(2, n + 1):
            f[i] = min(f[i - 2] + cost[i - 2], f[i - 1] + cost[i - 1])
        return f[n]
```

#### Java

```java
class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int n = cost.length;
        int[] f = new int[n + 1];
        for (int i = 2; i <= n; ++i) {
            f[i] = Math.min(f[i - 2] + cost[i - 2], f[i - 1] + cost[i - 1]);
        }
        return f[n];
    }
}
```

#### C++

```cpp
class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        int n = cost.size();
        vector<int> f(n + 1);
        for (int i = 2; i <= n; ++i) {
            f[i] = min(f[i - 2] + cost[i - 2], f[i - 1] + cost[i - 1]);
        }
        return f[n];
    }
};
```

#### Go

```go
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 2; i <= n; i++ {
		f[i] = min(f[i-1]+cost[i-1], f[i-2]+cost[i-2])
	}
	return f[n]
}
```

#### TypeScript

```ts
function minCostClimbingStairs(cost: number[]): number {
    const n = cost.length;
    const f: number[] = Array(n + 1).fill(0);
    for (let i = 2; i <= n; ++i) {
        f[i] = Math.min(f[i - 1] + cost[i - 1], f[i - 2] + cost[i - 2]);
    }
    return f[n];
}
```

#### Rust

```rust
impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let n = cost.len();
        let mut f = vec![0; n + 1];
        for i in 2..=n {
            f[i] = std::cmp::min(f[i - 2] + cost[i - 2], f[i - 1] + cost[i - 1]);
        }
        f[n]
    }
}
```

<!-- tabs:end -->

<!-- solution:end -->

<!-- solution:start -->

### 方法二

<!-- tabs:start -->

#### Python3

```python
class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        f = g = 0
        for i in range(2, len(cost) + 1):
            f, g = g, min(f + cost[i - 2], g + cost[i - 1])
        return g
```

#### Java

```java
class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int f = 0, g = 0;
        for (int i = 2; i <= cost.length; ++i) {
            int gg = Math.min(f + cost[i - 2], g + cost[i - 1]);
            f = g;
            g = gg;
        }
        return g;
    }
}
```

#### C++

```cpp
class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        int f = 0, g = 0;
        for (int i = 2; i <= cost.size(); ++i) {
            int gg = min(f + cost[i - 2], g + cost[i - 1]);
            f = g;
            g = gg;
        }
        return g;
    }
};
```

#### Go

```go
func minCostClimbingStairs(cost []int) int {
	var f, g int
	for i := 2; i <= n; i++ {
		f, g = g, min(f+cost[i-2], g+cost[i-1])
	}
	return g
}
```

#### TypeScript

```ts
function minCostClimbingStairs(cost: number[]): number {
    let a = 0,
        b = 0;
    for (let i = 1; i < cost.length; ++i) {
        [a, b] = [b, Math.min(a + cost[i - 1], b + cost[i])];
    }
    return b;
}
```

#### Rust

```rust
impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let (mut f, mut g) = (0, 0);
        for i in 2..=cost.len() {
            let gg = std::cmp::min(f + cost[i - 2], g + cost[i - 1]);
            f = g;
            g = gg;
        }
        g
    }
}
```

<!-- tabs:end -->

<!-- solution:end -->

<!-- problem:end -->
