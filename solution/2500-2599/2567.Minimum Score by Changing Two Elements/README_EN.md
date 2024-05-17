---
comments: true
difficulty: Medium
edit_url: https://github.com/doocs/leetcode/edit/main/solution/2500-2599/2567.Minimum%20Score%20by%20Changing%20Two%20Elements/README_EN.md
rating: 1608
source: Biweekly Contest 98 Q2
tags:
    - Greedy
    - Array
    - Sorting
---

<!-- problem:start -->

# [2567. Minimum Score by Changing Two Elements](https://leetcode.com/problems/minimum-score-by-changing-two-elements)

[中文文档](/solution/2500-2599/2567.Minimum%20Score%20by%20Changing%20Two%20Elements/README.md)

## Description

<!-- description:start -->

<p>You are given a <strong>0-indexed</strong> integer array <code>nums</code>.</p>

<ul>
	<li>The <strong>low</strong> score of <code><font face="monospace">nums</font></code> is the minimum value of <code>|nums[i]&nbsp;- nums[j]|</code> over all <code>0 &lt;= i &lt; j &lt; nums.length</code>.</li>
	<li>The <strong>high</strong> score of&nbsp;<code><font face="monospace">nums</font></code> is the maximum value of <code>|nums[i]&nbsp;- nums[j]|</code> over all <code>0 &lt;= i &lt; j &lt; nums.length</code>.</li>
	<li>The <strong>score</strong> of <code>nums</code> is the sum of the <strong>high</strong> and <strong>low</strong> scores of nums.</li>
</ul>

<p>To minimize the score of <code>nums</code>, we can change the value of <strong>at most two</strong> elements of <code>nums</code>.</p>

<p>Return <em>the <strong>minimum</strong> possible <strong>score</strong> after changing&nbsp;the value of <strong>at most two</strong> elements o</em>f <code>nums</code>.</p>

<p>Note that <code>|x|</code> denotes the absolute value of <code>x</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,4,3]
<strong>Output:</strong> 0
<strong>Explanation:</strong> Change value of nums[1] and nums[2] to 1 so that nums becomes [1,1,1]. Now, the value of <code>|nums[i] - nums[j]|</code> is always equal to 0, so we return 0 + 0 = 0.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,4,7,8,5]
<strong>Output:</strong> 3
<strong>Explanation:</strong> Change nums[0] and nums[1] to be 6. Now nums becomes [6,6,7,8,5].
Our low score is achieved when i = 0 and j = 1, in which case |<code>nums[i] - nums[j]</code>| = |6 - 6| = 0.
Our high score is achieved when i = 3 and j = 4, in which case |<code>nums[i] - nums[j]</code>| = |8 - 5| = 3.
The sum of our high and low score is 3, which we can prove to be minimal.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

<!-- description:end -->

## Solutions

<!-- solution:start -->

### Solution 1: Sorting + Greedy

From the problem description, we know that the minimum score is actually the minimum difference between two adjacent elements in the sorted array, and the maximum score is the difference between the first and last elements of the sorted array. The score of the array $nums$ is the sum of the minimum score and the maximum score.

Therefore, we can first sort the array. Since the problem allows us to modify the values of at most two elements in the array, we can modify a number to make it the same as another number in the array, making the minimum score $0$. In this case, the score of the array $nums$ is actually the maximum score. We can choose to make one of the following modifications:

Modify the smallest two numbers to $nums[2]$, then the maximum score is $nums[n - 1] - nums[2]$;
Modify the smallest number to $nums[1]$ and the largest number to $nums[n - 2]$, then the maximum score is $nums[n - 2] - nums[1]$;
Modify the largest two numbers to $nums[n - 3]$, then the maximum score is $nums[n - 3] - nums[0]$.
Finally, we return the minimum score of the above three modifications.

The time complexity is $O(n \log n)$, and the space complexity is $O(\log n)$. Here, $n$ is the length of the array $nums$.

Similar problems:

-[1509. Minimum Difference Between Largest and Smallest Value in Three Moves](https://github.com/doocs/leetcode/blob/main/solution/1500-1599/1509.Minimum%20Difference%20Between%20Largest%20and%20Smallest%20Value%20in%20Three%20Moves/README_EN.md)

<!-- tabs:start -->

#### Python3

```python
class Solution:
    def minimizeSum(self, nums: List[int]) -> int:
        nums.sort()
        return min(nums[-1] - nums[2], nums[-2] - nums[1], nums[-3] - nums[0])
```

#### Java

```java
class Solution {
    public int minimizeSum(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        int a = nums[n - 1] - nums[2];
        int b = nums[n - 2] - nums[1];
        int c = nums[n - 3] - nums[0];
        return Math.min(a, Math.min(b, c));
    }
}
```

#### C++

```cpp
class Solution {
public:
    int minimizeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        return min({nums[n - 1] - nums[2], nums[n - 2] - nums[1], nums[n - 3] - nums[0]});
    }
};
```

#### Go

```go
func minimizeSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return min(nums[n-1]-nums[2], min(nums[n-2]-nums[1], nums[n-3]-nums[0]))
}
```

#### TypeScript

```ts
function minimizeSum(nums: number[]): number {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    return Math.min(nums[n - 3] - nums[0], nums[n - 2] - nums[1], nums[n - 1] - nums[2]);
}
```

#### Rust

```rust
impl Solution {
    pub fn minimize_sum(mut nums: Vec<i32>) -> i32 {
        nums.sort();
        let n = nums.len();
        (nums[n - 1] - nums[2]).min(nums[n - 2] - nums[1]).min(nums[n - 3] - nums[0])
    }
}
```

#### C

```c
#define min(a, b) (((a) < (b)) ? (a) : (b))

int cmp(const void* a, const void* b) {
    return *(int*) a - *(int*) b;
}

int minimizeSum(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    return min(nums[numsSize - 1] - nums[2], min(nums[numsSize - 2] - nums[1], nums[numsSize - 3] - nums[0]));
}
```

<!-- tabs:end -->

<!-- solution:end -->

<!-- problem:end -->
