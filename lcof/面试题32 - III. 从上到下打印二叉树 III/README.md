# [面试题 32 - III. 从上到下打印二叉树 III](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/)

## 题目描述

<p>请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。</p>

<p>&nbsp;</p>

<p>例如:<br>
给定二叉树:&nbsp;<code>[3,9,20,null,null,15,7]</code>,</p>

<pre>    3
   / \
  9  20
    /  \
   15   7
</pre>

<p>返回其层次遍历结果：</p>

<pre>[
  [3],
  [20,9],
  [15,7]
]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>节点总数 &lt;= 1000</code></li>
</ol>

## 解法

<!-- tabs:start -->

### **Python3**

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if root is None:
            return []
        q = deque()
        res = []
        q.append(root)
        while q:
            size = len(q)
            t = []
            for _ in range(size):
                node = q.popleft()
                t.append(node.val)
                if node.left is not None:
                    q.append(node.left)
                if node.right is not None:
                    q.append(node.right)
            res.append(t if len(res) & 1 == 0 else t[::-1])
        return res
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
    public List<List<Integer>> levelOrder(TreeNode root) {
        if (root == null) return Collections.emptyList();
        Deque<TreeNode> q = new ArrayDeque<>();
        List<List<Integer>> res = new ArrayList<>();
        q.offer(root);
        while (!q.isEmpty()) {
            int size = q.size();
            List<Integer> t = new ArrayList<>();
            while (size-- > 0) {
                TreeNode node = q.poll();
                t.add(node.val);
                if (node.left != null) q.offer(node.left);
                if (node.right != null) q.offer(node.right);
            }
            if ((res.size() & 1) == 1) Collections.reverse(t);
            res.add(t);
        }
        return res;
    }
}
```

### **JavaScript**

```js
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrder = function (root) {
    if (!root) return [];
    let queue = [root];
    let res = [];
    let depth = 0;
    let dir = true;
    while (queue.length) {
        let len = queue.length;
        for (let i = 0; i < len; i++) {
            let node = queue.shift();
            if (!node) continue;
            if (!res[depth]) res[depth] = [];
            if (dir) {
                res[depth].push(node.val);
            } else {
                res[depth].unshift(node.val);
            }
            queue.push(node.left, node.right);
        }
        depth++;
        dir = !dir;
    }
    return res;
};
```

### **Go**

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    res := [][]int{}
    queue := []*TreeNode{}
    queue = append(queue,root)
    level := 0
    for len(queue) != 0 {
        size := len(queue)
        ans := []int{}
        //size记录每层大小,level记录层数
        for size > 0 {
            cur := queue[0]
            if level & 1 == 0 {
                ans = append(ans, cur.Val)
            } else {
                ans = append([]int{cur.Val},ans...)
            }

            queue = queue[1:]
            size--
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        level++
        res = append(res, ans)
    }
    return res
}
```

### **C++**

```cpp
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> ans;
        if (root == NULL) return ans;
        queue<TreeNode*> q;
        q.push(root);
        bool flag = false;
        while (!q.empty()) {
            int n = q.size();
            vector<int> v;
            for (int i = 0; i < n; ++i) {
                TreeNode* node = q.front();
                q.pop();
                v.emplace_back(node->val);
                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
            }
            if (flag) reverse(v.begin(), v.end());
            flag = !flag;
            ans.emplace_back(v);
        }
        return ans;
    }
};
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

function levelOrder(root: TreeNode | null): number[][] {
    const res = [];
    if (root == null) {
        return res;
    }
    let isEven = false;
    const levelFn = (nodes: TreeNode[]) => {
        if (nodes.length === 0) {
            return res;
        }
        const nextNodes = [];
        const values = nodes.map(({ val, left, right }) => {
            left && nextNodes.push(left);
            right && nextNodes.push(right);
            return val;
        });
        res.push(isEven ? values.reverse() : values);
        isEven = !isEven;
        return levelFn(nextNodes);
    };
    return levelFn([root]);
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
use std::collections::VecDeque;

impl Solution {
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        if root.is_none() {
            return res;
        }
        let mut nodes = VecDeque::new();
        nodes.push_back(root.unwrap());
        let mut is_even = false;
        while !nodes.is_empty() {
            let mut values = Vec::new();
            for _ in 0..nodes.len() {
                let node = nodes.pop_front().unwrap();
                let mut node = node.borrow_mut();
                values.push(node.val);
                if node.left.is_some() {
                    nodes.push_back(node.left.take().unwrap())
                }
                if node.right.is_some() {
                    nodes.push_back(node.right.take().unwrap())
                }
            }
            if is_even {
                values.reverse()
            }
            res.push(values);
            is_even = !is_even
        }
        res
    }
}
```

### **...**

```

```

<!-- tabs:end -->
