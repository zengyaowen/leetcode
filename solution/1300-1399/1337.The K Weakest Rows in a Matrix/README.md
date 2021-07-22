# [1337. 矩阵中战斗力最弱的 K 行](https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix)

[English Version](/solution/1300-1399/1337.The%20K%20Weakest%20Rows%20in%20a%20Matrix/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给你一个大小为 <code>m * n</code> 的矩阵 <code>mat</code>，矩阵由若干军人和平民组成，分别用 1 和 0 表示。</p>

<p>请你返回矩阵中战斗力最弱的 <code>k</code> 行的索引，按从最弱到最强排序。</p>

<p>如果第 <em><strong>i</strong></em> 行的军人数量少于第 <em><strong>j</strong></em> 行，或者两行军人数量相同但<em><strong> i</strong></em> 小于 <em><strong>j</strong></em>，那么我们认为第<em><strong> i </strong></em>行的战斗力比第<em><strong> j </strong></em>行弱。</p>

<p>军人 <strong>总是</strong> 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>mat = 
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]], 
k = 3
<strong>输出：</strong>[2,0,3]
<strong>解释：</strong>
每行中的军人数目：
行 0 -> 2 
行 1 -> 4 
行 2 -> 1 
行 3 -> 2 
行 4 -> 5 
从最弱到最强对这些行排序后得到 [2,0,3,1,4]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>mat = 
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]], 
k = 2
<strong>输出：</strong>[0,2]
<strong>解释：</strong> 
每行中的军人数目：
行 0 -> 1 
行 1 -> 4 
行 2 -> 1 
行 3 -> 1 
从最弱到最强对这些行排序后得到 [0,2,3,1]
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == mat.length</code></li>
	<li><code>n == mat[i].length</code></li>
	<li><code>2 <= n, m <= 100</code></li>
	<li><code>1 <= k <= m</code></li>
	<li><code>matrix[i][j]</code> 不是 0 就是 1</li>
</ul>


## 解法

<!-- 这里可写通用的实现逻辑 -->

二分查找 + 排序。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        m, n = len(mat), len(mat[0])
        res = []
        for row in mat:
            left, right = 0, n
            while left < right:
                mid = (left + right) >> 1
                if row[mid] == 0:
                    right = mid
                else:
                    left = mid + 1
            res.append(left)
        idx = list(range(m))
        idx.sort(key=lambda x: res[x])
        return idx[:k]
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public int[] kWeakestRows(int[][] mat, int k) {
        int m = mat.length, n = mat[0].length;
        int[] res = new int[m];
        List<Integer> idx = new ArrayList<>();
        for (int i = 0; i < m; ++i) {
            idx.add(i);
            int[] row = mat[i];
            int left = 0, right = n;
            while (left < right) {
                int mid = (left + right) >> 1;
                if (row[mid] == 0) {
                    right = mid;
                } else {
                    left = mid + 1;
                }
            }
            res[i] = left;
        }
        idx.sort(Comparator.comparingInt(a -> res[a]));
        int[] ans = new int[k];
        for (int i = 0; i < k; ++i) {
            ans[i] = idx.get(i);
        }
        return ans;
    }
}
```

### **TypeScript**

```ts
function kWeakestRows(mat: number[][], k: number): number[] {
    let n = mat.length;
    let sumMap = mat.map((d, i)=> ([d.reduce((a, c) => a + c, 0), i]));
    let ans = [];
    // 冒泡排序
    for (let i = 0; i < k; i++) {
        for (let j = i; j < n; j++) {
            if (sumMap[j][0] < sumMap[i][0] || (sumMap[j][0] == sumMap[i][0]) && sumMap[i][1] > sumMap[j][1]) {
                [sumMap[i], sumMap[j]] = [sumMap[j], sumMap[i]];
            }
        }
        ans.push(sumMap[i][1]);
    }
    return ans;
};
```

### **...**

```

```

<!-- tabs:end -->
