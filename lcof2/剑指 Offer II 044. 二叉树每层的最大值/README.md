# [剑指 Offer II 044. 二叉树每层的最大值](https://leetcode-cn.com/problems/hPov7L)

## 题目描述

<!-- 这里写题目描述 -->

<p>给定一棵二叉树的根节点&nbsp;<code>root</code> ，请找出该二叉树中每一层的最大值。</p>

<p>&nbsp;</p>

<p><strong>示例1：</strong></p>

<pre>
<strong>输入: </strong>root = [1,3,2,5,3,null,9]
<strong>输出: </strong>[1,3,9]
<strong>解释:</strong>
          1
         / \
        3   2
       / \   \  
      5   3   9 
</pre>

<p><strong>示例2：</strong></p>

<pre>
<strong>输入: </strong>root = [1,2,3]
<strong>输出: </strong>[1,3]
<strong>解释:</strong>
          1
         / \
        2   3
</pre>

<p><strong>示例3：</strong></p>

<pre>
<strong>输入: </strong>root = [1]
<strong>输出: </strong>[1]
</pre>

<p><strong>示例4：</strong></p>

<pre>
<strong>输入: </strong>root = [1,null,2]
<strong>输出: </strong>[1,2]
<strong>解释:</strong>      
&nbsp;          1 
&nbsp;           \
&nbsp;            2     
</pre>

<p><strong>示例5：</strong></p>

<pre>
<strong>输入: </strong>root = []
<strong>输出: </strong>[]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>二叉树的节点个数的范围是 <code>[0,10<sup>4</sup>]</code></li>
	<li><meta charset="UTF-8" /><code>-2<sup>31</sup>&nbsp;&lt;= Node.val &lt;= 2<sup>31</sup>&nbsp;- 1</code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 515&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/">https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/</a></p>


## 解法

<!-- 这里可写通用的实现逻辑 -->

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python

```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java

```

### **C++**

```cpp
class Solution {
public:
    vector<int> largestValues(TreeNode* root) {
        vector<int> res;
        if( !root )
        {
            return res;
        }

        deque<TreeNode* > deq;
        deq.push_back(root);
        while( !deq.empty())
        {
            int size = deq.size();
            int maxnum = INT_MIN;
            for (int i = 0; i < size; i++)
            {
                TreeNode* ptr = deq.front();
                deq.pop_front();
                if(maxnum < ptr->val)
                {
                    maxnum = ptr->val;
                }

                if(ptr->left)
                {
                    deq.push_back(ptr->left);
                }

                if(ptr->right)
                {
                    deq.push_back(ptr->right);
                }
            }

            res.push_back(maxnum);
        }

        return res;
    }
};
```

### **...**

```

```

<!-- tabs:end -->

