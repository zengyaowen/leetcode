# [2670. 找出不同元素数目差数组](https://leetcode.cn/problems/find-the-distinct-difference-array)

[English Version](/solution/2600-2699/2670.Find%20the%20Distinct%20Difference%20Array/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>nums</code> ，数组长度为 <code>n</code> 。</p>

<p><code>nums</code> 的 <strong>不同元素数目差</strong> 数组可以用一个长度为 <code>n</code> 的数组 <code>diff</code> 表示，其中 <code>diff[i]</code> 等于前缀 <code>nums[0, ..., i]</code> 中不同元素的数目 <strong>减去</strong> 后缀 <code>nums[i + 1, ..., n - 1]</code> 中不同元素的数目。</p>

<p>返回<em> </em><code>nums</code> 的 <strong>不同元素数目差</strong> 数组。</p>

<p>注意 <code>nums[i, ..., j]</code> 表示 <code>nums</code> 的一个从下标 <code>i</code> 开始到下标 <code>j</code> 结束的子数组（包含下标 <code>i</code> 和 <code>j</code> 对应元素）。特别需要说明的是，如果 <code>i &gt; j</code> ，则 <code>nums[i, ..., j]</code> 表示一个空子数组。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3,4,5]
<strong>输出：</strong>[-3,-1,1,3,5]
<strong>解释：
</strong>对于 i = 0，前缀中有 1 个不同的元素，而在后缀中有 4 个不同的元素。因此，diff[0] = 1 - 4 = -3 。
对于 i = 1，前缀中有 2 个不同的元素，而在后缀中有 3 个不同的元素。因此，diff[1] = 2 - 3 = -1 。
对于 i = 2，前缀中有 3 个不同的元素，而在后缀中有 2 个不同的元素。因此，diff[2] = 3 - 2 = 1 。
对于 i = 3，前缀中有 4 个不同的元素，而在后缀中有 1 个不同的元素。因此，diff[3] = 4 - 1 = 3 。
对于 i = 4，前缀中有 5 个不同的元素，而在后缀中有 0 个不同的元素。因此，diff[4] = 5 - 0 = 5 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [3,2,3,4,2]
<strong>输出：</strong>[-2,-1,0,2,3]
<strong>解释：</strong>
对于 i = 0，前缀中有 1 个不同的元素，而在后缀中有 3 个不同的元素。因此，diff[0] = 1 - 3 = -2 。
对于 i = 1，前缀中有 2 个不同的元素，而在后缀中有 3 个不同的元素。因此，diff[1] = 2 - 3 = -1 。
对于 i = 2，前缀中有 2 个不同的元素，而在后缀中有 2 个不同的元素。因此，diff[2] = 2 - 2 = 0 。
对于 i = 3，前缀中有 3 个不同的元素，而在后缀中有 1 个不同的元素。因此，diff[3] = 3 - 1 = 2 。
对于 i = 4，前缀中有 3 个不同的元素，而在后缀中有 0 个不同的元素。因此，diff[4] = 3 - 0 = 3 。 
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length&nbsp;&lt;= 50</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 50</code></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def distinctDifferenceArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n
        for i in range(n):
            a = len(set(nums[: i + 1]))
            b = len(set(nums[i + 1 :]))
            ans[i] = a - b
        return ans
```

```python
class Solution:
    def distinctDifferenceArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        suf = [0] * (n + 1)
        s = set()
        for i in range(n - 1, -1, -1):
            s.add(nums[i])
            suf[i] = len(s)

        s.clear()
        ans = [0] * n
        for i, x in enumerate(nums):
            s.add(x)
            ans[i] = len(s) - suf[i + 1]
        return ans
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public int[] distinctDifferenceArray(int[] nums) {
        int n = nums.length;
        int[] suf = new int[n + 1];
        Set<Integer> s = new HashSet<>();
        for (int i = n - 1; i >= 0; --i) {
            s.add(nums[i]);
            suf[i] = s.size();
        }
        s.clear();
        int[] ans = new int[n];
        for (int i = 0; i < n; ++i) {
            s.add(nums[i]);
            ans[i] = s.size() - suf[i + 1];
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<int> distinctDifferenceArray(vector<int>& nums) {
        int n = nums.size();
        vector<int> suf(n + 1);
        unordered_set<int> s;
        for (int i = n - 1; i >= 0; --i) {
            s.insert(nums[i]);
            suf[i] = s.size();
        }
        s.clear();
        vector<int> ans(n);
        for (int i = 0; i < n; ++i) {
            s.insert(nums[i]);
            ans[i] = s.size() - suf[i + 1];
        }
        return ans;
    }
};
```

### **Go**

```go
func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suf := make([]int, n+1)
	s := map[int]bool{}
	for i := n - 1; i >= 0; i-- {
		s[nums[i]] = true
		suf[i] = len(s)
	}
	ans := make([]int, n)
	s = map[int]bool{}
	for i, x := range nums {
		s[x] = true
		ans[i] = len(s) - suf[i+1]
	}
	return ans
}
```

### **TypeScript**

```ts
function distinctDifferenceArray(nums: number[]): number[] {
    const n = nums.length;
    const suf: number[] = new Array(n + 1).fill(0);
    const s: Set<number> = new Set();
    for (let i = n - 1; i >= 0; --i) {
        s.add(nums[i]);
        suf[i] = s.size;
    }
    s.clear();
    const ans: number[] = new Array(n);
    for (let i = 0; i < n; ++i) {
        s.add(nums[i]);
        ans[i] = s.size - suf[i + 1];
    }
    return ans;
}
```

### **...**

```

```

<!-- tabs:end -->
