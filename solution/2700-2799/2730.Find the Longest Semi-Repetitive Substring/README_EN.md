# [2730. Find the Longest Semi-Repetitive Substring](https://leetcode.com/problems/find-the-longest-semi-repetitive-substring)

[中文文档](/solution/2700-2799/2730.Find%20the%20Longest%20Semi-Repetitive%20Substring/README.md)

## Description

<p>You are given a <strong>0-indexed</strong> string <code>s</code> that consists of digits from <code>0</code> to <code>9</code>.</p>

<p>A string <code>t</code> is called a <strong>semi-repetitive</strong> if there is at most one consecutive pair of the same digits inside <code>t</code>. For example, <code>0010</code>, <code>002020</code>, <code>0123</code>, <code>2002</code>, and <code>54944</code> are semi-repetitive while&nbsp;<code>00101022</code>, and <code>1101234883</code> are not.</p>

<p>Return <em>the length of the longest semi-repetitive substring inside</em> <code>s</code>.</p>

<p>A <b>substring</b> is a contiguous <strong>non-empty</strong> sequence of characters within a string.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;52233&quot;
<strong>Output:</strong> 4
<strong>Explanation:</strong> The longest semi-repetitive substring is &quot;5223&quot;, which starts at i = 0 and ends at j = 3. 
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;5494&quot;
<strong>Output:</strong> 4
<strong>Explanation:</strong> s is a semi-reptitive string, so the answer is 4.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;1111111&quot;
<strong>Output:</strong> 2
<strong>Explanation:</strong> The longest semi-repetitive substring is &quot;11&quot;, which starts at i = 0 and ends at j = 1.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 50</code></li>
	<li><code>&#39;0&#39; &lt;= s[i] &lt;= &#39;9&#39;</code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def longestSemiRepetitiveSubstring(self, s: str) -> int:
        n = len(s)
        ans = cnt = j = 0
        for i in range(n):
            if i and s[i] == s[i - 1]:
                cnt += 1
            while cnt > 1:
                if s[j] == s[j + 1]:
                    cnt -= 1
                j += 1
            ans = max(ans, i - j + 1)
        return ans
```

### **Java**

```java
class Solution {
    public int longestSemiRepetitiveSubstring(String s) {
        int n = s.length();
        int ans = 0;
        for (int i = 0, j = 0, cnt = 0; i < n; ++i) {
            if (i > 0 && s.charAt(i) == s.charAt(i - 1)) {
                ++cnt;
            }
            while (cnt > 1) {
                if (s.charAt(j) == s.charAt(j + 1)) {
                    --cnt;
                }
                ++j;
            }
            ans = Math.max(ans, i - j + 1);
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int longestSemiRepetitiveSubstring(string s) {
        int n = s.size();
        int ans = 0;
        for (int i = 0, j = 0, cnt = 0; i < n; ++i) {
            if (i && s[i] == s[i - 1]) {
                ++cnt;
            }
            while (cnt > 1) {
                if (s[j] == s[j + 1]) {
                    --cnt;
                }
                ++j;
            }
            ans = max(ans, i - j + 1);
        }
        return ans;
    }
};
```

### **Go**

```go
func longestSemiRepetitiveSubstring(s string) (ans int) {
	n := len(s)
	for i, j, cnt := 0, 0, 0; i < n; i++ {
		if i > 0 && s[i] == s[i-1] {
			cnt++
		}
		for cnt > 1 {
			if s[j] == s[j+1] {
				cnt--
			}
			j++
		}
		ans = max(ans, i-j+1)
	}
	return
}
```

### **TypeScript**

```ts
function longestSemiRepetitiveSubstring(s: string): number {
    const n = s.length;
    let ans = 0;
    for (let i = 0, j = 0, cnt = 0; i < n; ++i) {
        if (i > 0 && s[i] === s[i - 1]) {
            ++cnt;
        }
        while (cnt > 1) {
            if (s[j] === s[j + 1]) {
                --cnt;
            }
            ++j;
        }
        ans = Math.max(ans, i - j + 1);
    }
    return ans;
}
```

### **...**

```

```

<!-- tabs:end -->
