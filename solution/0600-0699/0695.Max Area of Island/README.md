# [695. 岛屿的最大面积](https://leetcode-cn.com/problems/max-area-of-island)

[English Version](/solution/0600-0699/0695.Max%20Area%20of%20Island/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给定一个包含了一些 <code>0</code> 和 <code>1</code> 的非空二维数组&nbsp;<code>grid</code> 。</p>

<p>一个&nbsp;<strong>岛屿</strong>&nbsp;是由一些相邻的&nbsp;<code>1</code>&nbsp;(代表土地) 构成的组合，这里的「相邻」要求两个 <code>1</code> 必须在水平或者竖直方向上相邻。你可以假设&nbsp;<code>grid</code> 的四个边缘都被 <code>0</code>（代表水）包围着。</p>

<p>找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 <code>0</code> 。)</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,<strong>1</strong>,0,<strong>1</strong>,0,0],
 [0,1,0,0,1,1,0,0,<strong>1</strong>,<strong>1</strong>,<strong>1</strong>,0,0],
 [0,0,0,0,0,0,0,0,0,0,<strong>1</strong>,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
</pre>

<p>对于上面这个给定矩阵应返回&nbsp;<code>6</code>。注意答案不应该是 <code>11</code> ，因为岛屿只能包含水平或垂直的四个方向的 <code>1</code> 。</p>

<p><strong>示例 2:</strong></p>

<pre>[[0,0,0,0,0,0,0,0]]</pre>

<p>对于上面这个给定的矩阵, 返回&nbsp;<code>0</code>。</p>

<p>&nbsp;</p>

<p><strong>注意:&nbsp;</strong>给定的矩阵<code>grid</code>&nbsp;的长度和宽度都不超过 50。</p>

## 解法

<!-- 这里可写通用的实现逻辑 -->

DFS 或并查集实现。

并查集模板：

模板 1——朴素并查集：

```python
# 初始化，p存储每个点的父节点
p = list(range(n))

# 返回x的祖宗节点
def find(x):
    if p[x] != x:
        # 路径压缩
        p[x] = find(p[x])
    return p[x]

# 合并a和b所在的两个集合
p[find(a)] = find(b)
```

模板 2——维护 size 的并查集：

```python
# 初始化，p存储每个点的父节点，size只有当节点是祖宗节点时才有意义，表示祖宗节点所在集合中，点的数量
p = list(range(n))
size = [1] * n

# 返回x的祖宗节点
def find(x):
    if p[x] != x:
        # 路径压缩
        p[x] = find(p[x])
    return p[x]

# 合并a和b所在的两个集合
if find(a) != find(b):
    size[find(b)] += size[find(a)]
    p[find(a)] = find(b)
```

模板 3——维护到祖宗节点距离的并查集：

```python
# 初始化，p存储每个点的父节点，d[x]存储x到p[x]的距离
p = list(range(n))
d = [0] * n

# 返回x的祖宗节点
def find(x):
    if p[x] != x:
        t = find(p[x])
        d[x] += d[p[x]]
        p[x] = t
    return p[x]

# 合并a和b所在的两个集合
p[find(a)] = find(b)
d[find(a)] = distance
```

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

DFS：

```python
class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        def dfs(i, j):
            if i < 0 or i >= len(grid) or j < 0 or j >= len(grid[0]) or grid[i][j] == 0:
                return 0
            grid[i][j] = 0
            res = 1
            for x, y in [[0, 1], [0, -1], [1, 0], [-1, 0]]:
                res += dfs(i + x, j + y)
            return res
        
        m, n = len(grid), len(grid[0])
        res = 0
        for i in range(m):
            for j in range(n):
                t = dfs(i, j)
                res = max(res, t)
        return res
```

并查集：

```python
class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        p = list(range(m * n))
        size = [1] * (m * n)

        def find(x):
            if p[x] != x:
                p[x] = find(p[x])
            return p[x]

        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    if i < m - 1 and grid[i + 1][j] == 1:
                        a, b = find(i * n + j), find((i + 1) * n + j)
                        if a != b:
                            size[a] += size[b]
                        p[b] = a
                    if j < n - 1 and grid[i][j + 1] == 1:
                        a, b = find(i * n + j), find(i * n + j + 1)
                        if a != b:
                            size[a] += size[b]
                        p[b] = a

        res = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    res = max(res, size[i * n + j])
        return res
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

DFS：

```java
class Solution {
    private int[][] directions = {{0, 1}, {0, - 1}, {1, 0}, {-1, 0}};

    public int maxAreaOfIsland(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int res = 0;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                int t = dfs(grid, i, j, m, n);
                res = Math.max(res, t);
            }
        }
        return res;
    }

    private int dfs(int[][] grid, int i, int j, int m, int n) {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0) {
            return 0;
        }
        grid[i][j] = 0;
        int res = 1;
        for (int[] direction : directions) {
            res += dfs(grid, i + direction[0], j + direction[1], m, n);
        }
        return res;
    }
}
```

并查集：

```java
class Solution {
    private int[] p;
    private int[] size;

