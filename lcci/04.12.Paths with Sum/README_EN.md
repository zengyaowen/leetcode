# [04.12. Paths with Sum](https://leetcode.cn/problems/paths-with-sum-lcci)

[中文文档](/lcci/04.12.Paths%20with%20Sum/README.md)

## Description

<p>You are given a binary tree in which each node contains an integer value (which might be positive or negative). Design an algorithm to count the number of paths that sum to a given value. The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).</p>

<p><strong>Example:</strong><br />

Given the following tree and &nbsp;<code>sum = 22,</code></p>

<pre>

              5

             / \

            4   8

           /   / \

          11  13  4

         /  \    / \

        7    2  5   1

</pre>

<p>Output:</p>

<pre>

3

<strong>Explanation: </strong>Paths that have sum 22 are: [5,4,11,2], [5,8,4,5], [4,11,7]</pre>

<p>Note:</p>

<ul>
	<li><code>node number &lt;= 10000</code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> int:
        def dfs(root, sum, flag):
            nonlocal ans
            if not root:
                return 0
            if sum - root.val == 0:
                ans += 1
            if flag == 0:
                dfs(root.left, sum, 0)
                dfs(root.right, sum, 0)
            dfs(root.left, sum - root.val, 1)
            dfs(root.right, sum - root.val, 1)

        if not root:
            return 0
        ans = 0
        dfs(root, sum, 0)
        return ans
```

### **Java**

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    private Map<Long, Integer> cnt = new HashMap<>();
    private int target;

    public int pathSum(TreeNode root, int sum) {
        cnt.put(0L, 1);
        target = sum;
        return dfs(root, 0);
    }

    private int dfs(TreeNode root, long s) {
        if (root == null) {
            return 0;
        }
        s += root.val;
        int ans = cnt.getOrDefault(s - target, 0);
        cnt.merge(s, 1, Integer::sum);
        ans += dfs(root.left, s);
        ans += dfs(root.right, s);
        cnt.merge(s, -1, Integer::sum);
        return ans;
    }
}
```

### **C++**

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int pathSum(TreeNode* root, int sum) {
        unordered_map<long long, int> cnt;
        cnt[0] = 1;
        function<int(TreeNode*, long long)> dfs = [&](TreeNode* root, long long s) {
            if (!root) {
                return 0;
            }
            s += root->val;
            int ans = cnt[s - sum];
            ++cnt[s];
            ans += dfs(root->left, s);
            ans += dfs(root->right, s);
            --cnt[s];
            return ans;
        };
        return dfs(root, 0);
    }
};
```

### **Go**

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, s int) int {
		if root == nil {
			return 0
		}
		s += root.Val
		ans := cnt[s-sum]
		cnt[s]++
		ans += dfs(root.Left, s)
		ans += dfs(root.Right, s)
		cnt[s]--
		return ans
	}
	return dfs(root, 0)
}
```

### **TypeScript**

```ts
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function pathSum(root: TreeNode | null, sum: number): number {
    const cnt: Map<number, number> = new Map();
    cnt.set(0, 1);
    const dfs = (root: TreeNode | null, s: number): number => {
        if (!root) {
            return 0;
        }
        s += root.val;
        let ans = cnt.get(s - sum) ?? 0;
        cnt.set(s, (cnt.get(s) ?? 0) + 1);
        ans += dfs(root.left, s);
        ans += dfs(root.right, s);
        cnt.set(s, (cnt.get(s) ?? 0) - 1);
        return ans;
    };
    return dfs(root, 0);
}
```

### **Rust**

```rust
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
use std::collections::HashMap;
impl Solution {
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> i32 {
        let mut cnt = HashMap::new();
        cnt.insert(0, 1);
        return Self::dfs(root, sum, 0, &mut cnt);
    }

    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, sum: i32, s: i32, cnt: &mut HashMap<i32, i32>) -> i32 {
        if let Some(node) = root {
            let node = node.borrow();
            let s = s + node.val;
            let mut ans = *cnt.get(&(s - sum)).unwrap_or(&0);
            *cnt.entry(s).or_insert(0) += 1;
            ans += Self::dfs(node.left.clone(), sum, s, cnt);
            ans += Self::dfs(node.right.clone(), sum, s, cnt);
            *cnt.entry(s).or_insert(0) -= 1;
            return ans;
        }
        return 0;
    }
}
```

### **...**

```

```

<!-- tabs:end -->
