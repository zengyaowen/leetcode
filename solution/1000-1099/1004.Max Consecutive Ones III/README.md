# [1004. 最大连续 1 的个数 III](https://leetcode-cn.com/problems/max-consecutive-ones-iii)

[English Version](/solution/1000-1099/1004.Max%20Consecutive%20Ones%20III/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给定一个二进制数组&nbsp;<code>nums</code>&nbsp;和一个整数&nbsp;<code><font color="#c7254e"><font face="Menlo, Monaco, Consolas, Courier New, monospace"><span style="font-size:12.6px"><span style="background-color:#f9f2f4">k</span></span></font></font></code>&nbsp;，如果可以翻转最多<code><font color="#c7254e"><font face="Menlo, Monaco, Consolas, Courier New, monospace"><span style="font-size:12.6px"><span style="background-color:#f9f2f4">k</span></span></font></font></code>&nbsp;个 <code>0</code> ，则返回 <em>数组中连续 <code>1</code> 的最大个数</em> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
<strong>输出：</strong>6
<strong>解释：</strong>[1,1,1,0,0,<strong>1</strong>,1,1,1,1,<strong>1</strong>]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
<strong>输出：</strong>10
<strong>解释：</strong>[0,0,1,1,<strong>1</strong>,<strong>1</strong>,1,1,1,<strong>1</strong>,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>nums[i]</code>&nbsp;不是&nbsp;<code>0</code>&nbsp;就是&nbsp;<code>1</code></li>
	<li><code>0 &lt;= k &lt;= nums.length</code></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

思路同 [2024. 考试的最大困扰度](/solution/2000-2099/2024.Maximize%20the%20Confusion%20of%20an%20Exam/README.md)

维护一个单调变长的窗口。这种窗口经常出现在寻求“最大窗口”的问题中：因为要求的是”最大“，所以我们没有必要缩短窗口，于是代码就少了缩短窗口的部分；从另一个角度讲，本题里的 K 是资源数，一旦透支，窗口就不能再增长了。

-   l 是窗口左端点，负责移动起始位置
-   r 是窗口右端点，负责扩展窗口
-   k 是资源数，每次要替换 0，k 减 1，同时 r 向右移动
-   `r++` 每次都会执行，`l++` 只有资源 `k < 0` 时才触发，因此 `r - l` 的值只会单调递增（或保持不变）
-   移动左端点时，如果当前元素是 0，说明可以释放一个资源，k 加 1

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        l = r = -1
        while r < len(nums) - 1:
            r += 1
            if nums[r] == 0:
                k -= 1
            if k < 0:
                l += 1
                if nums[l] == 0:
                    k += 1
        return r - l
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public int longestOnes(int[] nums, int k) {
        int l = 0, r = 0;
        while (r < nums.length) {
            if (nums[r++] == 0) {
                --k;
            }
            if (k < 0 && nums[l++] == 0) {
                ++k;
            }
        }
        return r - l;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int l = 0, r = 0;
        while (r < nums.size())
        {
            if (nums[r++] == 0) --k;
            if (k < 0 && nums[l++] == 0) ++k;
        }
        return r - l;
    }
};
```

### **Go**

```go
func longestOnes(nums []int, k int) int {
	l, r := -1, -1
	for r < len(nums)-1 {
		r++
		if nums[r] == 0 {
			k--
		}
		if k < 0 {
			l++
			if nums[l] == 0 {
				k++
			}
		}
	}
	return r - l
}
```

### **TypeScript**

```ts
function longestOnes(nums: number[], k: number): number {
    const n = nums.length;
    let l = 0;
    for (const num of nums) {
        if (num === 0) {
            k--;
        }
        if (k < 0 && nums[l++] === 0) {
            k++;
        }
    }
    return n - l;
}
```

```ts
function longestOnes(nums: number[], k: number): number {
    const n = nums.length;
    let l = 0;
    let res = k;
    const count = [0, 0];
    for (let r = 0; r < n; r++) {
        count[nums[r]]++;
        res = Math.max(res, r - l);
        while (count[0] > k) {
            count[nums[l]]--;
            l++;
        }
    }
    return Math.max(res, n - l);
}
```

### **Rust**

```rust
impl Solution {
    pub fn longest_ones(nums: Vec<i32>, mut k: i32) -> i32 {
        let n = nums.len();
        let mut l = 0;
        for num in nums.iter() {
            if num == &0 {
                k -= 1;
            }
            if k < 0 {
                if nums[l] == 0 {
                    k += 1;
                }
                l += 1;
            }
        }
        (n - l) as i32
    }
}
```

### **...**

```

```

<!-- tabs:end -->
