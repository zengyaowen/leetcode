# [01.01. Is Unique](https://leetcode-cn.com/problems/is-unique-lcci)

[中文文档](/lcci/01.01.Is%20Unique/README.md)

## Description

<p>Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?</p>

<p><strong>Example 1:</strong></p>

<pre>

<strong>Input: </strong><code>s</code> = &quot;leetcode&quot;

<strong>Output: </strong>false

</pre>

<p><strong>Example 2:</strong></p>

<pre>

<strong>Input: </strong><code>s</code> = &quot;abc&quot;

<strong>Output: </strong>true

</pre>

<p><strong>Note:</strong></p>

<ul>
	<li><code>0 &lt;= len(s) &lt;= 100 </code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def isUnique(self, astr: str) -> bool:
        bitmap = 0
        for c in astr:
            pos = ord(c) - ord('a')
            if (bitmap & (1 << pos)) != 0:
                return False
            bitmap |= (1 << pos)
        return True
```

### **Java**

```java
class Solution {
    public boolean isUnique(String astr) {
        int bitmap = 0;
        for (int i = 0, n = astr.length(); i < n; ++i) {
            int pos = astr.charAt(i) - 'a';
            if ((bitmap & (1 << pos)) != 0) {
                return false;
            }
            bitmap |= (1 << pos);
        }
        return true;
    }
}
```

### **...**

```

```

<!-- tabs:end -->
