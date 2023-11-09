# [907. Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums)

[中文文档](/solution/0900-0999/0907.Sum%20of%20Subarray%20Minimums/README.md)

## Description

<p>Given an array of integers arr, find the sum of <code>min(b)</code>, where <code>b</code> ranges over every (contiguous) subarray of <code>arr</code>. Since the answer may be large, return the answer <strong>modulo</strong> <code>10<sup>9</sup> + 7</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> arr = [3,1,2,4]
<strong>Output:</strong> 17
<strong>Explanation:</strong> 
Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]. 
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Sum is 17.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> arr = [11,81,94,43,3]
<strong>Output:</strong> 444
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= arr.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>1 &lt;= arr[i] &lt;= 3 * 10<sup>4</sup></code></li>
</ul>

## Solutions

The problem asks for the sum of the minimum values of each subarray, which is actually equivalent to finding the number of subarrays for each element $arr[i]$ where $arr[i]$ is the minimum, multiplying each by $arr[i]$, and then summing these products.

Thus, the focus of the problem is translated to finding the number of subarrays for which $arr[i]$ is the minimum.

For each $arr[i]$, we identify the first position $left[i]$ to its left that is smaller than $arr[i]$ and the first position $right[i]$ to its right that is less than or equal to $arr[i]$.

The number of subarrays where $arr[i]$ is the minimum can then be given by $(i - left[i]) \times (right[i] - i)$.

It's important to note why we are looking for the first position $right[i]$ that is less than or equal to $arr[i]$ and not less than $arr[i]$.

If we were to look for the first position less than $arr[i]$, we would end up double-counting.

For instance, consider the following array:

The element at index $3$ is $2$, and the first element less than $2$ to its left is at index $0$. If we find the first element less than $2$ to its right, we would end up at index $7$. That means the subarray interval is $(0, 7)$. Note that this is an open interval.

```
0 4 3 2 5 3 2 1
*     ^       *
```

If we calculate the subarray interval for the element at index $6$ using the same method, we would find that its interval is also $(0, 7)$.

```
0 4 3 2 5 3 2 1
*           ^ *
```

Therefore, the subarray intervals of the elements at index $3$ and $6$ are overlapping, leading to double-counting.

If we were to find the first element less than or equal to $arr[i]$ to its right, we wouldn't have this problem.

The subarray interval for the element at index $3$ would become $(0, 6)$ and for the element at index $6$ it would be $(0, 7)$, and these two are not overlapping.

To solve this problem, we just need to traverse the array.

For each element $arr[i]$, we use a monotonic stack to find its $left[i]$ and $right[i]$.

Then the number of subarrays where $arr[i]$ is the minimum can be calculated by $(i - left[i]) \times (right[i] - i)$. Multiply this by $arr[i]$ and sum these values for all $i$ to get the final answer.

Remember to take care of data overflow and modulus operation.

The time complexity is $O(n)$, where $n$ represents the length of the array $arr$.

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        n = len(arr)
        left = [-1] * n
        right = [n] * n
        stk = []
        for i, v in enumerate(arr):
            while stk and arr[stk[-1]] >= v:
                stk.pop()
            if stk:
                left[i] = stk[-1]
            stk.append(i)

        stk = []
        for i in range(n - 1, -1, -1):
            while stk and arr[stk[-1]] > arr[i]:
                stk.pop()
            if stk:
                right[i] = stk[-1]
            stk.append(i)
        mod = 10**9 + 7
        return sum((i - left[i]) * (right[i] - i) * v for i, v in enumerate(arr)) % mod
```

### **Java**

```java
class Solution {
    public int sumSubarrayMins(int[] arr) {
        int n = arr.length;
        int[] left = new int[n];
        int[] right = new int[n];
        Arrays.fill(left, -1);
        Arrays.fill(right, n);
        Deque<Integer> stk = new ArrayDeque<>();
        for (int i = 0; i < n; ++i) {
            while (!stk.isEmpty() && arr[stk.peek()] >= arr[i]) {
                stk.pop();
            }
            if (!stk.isEmpty()) {
                left[i] = stk.peek();
            }
            stk.push(i);
        }
        stk.clear();
        for (int i = n - 1; i >= 0; --i) {
            while (!stk.isEmpty() && arr[stk.peek()] > arr[i]) {
                stk.pop();
            }
            if (!stk.isEmpty()) {
                right[i] = stk.peek();
            }
            stk.push(i);
        }
        int mod = (int) 1e9 + 7;
        long ans = 0;
        for (int i = 0; i < n; ++i) {
            ans += (long) (i - left[i]) * (right[i] - i) % mod * arr[i] % mod;
            ans %= mod;
        }
        return (int) ans;
    }
}
```

### **C++**

```cpp
using ll = long long;
const int mod = 1e9 + 7;

