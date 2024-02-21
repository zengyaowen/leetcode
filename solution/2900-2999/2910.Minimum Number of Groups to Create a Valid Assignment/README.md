# [2910. 合法分组的最少组数](https://leetcode.cn/problems/minimum-number-of-groups-to-create-a-valid-assignment)

[English Version](/solution/2900-2999/2910.Minimum%20Number%20of%20Groups%20to%20Create%20a%20Valid%20Assignment/README_EN.md)

<!-- tags:贪心,数组,哈希表 -->

## 题目描述

<!-- 这里写题目描述 -->

<p>给你一个长度为 <code>n</code>&nbsp;下标从 <strong>0</strong>&nbsp;开始的整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>我们想将下标进行分组，使得&nbsp;<code>[0, n - 1]</code>&nbsp;内所有下标&nbsp;<code>i</code>&nbsp;都 <strong>恰好</strong> 被分到其中一组。</p>

<p>如果以下条件成立，我们说这个分组方案是合法的：</p>

<ul>
	<li>对于每个组&nbsp;<code>g</code>&nbsp;，同一组内所有下标在&nbsp;<code>nums</code>&nbsp;中对应的数值都相等。</li>
	<li>对于任意两个组&nbsp;<code>g<sub>1</sub></code> 和&nbsp;<code>g<sub>2</sub></code>&nbsp;，两个组中&nbsp;<strong>下标数量</strong> 的&nbsp;<strong>差值不超过&nbsp;</strong><code>1</code>&nbsp;。</li>
</ul>

<p>请你返回一个整数，表示得到一个合法分组方案的 <strong>最少</strong>&nbsp;组数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,2,3,2,3]
<b>输出：</b>2
<b>解释：</b>一个得到 2 个分组的方案如下，中括号内的数字都是下标：
组 1 -&gt; [0,2,4]
组 2 -&gt; [1,3]
所有下标都只属于一个组。
组 1 中，nums[0] == nums[2] == nums[4] ，所有下标对应的数值都相等。
组 2 中，nums[1] == nums[3] ，所有下标对应的数值都相等。
组 1 中下标数目为 3 ，组 2 中下标数目为 2 。
两者之差不超过 1 。
无法得到一个小于 2 组的答案，因为如果只有 1 组，组内所有下标对应的数值都要相等。
所以答案为 2 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [10,10,10,3,1,1]
<b>输出：</b>4
<b>解释：</b>一个得到 2 个分组的方案如下，中括号内的数字都是下标：
组 1 -&gt; [0]
组 2 -&gt; [1,2]
组 3 -&gt; [3]
组 4 -&gt; [4,5]
分组方案满足题目要求的两个条件。
无法得到一个小于 4 组的答案。
所以答案为 4 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

## 解法

### 方法一：哈希表 + 枚举

我们用一个哈希表 $cnt$ 统计数组 $nums$ 中每个数字出现的次数，我们记数字次数的最小值为 $k$，那么我们可以在 $[k,..1]$ 的范围内枚举分组的大小。由于每个组的大小差值不超过 $1$，那么分组大小为 $k$ 或 $k+1$。

对于当前枚举到的分组大小 $k$，我们遍历哈希表中的每个次数 $v$，如果 $\lfloor \frac{v}{k} \rfloor < v \bmod k$，那么说明无法将这个次数 $v$ 分成 $k$ 个或 $k+1$ 个数值相同的组，因此我们可以直接跳过这个分组大小 $k$。否则，说明可以分组，我们只需要尽可能分出最多的分组大小 $k+1$，即可保证得到最小的分组数，因此我们可以将 $v$ 个数分成 $\lceil \frac{v}{k+1} \rceil$ 组，累加到当前枚举的答案中。由于我们是按照 $k$ 从大到小枚举的，因此只要找到了一个合法的分组方案，那么一定是最优的。

时间复杂度 $O(n)$，空间复杂度 $O(n)$。其中 $n$ 为数组 $nums$ 的长度。

<!-- tabs:start -->

```python
class Solution:
    def minGroupsForValidAssignment(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        for k in range(min(cnt.values()), 0, -1):
            ans = 0
            for v in cnt.values():
                if v // k < v % k:
                    ans = 0
                    break
                ans += (v + k) // (k + 1)
            if ans:
                return ans
```

```java
class Solution {
    public int minGroupsForValidAssignment(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum);
        }
        int k = nums.length;
        for (int v : cnt.values()) {
            k = Math.min(k, v);
        }
        for (;; --k) {
            int ans = 0;
            for (int v : cnt.values()) {
                if (v / k < v % k) {
                    ans = 0;
                    break;
                }
                ans += (v + k) / (k + 1);
            }
            if (ans > 0) {
                return ans;
            }
        }
    }
}
```

```cpp
class Solution {
public:
    int minGroupsForValidAssignment(vector<int>& nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }
        int k = 1e9;
        for (auto& [_, v] : cnt) {
            ans = min(ans, v);
        }
        for (;; --k) {
            int ans = 0;
            for (auto& [_, v] : cnt) {
                if (v / k < v % k) {
                    ans = 0;
                    break;
                }
                ans += (v + k) / (k + 1);
            }
            if (ans) {
                return ans;
            }
        }
    }
};
```

```go
func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, v := range cnt {
		k = min(k, v)
	}
	for ; ; k-- {
		ans := 0
		for _, v := range cnt {
			if v/k < v%k {
				ans = 0
				break
			}
			ans += (v + k) / (k + 1)
		}
		if ans > 0 {
			return ans
		}
	}
}
```

```ts
function minGroupsForValidAssignment(nums: number[]): number {
    const cnt: Map<number, number> = new Map();
    for (const x of nums) {
        cnt.set(x, (cnt.get(x) || 0) + 1);
    }
    for (let k = Math.min(...cnt.values()); ; --k) {
        let ans = 0;
        for (const [_, v] of cnt) {
            if (((v / k) | 0) < v % k) {
                ans = 0;
                break;
            }
            ans += Math.ceil(v / (k + 1));
        }
        if (ans) {
            return ans;
        }
    }
}
```

```rust
use std::collections::HashMap;

impl Solution {
    pub fn min_groups_for_valid_assignment(nums: Vec<i32>) -> i32 {
        let mut cnt: HashMap<i32, i32> = HashMap::new();

        for x in nums.iter() {
            let count = cnt.entry(*x).or_insert(0);
            *count += 1;
        }

        let mut k = i32::MAX;

        for &v in cnt.values() {
            k = k.min(v);
        }

        for k in (1..=k).rev() {
            let mut ans = 0;

            for &v in cnt.values() {
                if v / k < v % k {
                    ans = 0;
                    break;
                }

                ans += (v + k) / (k + 1);
            }

            if ans > 0 {
                return ans;
            }
        }

        0
    }
}
```

<!-- tabs:end -->

<!-- end -->