    public int maxAreaOfIsland(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        p = new int[m * n];
        size = new int[m * n];
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                p[i * n + j] = i * n + j;
                size[i * n + j] = 1;
            }
        }
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1) {
                    if (i < m - 1 && grid[i + 1][j] == 1) {
                        int a = find(i * n + j), b = find((i + 1) * n + j);
                        if (a != b) {
                            size[a] += size[b];
                        }
                        p[b] = a;
                    }
                    if (j < n - 1 && grid[i][j + 1] == 1) {
                        int a = find(i * n + j), b = find(i * n + j + 1);
                        if (a != b) {
                            size[a] += size[b];
                        }
                        p[b] = a;
                    }
                }
            }
        }
        int res = 0;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1) {
                    res = Math.max(res, size[i * n + j]);
                }
            }
        }
        return res;
    }

    private int find(int x) {
        if (p[x] != x) {
            p[x] = find(p[x]);
        }
        return p[x];
    }
}
```

### **TypeScript**

DFS：

```ts
function maxAreaOfIsland(grid: number[][]): number {
    let m = grid.length, n = grid[0].length;
    let res = 0;
    for (let i = 0; i < m; ++i) {
        for (let j = 0; j < n; ++j) {
            if (grid[i][j] == 1) {
                res = Math.max(dfs(grid, i, j), res);
            }
        }
    }
    return res;
};

function dfs(grid: number[][], i: number, j: number): number {
    let m = grid.length, n = grid[0].length;
    if (i < 0 || i > m - 1 || j < 0 || j > n - 1 || grid[i][j] == 0) {
        return 0;
    }
    grid[i][j] = 0;
    let res = 1;
    for (let [dx, dy] of [[0, 1], [0, -1], [1, 0], [-1, 0]]) {
        res += dfs(grid, i + dx, j + dy);
    }
    return res;
}
```

### **C++**

DFS：

```cpp
class Solution {
public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        int res = 0;
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                int t = dfs(grid, i, j, m, n);
                res = max(res, t);
            }
        }
        return res;
    }
private:
    vector<vector<int>> dirs = {{0, 1}, {0, - 1}, {1, 0}, {-1, 0}};

    int dfs(vector<vector<int>>& grid, int i, int j, int m, int n) {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0) return 0;
        grid[i][j] = 0;
        int res = 1;
        for (auto dir : dirs)
            res += dfs(grid, i + dir[0], j + dir[1], m, n);
        return res;
    }
};
```

并查集：

```cpp
class Solution {
public:
    vector<int> p;

    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<int> size(m * n, 1);
        p.resize(m * n);
        for (int i = 0; i < p.size(); ++i) p[i] = i;
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (grid[i][j] == 1)
                {
                    if (i < m - 1 && grid[i + 1][j] == 1)
                    {
                        int a = find(i * n + j), b = find((i + 1) * n + j);
                        if (a != b) size[a] += size[b];
                        p[b] = a;
                    }
                    if (j < n - 1 && grid[i][j + 1] == 1)
                    {
                        int a = find(i * n + j), b = find(i * n + j + 1);
                        if (a != b) size[a] += size[b];
                        p[b] = a;
                    }
                }
            }
        }
        int res = 0;
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (grid[i][j] == 1) res = max(res, size[i * n + j]);
            }
        }
        return res;
    }

    int find(int x) {
        if (p[x] != x) p[x] = find(p[x]);
        return p[x];
    }
};
```

### **Go**

并查集：

```go
var p []int

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	p = make([]int, m*n)
	size := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p[i*n+j] = i*n + j
			size[i*n+j] = 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i < m-1 && grid[i+1][j] == 1 {
					a, b := find(i*n+j), find((i+1)*n+j)
					if a != b {
						size[a] += size[b]
					}
					p[b] = a
				}
				if j < n-1 && grid[i][j+1] == 1 {
					a, b := find(i*n+j), find(i*n+j+1)
					if a != b {
						size[a] += size[b]
					}
					p[b] = a
				}
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, size[i*n+j])
			}
		}
	}
	return res
}

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### **...**

```

```

<!-- tabs:end -->