class Solution {
public:
    int sumSubarrayMins(vector<int>& arr) {
        int n = arr.size();
        vector<int> left(n, -1);
        vector<int> right(n, n);
        stack<int> stk;
        for (int i = 0; i < n; ++i) {
            while (!stk.empty() && arr[stk.top()] >= arr[i]) stk.pop();
            if (!stk.empty()) left[i] = stk.top();
            stk.push(i);
        }
        stk = stack<int>();
        for (int i = n - 1; i >= 0; --i) {
            while (!stk.empty() && arr[stk.top()] > arr[i]) stk.pop();
            if (!stk.empty()) right[i] = stk.top();
            stk.push(i);
        }
        ll ans = 0;
        for (int i = 0; i < n; ++i) {
            ans += (ll) (i - left[i]) * (right[i] - i) * arr[i] % mod;
            ans %= mod;
        }
        return ans;
    }
};
```

### **Rust**

```rust
const MOD: i64 = (1e9 as i64) + 7;

impl Solution {
    pub fn sum_subarray_mins(arr: Vec<i32>) -> i32 {
        let n: usize = arr.len();
        let mut ret: i64 = 0;
        let mut left: Vec<i32> = vec![-1; n];
        let mut right: Vec<i32> = vec![n as i32; n];
        // Index stack, store the index of the value in the given array
        let mut stack: Vec<i32> = Vec::new();

        // Find the first element that's less than the current value for the left side
        // The default value of which is -1
        for i in 0..n {
            while !stack.is_empty() && arr[*stack.last().unwrap() as usize] >= arr[i] {
                stack.pop();
            }
            if !stack.is_empty() {
                left[i] = *stack.last().unwrap();
            }
            stack.push(i as i32);
        }

        stack.clear();

        // Find the first element that's less or equal than the current value for the right side
        // The default value of which is n
        for i in (0..n).rev() {
            while !stack.is_empty() && arr[*stack.last().unwrap() as usize] > arr[i] {
                stack.pop();
            }
            if !stack.is_empty() {
                right[i] = *stack.last().unwrap();
            }
            stack.push(i as i32);
        }

        // Traverse the array, to find the sum
        for i in 0..n {
            ret +=
                ((((right[i] - (i as i32)) * ((i as i32) - left[i])) as i64) * (arr[i] as i64)) %
                MOD;
            ret %= MOD;
        }

        (ret % (MOD as i64)) as i32
    }
}
```

### **Go**

```go
func sumSubarrayMins(arr []int) int {
	mod := int(1e9) + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	for i := range left {
		left[i] = -1
		right[i] = n
	}
	stk := []int{}
	for i, v := range arr {
		for len(stk) > 0 && arr[stk[len(stk)-1]] >= v {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && arr[stk[len(stk)-1]] > arr[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	ans := 0
	for i, v := range arr {
		ans += (i - left[i]) * (right[i] - i) * v % mod
		ans %= mod
	}
	return ans
}
```

### **TypeScript**

```ts
function sumSubarrayMins(arr: number[]): number {
    const n = arr.length;
    function getEle(i: number): number {
        if (i == -1 || i == n) return Number.MIN_SAFE_INTEGER;
        return arr[i];
    }
    let ans = 0;
    const mod = 10 ** 9 + 7;
    let stack = [];
    for (let i = -1; i <= n; i++) {
        while (stack.length && getEle(stack[0]) > getEle(i)) {
            const idx = stack.shift();
            ans = (ans + arr[idx] * (idx - stack[0]) * (i - idx)) % mod;
        }
        stack.unshift(i);
    }
    return ans;
}
```

### **...**

```

```

<!-- tabs:end -->
