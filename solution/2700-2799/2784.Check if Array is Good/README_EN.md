# [2784. Check if Array is Good](https://leetcode.com/problems/check-if-array-is-good)

[中文文档](/solution/2700-2799/2784.Check%20if%20Array%20is%20Good/README.md)

## Description

<p>You are given an integer array <code>nums</code>. We consider an array <strong>good </strong>if it is a permutation of an array <code>base[n]</code>.</p>

<p><code>base[n] = [1, 2, ..., n - 1, n, n] </code>(in other words, it is an array of length <code>n + 1</code> which contains <code>1</code> to <code>n - 1 </code>exactly once, plus two occurrences of <code>n</code>). For example, <code>base[1] = [1, 1]</code> and<code> base[3] = [1, 2, 3, 3]</code>.</p>

<p>Return <code>true</code> <em>if the given array is good, otherwise return</em><em> </em><code>false</code>.</p>

<p><strong>Note: </strong>A permutation of integers represents an arrangement of these numbers.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [2, 1, 3]
<strong>Output:</strong> false
<strong>Explanation:</strong> Since the maximum element of the array is 3, the only candidate n for which this array could be a permutation of base[n], is n = 3. However, base[3] has four elements but array nums has three. Therefore, it can not be a permutation of base[3] = [1, 2, 3, 3]. So the answer is false.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1, 3, 3, 2]
<strong>Output:</strong> true
<strong>Explanation:</strong> Since the maximum element of the array is 3, the only candidate n for which this array could be a permutation of base[n], is n = 3. It can be seen that nums is a permutation of base[3] = [1, 2, 3, 3] (by swapping the second and fourth elements in nums, we reach base[3]). Therefore, the answer is true.</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [1, 1]
<strong>Output:</strong> true
<strong>Explanation:</strong> Since the maximum element of the array is 1, the only candidate n for which this array could be a permutation of base[n], is n = 1. It can be seen that nums is a permutation of base[1] = [1, 1]. Therefore, the answer is true.</pre>

<p><strong class="example">Example 4:</strong></p>

<pre>
<strong>Input:</strong> nums = [3, 4, 4, 1, 2, 1]
<strong>Output:</strong> false
<strong>Explanation:</strong> Since the maximum element of the array is 4, the only candidate n for which this array could be a permutation of base[n], is n = 4. However, base[4] has five elements but array nums has six. Therefore, it can not be a permutation of base[4] = [1, 2, 3, 4, 4]. So the answer is false.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= num[i] &lt;= 200</code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def isGood(self, nums: List[int]) -> bool:
        n = len(nums) - 1
        cnt = Counter(nums)
        cnt[n] -= 2
        for i in range(1, n):
            cnt[i] -= 1
        return all(v == 0 for v in cnt.values())
```

### **Java**

```java
class Solution {
    public boolean isGood(int[] nums) {
        int n = nums.length - 1;
        int[] cnt = new int[201];
        for (int x : nums) {
            ++cnt[x];
        }
        cnt[n] -= 2;
        for (int i = 1; i < n; ++i) {
            cnt[i] -= 1;
        }
        for (int x : cnt) {
            if (x != 0) {
                return false;
            }
        }
        return true;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    bool isGood(vector<int>& nums) {
        int n = nums.size() - 1;
        vector<int> cnt(201);
        for (int x : nums) {
            ++cnt[x];
        }
        cnt[n] -= 2;
        for (int i = 1; i < n; ++i) {
            --cnt[i];
        }
        for (int x : cnt) {
            if (x) {
                return false;
            }
        }
        return true;
    }
};
```

### **Go**

```go
func isGood(nums []int) bool {
	n := len(nums) - 1
	cnt := [201]int{}
	for _, x := range nums {
		cnt[x]++
	}
	cnt[n] -= 2
	for i := 1; i < n; i++ {
		cnt[i]--
	}
	for _, x := range cnt {
		if x != 0 {
			return false
		}
	}
	return true
}
```

### **TypeScript**

```ts
function isGood(nums: number[]): boolean {
    const n = nums.length - 1;
    const cnt: number[] = new Array(201).fill(0);
    for (const x of nums) {
        ++cnt[x];
    }
    cnt[n] -= 2;
    for (let i = 1; i < n; ++i) {
        cnt[i]--;
    }
    return cnt.every(x => x >= 0);
}
```

### **...**

```

```

<!-- tabs:end -->
