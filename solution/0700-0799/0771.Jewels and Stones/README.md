# [771. 宝石与石头](https://leetcode-cn.com/problems/jewels-and-stones)

[English Version](/solution/0700-0799/0771.Jewels%20and%20Stones/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>&nbsp;给你一个字符串 <code>jewels</code>&nbsp;代表石头中宝石的类型，另有一个字符串 <code>stones</code> 代表你拥有的石头。&nbsp;<code>stones</code>&nbsp;中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。</p>

<p>字母区分大小写，因此 <code>"a"</code> 和 <code>"A"</code> 是不同类型的石头。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>jewels = "aA", stones = "aAAbbbb"
<strong>输出：</strong>3
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>jewels = "z", stones = "ZZ"
<strong>输出：</strong>0<strong>
</strong></pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;=&nbsp;jewels.length, stones.length &lt;= 50</code></li>
	<li><code>jewels</code> 和 <code>stones</code> 仅由英文字母组成</li>
	<li><code>jewels</code> 中的所有字符都是 <strong>唯一的</strong></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

哈希表实现。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        s = set(jewels)
        return sum([1 for c in stones if c in s])
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public int numJewelsInStones(String jewels, String stones) {
        Set<Character> s = new HashSet<>();
        for (char c : jewels.toCharArray()) {
            s.add(c);
        }
        int res = 0;
        for (char c : stones.toCharArray()) {
            res += (s.contains(c) ? 1 : 0);
        }
        return res;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int numJewelsInStones(string jewels, string stones) {
        unordered_set<char> s;
        for (char c : jewels) {
            s.insert(c);
        }
        int res = 0;
        for (char c : stones) {
            res += s.count(c);
        }
        return res;
    }
};
```

### **Go**

```go
func numJewelsInStones(jewels string, stones string) int {
	s := make(map[rune]bool)
	for _, c := range jewels {
		s[c] = true
	}
	res := 0
	for _, c := range stones {
		if s[c] {
			res++
		}
	}
	return res
}
```

### **...**

```

```

<!-- tabs:end -->
