# [2696. Minimum String Length After Removing Substrings](https://leetcode.com/problems/minimum-string-length-after-removing-substrings)

[中文文档](/solution/2600-2699/2696.Minimum%20String%20Length%20After%20Removing%20Substrings/README.md)

## Description

<p>You are given a string <code>s</code> consisting only of <strong>uppercase</strong> English letters.</p>

<p>You can apply some operations to this string where, in one operation, you can remove <strong>any</strong> occurrence of one of the substrings <code>&quot;AB&quot;</code> or <code>&quot;CD&quot;</code> from <code>s</code>.</p>

<p>Return <em>the <strong>minimum</strong> possible length of the resulting string that you can obtain</em>.</p>

<p><strong>Note</strong> that the string concatenates after removing the substring and could produce new <code>&quot;AB&quot;</code> or <code>&quot;CD&quot;</code> substrings.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;ABFCACDB&quot;
<strong>Output:</strong> 2
<strong>Explanation:</strong> We can do the following operations:
- Remove the substring &quot;<u>AB</u>FCACDB&quot;, so s = &quot;FCACDB&quot;.
- Remove the substring &quot;FCA<u>CD</u>B&quot;, so s = &quot;FCAB&quot;.
- Remove the substring &quot;FC<u>AB</u>&quot;, so s = &quot;FC&quot;.
So the resulting length of the string is 2.
It can be shown that it is the minimum length that we can obtain.</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;ACBBD&quot;
<strong>Output:</strong> 5
<strong>Explanation:</strong> We cannot do any operations on the string so the length remains the same.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code>&nbsp;consists only of uppercase English letters.</li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def minLength(self, s: str) -> int:
        stk = [""]
        for c in s:
            if (c == "B" and stk[-1] == "A") or (c == "D" and stk[-1] == "C"):
                stk.pop()
            else:
                stk.append(c)
        return len(stk) - 1
```

### **Java**

```java
class Solution {
    public int minLength(String s) {
        Deque<Character> stk = new ArrayDeque<>();
        stk.push(' ');
        for (char c : s.toCharArray()) {
            if ((c == 'B' && stk.peek() == 'A') || (c == 'D' && stk.peek() == 'C')) {
                stk.pop();
            } else {
                stk.push(c);
            }
        }
        return stk.size() - 1;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int minLength(string s) {
        string stk = " ";
        for (char& c : s) {
            if ((c == 'B' && stk.back() == 'A') || (c == 'D' && stk.back() == 'C')) {
                stk.pop_back();
            } else {
                stk.push_back(c);
            }
        }
        return stk.size() - 1;
    }
};
```

### **Go**

```go
func minLength(s string) int {
	stk := []byte{' '}
	for _, c := range s {
		if (c == 'B' && stk[len(stk)-1] == 'A') || (c == 'D' && stk[len(stk)-1] == 'C') {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, byte(c))
		}
	}
	return len(stk) - 1
}
```

### **TypeScript**

```ts
function minLength(s: string): number {
    const stk: string[] = [''];
    for (const c of s) {
        if (c === 'B' && stk[stk.length - 1] === 'A') {
            stk.pop();
        } else if (c == 'D' && stk[stk.length - 1] === 'C') {
            stk.pop();
        } else {
            stk.push(c);
        }
    }
    return stk.length - 1;
}
```

### **...**

```

```

<!-- tabs:end -->
