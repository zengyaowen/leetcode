# [2788. Split Strings by Separator](https://leetcode.com/problems/split-strings-by-separator)

[中文文档](/solution/2700-2799/2788.Split%20Strings%20by%20Separator/README.md)

## Description

<p>Given an array of strings <code>words</code> and a character <code>separator</code>, <strong>split</strong> each string in <code>words</code> by <code>separator</code>.</p>

<p>Return <em>an array of strings containing the new strings formed after the splits, <strong>excluding empty strings</strong>.</em></p>

<p><strong>Notes</strong></p>

<ul>
	<li><code>separator</code> is used to determine where the split should occur, but it is not included as part of the resulting strings.</li>
	<li>A split may result in more than two strings.</li>
	<li>The resulting strings must maintain the same order as they were initially given.</li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;one.two.three&quot;,&quot;four.five&quot;,&quot;six&quot;], separator = &quot;.&quot;
<strong>Output:</strong> [&quot;one&quot;,&quot;two&quot;,&quot;three&quot;,&quot;four&quot;,&quot;five&quot;,&quot;six&quot;]
<strong>Explanation: </strong>In this example we split as follows:

&quot;one.two.three&quot; splits into &quot;one&quot;, &quot;two&quot;, &quot;three&quot;
&quot;four.five&quot; splits into &quot;four&quot;, &quot;five&quot;
&quot;six&quot; splits into &quot;six&quot; 

Hence, the resulting array is [&quot;one&quot;,&quot;two&quot;,&quot;three&quot;,&quot;four&quot;,&quot;five&quot;,&quot;six&quot;].</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;&#36;easy&#36;&quot;,&quot;&#36;problem&#36;&quot;], separator = &quot;&#36;&quot;
<strong>Output:</strong> [&quot;easy&quot;,&quot;problem&quot;]
<strong>Explanation:</strong> In this example we split as follows: 

&quot;&#36;easy&#36;&quot; splits into &quot;easy&quot; (excluding empty strings)
&quot;&#36;problem&#36;&quot; splits into &quot;problem&quot; (excluding empty strings)

Hence, the resulting array is [&quot;easy&quot;,&quot;problem&quot;].
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;|||&quot;], separator = &quot;|&quot;
<strong>Output:</strong> []
<strong>Explanation:</strong> In this example the resulting split of &quot;|||&quot; will contain only empty strings, so we return an empty array []. </pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 100</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 20</code></li>
	<li>characters in <code>words[i]</code> are either lowercase English letters or characters from the string <code>&quot;.,|&#36;#@&quot;</code> (excluding the quotes)</li>
	<li><code>separator</code> is a character from the string <code>&quot;.,|&#36;#@&quot;</code> (excluding the quotes)</li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def splitWordsBySeparator(self, words: List[str], separator: str) -> List[str]:
        return [s for w in words for s in w.split(separator) if s]
```

### **Java**

```java
import java.util.regex.Pattern;

class Solution {
    public List<String> splitWordsBySeparator(List<String> words, char separator) {
        List<String> ans = new ArrayList<>();
        for (var w : words) {
            for (var s : w.split(Pattern.quote(String.valueOf(separator)))) {
                if (s.length() > 0) {
                    ans.add(s);
                }
            }
        }
        return ans;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<string> splitWordsBySeparator(vector<string>& words, char separator) {
        vector<string> ans;
        for (auto& w : words) {
            for (auto& s : split(w, separator)) {
                if (!s.empty()) {
                    ans.emplace_back(s);
                }
            }
        }
        return ans;
    }

    vector<string> split(string& s, char c) {
        vector<string> res;
        stringstream ss(s);
        string t;
        while (getline(ss, t, c)) {
            res.push_back(t);
        }
        return res;
    }
};
```

### **Go**

```go
func splitWordsBySeparator(words []string, separator byte) (ans []string) {
	for _, w := range words {
		for _, s := range strings.Split(w, string(separator)) {
			if s != "" {
				ans = append(ans, s)
			}
		}
	}
	return
}
```

### **TypeScript**

```ts
function splitWordsBySeparator(words: string[], separator: string): string[] {
    const ans: string[] = [];
    for (const w of words) {
        for (const s of w.split(separator)) {
            if (s.length > 0) {
                ans.push(s);
            }
        }
    }
    return ans;
}
```

### **...**

```

```

<!-- tabs:end -->
