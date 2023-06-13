# [78. Subsets](https://leetcode.com/problems/subsets)

[中文文档](/solution/0000-0099/0078.Subsets/README.md)

## Description

<p>Given an integer array <code>nums</code> of <strong>unique</strong> elements, return <em>all possible</em> <span data-keyword="subset"><em>subsets</em></span> <em>(the power set)</em>.</p>

<p>The solution set <strong>must not</strong> contain duplicate subsets. Return the solution in <strong>any order</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,3]
<strong>Output:</strong> [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0]
<strong>Output:</strong> [[],[0]]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10</code></li>
	<li><code>-10 &lt;= nums[i] &lt;= 10</code></li>
	<li>All the numbers of&nbsp;<code>nums</code> are <strong>unique</strong>.</li>
</ul>

## Solutions

DFS.

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        def dfs(i: int):
            if i == len(nums):
                ans.append(t[:])
                return
            dfs(i + 1)
            t.append(nums[i])
            dfs(i + 1)
            t.pop()

        ans = []
        t = []
        dfs(0)
        return ans
```

```python
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = []
        for mask in range(1 << len(nums)):
            t = [x for i, x in enumerate(nums) if mask >> i & 1]
            ans.append(t)
        return ans
```

### **Java**

```java
class Solution {
    private List<List<Integer>> ans = new ArrayList<>();
    private List<Integer> t = new ArrayList<>();
    private int[] nums;

    public List<List<Integer>> subsets(int[] nums) {
        this.nums = nums;
        dfs(0);
        return ans;
    }

    private void dfs(int i) {
        if (i == nums.length) {
            ans.add(new ArrayList<>(t));
            return;
        }
        dfs(i + 1);
        t.add(nums[i]);
        dfs(i + 1);
        t.remove(t.size() - 1);
    }
}
```

```java
class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        int n = nums.length;
        List<List<Integer>> ans = new ArrayList<>();
        for (int mask = 0; mask < 1 << n; ++mask) {
            List<Integer> t = new ArrayList<>();
            for (int i = 0; i < n; ++i) {
                if (((mask >> i) & 1) == 1) {
                    t.add(nums[i]);
                }
            }
            ans.add(t);
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> ans;
        vector<int> t;
        function<void(int)> dfs = [&](int i) -> void {
            if (i == nums.size()) {
                ans.push_back(t);
                return;
            }
            dfs(i + 1);
            t.push_back(nums[i]);
            dfs(i + 1);
            t.pop_back();
        };
        dfs(0);
        return ans;
    }
};
```

```cpp
class Solution {
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        int n = nums.size();
        vector<vector<int>> ans;
        for (int mask = 0; mask < 1 << n; ++mask) {
            vector<int> t;
            for (int i = 0; i < n; ++i) {
                if (mask >> i & 1) {
                    t.emplace_back(nums[i]);
                }
            }
            ans.emplace_back(t);
        }
        return ans;
    }
};
```

### **Go**

```go
func subsets(nums []int) (ans [][]int) {
	t := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(i + 1)
		t = append(t, nums[i])
		dfs(i + 1)
		t = t[:len(t)-1]
	}
	dfs(0)
	return
}
```

```go
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, x := range nums {
			if mask>>i&1 == 1 {
				t = append(t, x)
			}
		}
		ans = append(ans, t)
	}
	return
}
```

### **TypeScript**

```ts
function subsets(nums: number[]): number[][] {
    const ans: number[][] = [];
    const t: number[] = [];
    const dfs = (i: number) => {
        if (i === nums.length) {
            ans.push(t.slice());
            return;
        }
        dfs(i + 1);
        t.push(nums[i]);
        dfs(i + 1);
        t.pop();
    };
    dfs(0);
    return ans;
}
```

```ts
function subsets(nums: number[]): number[][] {
    const n = nums.length;
    const ans: number[][] = [];
    for (let mask = 0; mask < 1 << n; ++mask) {
        const t: number[] = [];
        for (let i = 0; i < n; ++i) {
            if (((mask >> i) & 1) === 1) {
                t.push(nums[i]);
            }
        }
        ans.push(t);
    }
    return ans;
}
```

### **Rust**

```rust
impl Solution {
    fn dfs(i: usize, path: &mut Vec<i32>, res: &mut Vec<Vec<i32>>, nums: &Vec<i32>) {
        if i == nums.len() {
            res.push(path.clone());
            return;
        }
        Self::dfs(i + 1, path, res, nums);
        path.push(nums[i]);
        Self::dfs(i + 1, path, res, nums);
        path.pop();
    }

    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        Self::dfs(0, &mut Vec::new(), &mut res, &nums);
        res
    }
}
```

### **...**

```

```

<!-- tabs:end -->
