# [2641. 二叉树的堂兄弟节点 II](https://leetcode.cn/problems/cousins-in-binary-tree-ii)

[English Version](/solution/2600-2699/2641.Cousins%20in%20Binary%20Tree%20II/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给你一棵二叉树的根&nbsp;<code>root</code>&nbsp;，请你将每个节点的值替换成该节点的所有 <strong>堂兄弟节点值的和&nbsp;</strong>。</p>

<p>如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 <strong>堂兄弟</strong>&nbsp;。</p>

<p>请你返回修改值之后，树的根<em>&nbsp;</em><code>root</code><em>&nbsp;</em>。</p>

<p><strong>注意</strong>，一个节点的深度指的是从树根节点到这个节点经过的边数。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/2600-2699/2641.Cousins%20in%20Binary%20Tree%20II/images/example11.png" style="width: 571px; height: 151px;" /></p>

<pre>
<b>输入：</b>root = [5,4,9,1,10,null,7]
<b>输出：</b>[0,0,0,7,7,null,11]
<b>解释：</b>上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
- 值为 5 的节点没有堂兄弟，所以值修改为 0 。
- 值为 4 的节点没有堂兄弟，所以值修改为 0 。
- 值为 9 的节点没有堂兄弟，所以值修改为 0 。
- 值为 1 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
- 值为 10 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
- 值为 7 的节点有两个堂兄弟，值分别为 1 和 10 ，所以值修改为 11 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/2600-2699/2641.Cousins%20in%20Binary%20Tree%20II/images/diagram33.png" style="width: 481px; height: 91px;" /></p>

<pre>
<b>输入：</b>root = [3,1,2]
<b>输出：</b>[0,0,0]
<b>解释：</b>上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
- 值为 3 的节点没有堂兄弟，所以值修改为 0 。
- 值为 1 的节点没有堂兄弟，所以值修改为 0 。
- 值为 2 的节点没有堂兄弟，所以值修改为 0 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点数目的范围是&nbsp;<code>[1, 10<sup>5</sup>]</code> 。</li>
	<li><code>1 &lt;= Node.val &lt;= 10<sup>4</sup></code></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

**方法一：两次 DFS**

我们用一个数组 $s$ 记录二叉树每一层的节点值之和，其中 $s[i]$ 表示第 $i$ 层的节点值之和。

接下来，我们先跑一遍 DFS，计算出数组 $s$ 的值。然后再跑一遍 DFS，更新每个节点的子节点的值，子节点的值等于子节点所在层的节点值之和减去子节点及其兄弟节点的值。

时间复杂度 $O(n)$，空间复杂度 $O(n)$。其中 $n$ 是二叉树的节点数。

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
    def replaceValueInTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def dfs1(root, d):
            if root is None:
                return
            if len(s) <= d:
                s.append(0)
            s[d] += root.val
            dfs1(root.left, d + 1)
            dfs1(root.right, d + 1)

        def dfs2(root, d):
            if root is None:
                return
            t = 0
            if root.left:
                t += root.left.val
            if root.right:
                t += root.right.val
            if root.left:
                root.left.val = s[d] - t
            if root.right:
                root.right.val = s[d] - t
            dfs2(root.left, d + 1)
            dfs2(root.right, d + 1)

        s = []
        dfs1(root, 0)
        root.val = 0
        dfs2(root, 1)
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
    private List<Integer> s = new ArrayList<>();

    public TreeNode replaceValueInTree(TreeNode root) {
        dfs1(root, 0);
        root.val = 0;
        dfs2(root, 1);
        return root;
    }

    private void dfs1(TreeNode root, int d) {
        if (root == null) {
            return;
        }
        if (s.size() <= d) {
            s.add(0);
        }
        s.set(d, s.get(d) + root.val);
        dfs1(root.left, d + 1);
        dfs1(root.right, d + 1);
    }

    private void dfs2(TreeNode root, int d) {
        if (root == null) {
            return;
        }
        int l = root.left == null ? 0 : root.left.val;
        int r = root.right == null ? 0 : root.right.val;
        if (root.left != null) {
            root.left.val = s.get(d) - l - r;
        }
        if (root.right != null) {
            root.right.val = s.get(d) - l - r;
        }
        dfs2(root.left, d + 1);
        dfs2(root.right, d + 1);
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
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* replaceValueInTree(TreeNode* root) {
        vector<int> s;
        function<void(TreeNode*, int)> dfs1 = [&](TreeNode* root, int d) {
            if (!root) {
                return;
            }
            if (s.size() <= d) {
                s.push_back(0);
            }
            s[d] += root->val;
            dfs1(root->left, d + 1);
            dfs1(root->right, d + 1);
        };
        function<void(TreeNode*, int)> dfs2 = [&](TreeNode* root, int d) {
            if (!root) {
                return;
            }
            int l = root->left ? root->left->val : 0;
            int r = root->right ? root->right->val : 0;
            if (root->left) {
                root->left->val = s[d] - l - r;
            }
            if (root->right) {
                root->right->val = s[d] - l - r;
            }
            dfs2(root->left, d + 1);
            dfs2(root->right, d + 1);
        };
        dfs1(root, 0);
        root->val = 0;
        dfs2(root, 1);
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
func replaceValueInTree(root *TreeNode) *TreeNode {
	s := []int{}
	var dfs1 func(*TreeNode, int)
	dfs1 = func(root *TreeNode, d int) {
		if root == nil {
			return
		}
		if len(s) <= d {
			s = append(s, 0)
		}
		s[d] += root.Val
		dfs1(root.Left, d+1)
		dfs1(root.Right, d+1)
	}
	var dfs2 func(*TreeNode, int)
	dfs2 = func(root *TreeNode, d int) {
		if root == nil {
			return
		}
		l, r := 0, 0
		if root.Left != nil {
			l = root.Left.Val
		}
		if root.Right != nil {
			r = root.Right.Val
		}
		if root.Left != nil {
			root.Left.Val = s[d] - l - r
		}
		if root.Right != nil {
			root.Right.Val = s[d] - l - r
		}
		dfs2(root.Left, d+1)
		dfs2(root.Right, d+1)
	}
	dfs1(root, 0)
	root.Val = 0
	dfs2(root, 1)
	return root
}
```

### **...**

```

```

<!-- tabs:end -->
