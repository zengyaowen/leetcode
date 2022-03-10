# [1461. 检查一个字符串是否包含所有长度为 K 的二进制子串](https://leetcode-cn.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k)

[English Version](/solution/1400-1499/1461.Check%20If%20a%20String%20Contains%20All%20Binary%20Codes%20of%20Size%20K/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给你一个二进制字符串&nbsp;<code>s</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;。如果所有长度为 <code>k</code>&nbsp;的二进制字符串都是 <code>s</code>&nbsp;的子串，请返回 <code>true</code> ，否则请返回 <code>false</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "00110110", k = 2
<strong>输出：</strong>true
<strong>解释：</strong>长度为 2 的二进制串包括 "00"，"01"，"10" 和 "11"。它们分别是 s 中下标为 0，1，3，2 开始的长度为 2 的子串。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "0110", k = 1
<strong>输出：</strong>true
<strong>解释：</strong>长度为 1 的二进制串包括 "0" 和 "1"，显然它们都是 s 的子串。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "0110", k = 2
<strong>输出：</strong>false
<strong>解释：</strong>长度为 2 的二进制串 "00" 没有出现在 s 中。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 5 * 10<sup>5</sup></code></li>
	<li><code>s[i]</code> 不是<code>'0'</code> 就是 <code>'1'</code></li>
	<li><code>1 &lt;= k &lt;= 20</code></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

遍历字符串 s，用一个 set 存储所有长度为 k 的不同子串。只需要判断子串数能否达到 2<sup>k</sup> 即可。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        counter = 1 << k
        exists = set()
        for i in range(k, len(s) + 1):
            if s[i - k: i] not in exists:
                exists.add(s[i - k: i])
                counter -= 1
            if counter == 0:
                return True
        return False
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public boolean hasAllCodes(String s, int k) {
        int counter = 1 << k;
        Set<String> exists = new HashSet<>();
        for (int i = k; i <= s.length(); ++i) {
            String t = s.substring(i - k, i);
            if (!exists.contains(t)) {
                exists.add(t);
                --counter;
            }
            if (counter == 0) {
                return true;
            }
        }
        return false;
    }
}
```

### **...**

```

```

<!-- tabs:end -->
