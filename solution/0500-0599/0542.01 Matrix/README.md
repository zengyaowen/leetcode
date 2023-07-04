# [542. 01 矩阵](https://leetcode.cn/problems/01-matrix)

[English Version](/solution/0500-0599/0542.01%20Matrix/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给定一个由 <code>0</code> 和 <code>1</code> 组成的矩阵 <code>mat</code> ，请输出一个大小相同的矩阵，其中每一个格子是 <code>mat</code> 中对应位置元素到最近的 <code>0</code> 的距离。</p>

<p>两个相邻元素间的距离为 <code>1</code> 。</p>

<p> </p>

<p><b>示例 1：</b></p>

<p><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0500-0599/0542.01%20Matrix/images/1626667201-NCWmuP-image.png" style="width: 150px; " /></p>

<pre>
<strong>输入：</strong>mat =<strong> </strong>[[0,0,0],[0,1,0],[0,0,0]]
<strong>输出：</strong>[[0,0,0],[0,1,0],[0,0,0]]
</pre>

<p><b>示例 2：</b></p>

<p><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0500-0599/0542.01%20Matrix/images/1626667205-xFxIeK-image.png" style="width: 150px; " /></p>

<pre>
<b>输入：</b>mat =<b> </b>[[0,0,0],[0,1,0],[1,1,1]]
<strong>输出：</strong>[[0,0,0],[0,1,0],[1,2,1]]
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == mat.length</code></li>
	<li><code>n == mat[i].length</code></li>
	<li><code>1 <= m, n <= 10<sup>4</sup></code></li>
	<li><code>1 <= m * n <= 10<sup>4</sup></code></li>
	<li><code>mat[i][j] is either 0 or 1.</code></li>
	<li><code>mat</code> 中至少有一个 <code>0 </code></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

**方法一：多源 BFS**

初始化结果矩阵 ans，所有 0 的距离为 0，所以 1 的距离为 -1。初始化队列 q 存储 BFS 需要检查的位置，并将所有 0 的位置入队。

循环弹出队列 q 的元素 `p(i, j)`，检查邻居四个点。对于邻居 `(x, y)`，如果 `ans[x][y] = -1`，则更新 `ans[x][y] = ans[i][j] + 1`。同时将 `(x, y)` 入队。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m, n = len(mat), len(mat[0])
        ans = [[-1] * n for _ in range(m)]
        q = deque()
        for i, row in enumerate(mat):
            for j, v in enumerate(row):
                if v == 0:
                    ans[i][j] = 0
                    q.append((i, j))
        dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
        while q:
            i, j = q.popleft()
            for a, b in dirs:
                x, y = i + a, j + b
                if 0 <= x < m and 0 <= y < n and ans[x][y] == -1:
                    ans[x][y] = ans[i][j] + 1
                    q.append((x, y))
        return ans
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {

    public int[][] updateMatrix(int[][] mat) {
        int m = mat.length, n = mat[0].length;
        int[][] ans = new int[m][n];
        for (int i = 0; i < m; ++i) {
            Arrays.fill(ans[i], -1);
        }
        Deque<int[]> q = new LinkedList<>();
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (mat[i][j] == 0) {
                    ans[i][j] = 0;
                    q.offer(new int[] {i, j});
                }
            }
        }
        int[] dirs = new int[] {-1, 0, 1, 0, -1};
        while (!q.isEmpty()) {
            int[] t = q.poll();
            for (int i = 0; i < 4; ++i) {
                int x = t[0] + dirs[i];
                int y = t[1] + dirs[i + 1];
                if (x >= 0 && x < m && y >= 0 && y < n && ans[x][y] == -1) {
                    ans[x][y] = ans[t[0]][t[1]] + 1;
                    q.offer(new int[] {x, y});
                }
            }
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<vector<int>> updateMatrix(vector<vector<int>>& mat) {
        int m = mat.size(), n = mat[0].size();
        vector<vector<int>> ans(m, vector<int>(n, -1));
        queue<pair<int, int>> q;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (mat[i][j] == 0) {
                    ans[i][j] = 0;
                    q.emplace(i, j);
                }
            }
        }
        vector<int> dirs = {-1, 0, 1, 0, -1};
        while (!q.empty()) {
            auto p = q.front();
            q.pop();
            for (int i = 0; i < 4; ++i) {
                int x = p.first + dirs[i];
                int y = p.second + dirs[i + 1];
                if (x >= 0 && x < m && y >= 0 && y < n && ans[x][y] == -1) {
                    ans[x][y] = ans[p.first][p.second] + 1;
                    q.emplace(x, y);
                }
            }
        }
        return ans;
    }
};
```

### **Rust**

```rust
use std::collections::VecDeque;

impl Solution {
    #[allow(dead_code)]
    pub fn update_matrix(mat: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let n: usize = mat.len();
        let m: usize = mat[0].len();
        let mut ret_vec: Vec<Vec<i32>> = vec![vec![-1; m]; n];
        // The inner tuple is of <X, Y, Current Count>
        let mut the_q: VecDeque<(usize, usize)> = VecDeque::new();
        let traverse_vec: Vec<(i32, i32)> = vec![(-1, 0), (1, 0), (0, 1), (0, -1)];

        // Initialize the queue
        for i in 0..n {
            for j in 0..m {
                if mat[i][j] == 0 {
                    // For the zero cell, enqueue at first
                    the_q.push_back((i, j));
                    // Set to 0 in return vector
                    ret_vec[i][j] = 0;
                }
            }
        }

        while !the_q.is_empty() {
            let (x, y) = the_q.front().unwrap().clone();
            the_q.pop_front();
            for pair in &traverse_vec {
                let cur_x = pair.0 + x as i32;
                let cur_y = pair.1 + y as i32;
                if Solution::check_bounds(cur_x, cur_y, n as i32, m as i32) && ret_vec[cur_x as usize][cur_y as usize] == -1 {
                    // The current cell has not be updated yet, and is also in bound
                    ret_vec[cur_x as usize][cur_y as usize] = ret_vec[x][y] + 1;
                    the_q.push_back((cur_x as usize, cur_y as usize));
                }
            }
        }

        ret_vec
    }

    #[allow(dead_code)]
    pub fn check_bounds(i: i32, j: i32, n: i32, m: i32) -> bool {
        i >= 0 && i < n && j >= 0 && j < m
    }
}
```

### **Go**

```go
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	type pair struct{ x, y int }
	var q []pair
	for i, row := range mat {
		for j, v := range row {
			if v == 0 {
				ans[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}
	dirs := []int{-1, 0, 1, 0, -1}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			x, y := p.x+dirs[i], p.y+dirs[i+1]
			if x >= 0 && x < m && y >= 0 && y < n && ans[x][y] == -1 {
				ans[x][y] = ans[p.x][p.y] + 1
				q = append(q, pair{x, y})
			}
		}
	}
	return ans
}
```

### **...**

```

```

<!-- tabs:end -->
