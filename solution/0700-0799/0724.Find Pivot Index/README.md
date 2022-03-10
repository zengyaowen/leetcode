# [724. 寻找数组的中心下标](https://leetcode-cn.com/problems/find-pivot-index)

[English Version](/solution/0700-0799/0724.Find%20Pivot%20Index/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给你一个整数数组&nbsp;<code>nums</code> ，请计算数组的 <strong>中心下标 </strong>。</p>

<p>数组<strong> 中心下标</strong><strong> </strong>是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。</p>

<p>如果中心下标位于数组最左端，那么左侧数之和视为 <code>0</code> ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。</p>

<p>如果数组有多个中心下标，应该返回 <strong>最靠近左边</strong> 的那一个。如果数组不存在中心下标，返回 <code>-1</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1, 7, 3, 6, 5, 6]
<strong>输出：</strong>3
<strong>解释：</strong>
中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1, 2, 3]
<strong>输出：</strong>-1
<strong>解释：</strong>
数组中不存在满足此条件的中心下标。</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [2, 1, -1]
<strong>输出：</strong>0
<strong>解释：</strong>
中心下标是 0 。
左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>-1000 &lt;= nums[i] &lt;= 1000</code></li>
</ul>

<p>&nbsp;</p>

<p><strong>注意：</strong>本题与主站 1991 题相同：<a href="https://leetcode-cn.com/problems/find-the-middle-index-in-array/" target="_blank">https://leetcode-cn.com/problems/find-the-middle-index-in-array/</a></p>

## 解法

<!-- 这里可写通用的实现逻辑 -->

前缀和。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        s, presum = sum(nums), 0
        for i, v in enumerate(nums):
            if (presum << 1) == s - v:
                return i
            presum += v
        return -1
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public int pivotIndex(int[] nums) {
        int n = nums.length, s = 0;
        for (int e : nums) {
            s += e;
        }
        int presum = 0;
        for (int i = 0; i < n; ++i) {
            // presum == sums - nums[i] - presum
            if (presum << 1 == s - nums[i]) {
                return i;
            }
            presum += nums[i];
        }
        return -1;
    }
}
```

### **TypeScript**

```ts
function pivotIndex(nums: number[]): number {
    let rightSum = nums.reduce((a, c) => a + c, 0),
        leftSum = 0;
    for (let i = 0; i < nums.length; i++) {
        rightSum -= nums[i];
        if (leftSum == rightSum) return i;
        leftSum += nums[i];
    }
    return -1;
}
```

### **C++**

```cpp
class Solution {
public:
    int pivotIndex(vector<int> &nums) {
        int s = 0;
        for (int e : nums)
            s += e;
        int presum = 0;
        for (int i = 0; i < nums.size(); ++i)
        {
            if (presum * 2 == s - nums[i])
                return i;
            presum += nums[i];
        }
        return -1;
    }
};
```

### **Go**

```go
func pivotIndex(nums []int) int {
	s := 0
	for _, e := range nums {
		s += e
	}
	presum := 0
	for i, e := range nums {
		if presum<<1 == s-e {
			return i
		}
		presum += e
	}
	return -1
}
```

### **...**

```

```

<!-- tabs:end -->
