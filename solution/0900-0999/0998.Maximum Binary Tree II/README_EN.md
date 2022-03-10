# [998. Maximum Binary Tree II](https://leetcode.com/problems/maximum-binary-tree-ii)

[中文文档](/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/README.md)

## Description

<p>A <strong>maximum tree</strong> is a tree where every node has a value greater than any other value in its subtree.</p>

<p>You are given the <code>root</code> of a maximum binary tree and an integer <code>val</code>.</p>

<p>Just as in the <a href="https://leetcode.com/problems/maximum-binary-tree/" target="_blank">previous problem</a>, the given tree was constructed from a list <code>a</code> (<code>root = Construct(a)</code>) recursively with the following <code>Construct(a)</code> routine:</p>

<ul>
	<li>If <code>a</code> is empty, return <code>null</code>.</li>
	<li>Otherwise, let <code>a[i]</code> be the largest element of <code>a</code>. Create a <code>root</code> node with the value <code>a[i]</code>.</li>
	<li>The left child of <code>root</code> will be <code>Construct([a[0], a[1], ..., a[i - 1]])</code>.</li>
	<li>The right child of <code>root</code> will be <code>Construct([a[i + 1], a[i + 2], ..., a[a.length - 1]])</code>.</li>
	<li>Return <code>root</code>.</li>
</ul>

<p>Note that we were not given <code>a</code> directly, only a root node <code>root = Construct(a)</code>.</p>

<p>Suppose <code>b</code> is a copy of <code>a</code> with the value <code>val</code> appended to it. It is guaranteed that <code>b</code> has unique values.</p>

<p>Return <code>Construct(b)</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://cdn.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maxtree1.JPG" style="width: 376px; height: 235px;" />
<pre>
<strong>Input:</strong> root = [4,1,3,null,null,2], val = 5
<strong>Output:</strong> [5,4,null,1,3,null,null,2]
<strong>Explanation:</strong> a = [1,4,2,3], b = [1,4,2,3,5]
</pre>

<p><strong>Example 2:</strong></p>
<img alt="" src="https://cdn.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maxtree21.JPG" style="width: 358px; height: 156px;" />
<pre>
<strong>Input:</strong> root = [5,2,4,null,1], val = 3
<strong>Output:</strong> [5,2,4,null,1,null,3]
<strong>Explanation:</strong> a = [2,1,5,4], b = [2,1,5,4,3]
</pre>

<p><strong>Example 3:</strong></p>
<img alt="" src="https://cdn.jsdelivr.net/gh/doocs/leetcode@main/solution/0900-0999/0998.Maximum%20Binary%20Tree%20II/images/maxtree3.JPG" style="width: 404px; height: 180px;" />
<pre>
<strong>Input:</strong> root = [5,2,3,null,1], val = 4
<strong>Output:</strong> [5,2,4,null,1,3]
<strong>Explanation:</strong> a = [2,1,5,3], b = [2,1,5,3,4]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the tree is in the range <code>[1, 100]</code>.</li>
	<li><code>1 &lt;= Node.val &lt;= 100</code></li>
	<li>All the values of the tree are <strong>unique</strong>.</li>
	<li><code>1 &lt;= val &lt;= 100</code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

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
