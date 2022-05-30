# [998. 最大二叉树 II](https://leetcode.cn/problems/maximum-binary-tree-ii)

[English Version](/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p><strong>最大树</strong> 定义：一棵树，并满足：其中每个节点的值都大于其子树中的任何其他值。</p>

<p>给你最大树的根节点 <code>root</code> 和一个整数 <code>val</code> 。</p>

<p>就像 <a href="https://leetcode.cn/problems/maximum-binary-tree/" target="_blank">之前的问题</a> 那样，给定的树是利用 <code>Construct(a)</code>&nbsp;例程从列表&nbsp;<code>a</code>（<code>root = Construct(a)</code>）递归地构建的：</p>

<ul>
	<li>如果 <code>a</code> 为空，返回&nbsp;<code>null</code> 。</li>
	<li>否则，令&nbsp;<code>a[i]</code> 作为 <code>a</code> 的最大元素。创建一个值为&nbsp;<code>a[i]</code>&nbsp;的根节点 <code>root</code> 。</li>
	<li><code>root</code>&nbsp;的左子树将被构建为&nbsp;<code>Construct([a[0], a[1], ..., a[i - 1]])</code> 。</li>
	<li><code>root</code>&nbsp;的右子树将被构建为&nbsp;<code>Construct([a[i + 1], a[i + 2], ..., a[a.length - 1]])</code> 。</li>
	<li>返回&nbsp;<code>root</code> 。</li>
</ul>

<p>请注意，题目没有直接给出 <code>a</code> ，只是给出一个根节点&nbsp;<code>root = Construct(a)</code> 。</p>

<p>假设 <code>b</code> 是 <code>a</code> 的副本，并在末尾附加值 <code>val</code>。题目数据保证 <code>b</code> 中的值互不相同。</p>

<p>返回&nbsp;<code>Construct(b)</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><strong><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maximum-binary-tree-1-1.png" style="height: 160px; width: 159px;" /><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maximum-binary-tree-1-2.png" style="height: 160px; width: 169px;" /></strong></p>

<pre>
<strong>输入：</strong>root = [4,1,3,null,null,2], val = 5
<strong>输出：</strong>[5,4,null,1,3,null,null,2]
<strong>解释：</strong>a = [1,4,2,3], b = [1,4,2,3,5]</pre>

<p><strong>示例 2：<br />
<img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maximum-binary-tree-2-1.png" style="height: 160px; width: 180px;" /><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maximum-binary-tree-2-2.png" style="height: 160px; width: 214px;" /></strong></p>

<pre>
<strong>输入：</strong>root = [5,2,4,null,1], val = 3
<strong>输出：</strong>[5,2,4,null,1,null,3]
<strong>解释：</strong>a = [2,1,5,4], b = [2,1,5,4,3]</pre>

<p><strong>示例 3：<br />
<img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maximum-binary-tree-3-1.png" style="height: 160px; width: 180px;" /><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maximum-binary-tree-3-2.png" style="height: 160px; width: 201px;" /></strong></p>

<pre>
<strong>输入：</strong>root = [5,2,3,null,1], val = 4
<strong>输出：</strong>[5,2,4,null,1,3]
<strong>解释：</strong>a = [2,1,5,3], b = [2,1,5,3,4]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点数目在范围 <code>[1, 100]</code> 内</li>
	<li><code>1 &lt;= Node.val &lt;= 100</code></li>
	<li>树中的所有值 <strong>互不相同</strong></li>
	<li><code>1 &lt;= val &lt;= 100</code></li>
</ul>

<p>&nbsp;</p>

## 解法

<!-- 这里可写通用的实现逻辑 -->

已知最大树 A，插入一个值 val 后，返回插入后的树。

如果 val 是最大数，那么将 val 作为新的根节点，root 作为新的根节点的左子树。

如果 val 不是最大数，由于 val 是在最后追加的数，那么一定是在 root 的右边，所以将 val 作为新节点插入 root 的右子树即可。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def insertIntoMaxTree(self, root: TreeNode, val: int) -> TreeNode:
        if root is None or root.val < val:
            return TreeNode(val, root, None)
        root.right = self.insertIntoMaxTree(root.right, val)
        return root
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode insertIntoMaxTree(TreeNode root, int val) {
        if (root == null || root.val < val) {
            return new TreeNode(val, root, null);
        }
        root.right = insertIntoMaxTree(root.right, val);
        return root;
    }
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

function insertIntoMaxTree(
    root: TreeNode | null,
    val: number,
): TreeNode | null {
    if (root == null || val > root.val) {
        return new TreeNode(val, root);
    }
    root.right = insertIntoMaxTree(root.right, val);
    return root;
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
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* insertIntoMaxTree(TreeNode* root, int val) {
        if (root == nullptr || root->val < val) {
            return new TreeNode(val, root, nullptr);
        }
        root->right = insertIntoMaxTree(root->right, val);
        return root;
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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
```

### **...**

```

```

<!-- tabs:end -->
