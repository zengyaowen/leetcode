# [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number)

[English Version](/solution/0000-0099/0017.Letter%20Combinations%20of%20a%20Phone%20Number/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>给定一个仅包含数字&nbsp;<code>2-9</code>&nbsp;的字符串，返回所有它能表示的字母组合。答案可以按 <strong>任意顺序</strong> 返回。</p>

<p>给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。</p>

<p><img src="https://cdn.jsdelivr.net/gh/doocs/leetcode@main/solution/0000-0099/0017.Letter%20Combinations%20of%20a%20Phone%20Number/images/200px-telephone-keypad2svg.png" style="width: 200px;" /></p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>digits = "23"
<strong>输出：</strong>["ad","ae","af","bd","be","bf","cd","ce","cf"]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>digits = ""
<strong>输出：</strong>[]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>digits = "2"
<strong>输出：</strong>["a","b","c"]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= digits.length &lt;= 4</code></li>
	<li><code>digits[i]</code> 是范围 <code>['2', '9']</code> 的一个数字。</li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        n = len(digits)
        if n == 0:
            return []
        chars = ['abc', 'def', 'ghi', 'jkl', 'mno', 'pqrs', 'tuv', 'wxyz']
        strs = [chars[int(d) - 2] for d in digits]
        res = []
        for s in strs:
            if not res:
                res = list(s)
            else:
                cache = []
                for item in res:
                    for letter in s:
                        cache.append(item + letter)
                res = cache
        return res
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public List<String> letterCombinations(String digits) {
        int n;
        if ((n = digits.length()) == 0) return Collections.emptyList();
        List<String> chars = Arrays.asList("abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz");

        List<String> strs = new ArrayList<>();
        for (char c : digits.toCharArray()) {
            strs.add(chars.get(c - '0' - 2));
        }
        List<String> res = new ArrayList<>();
        for (String str : strs) {
            if (res.size() == 0) {
                for (char c : str.toCharArray()) {
                    res.add(String.valueOf(c));
                }
            } else {
                List<String> cache = new ArrayList<>();
                for (String item : res) {
                    for (char c : str.toCharArray()) {
                        cache.add(item + String.valueOf(c));
                    }
                }
                res = cache;
            }
        }
        return res;
    }
}
```

### **Go**

```go
var table = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}
	var result = table[string(digits[0])]
	for i := 1; i < len(digits); i++ {
		t := table[string(digits[i])]
		nr := make([]string, len(result)*len(t))
		for j := 0; j < len(result); j++ {
			for k := 0; k < len(t); k++ {
				nr[len(t)*j+k] = result[j] + t[k]
			}
		}
		result = nr
	}
	return result
}
```

### **C#**

```cs
using System.Collections.Generic;
using System.Linq;

public class Solution {
    private static string[] chars = {
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz"
    };

    public IList<string> LetterCombinations(string digits) {
        var numbers = digits.Where(d => d >= '2' && d <= '9').Select(d => d - '2').ToArray();
        var states = new int[numbers.Length];
        var results = new List<string>();
        if (numbers.Length == 0) return results;
        while (true) {
            results.Add(new string(states.Select((s, j) => chars[numbers[j]][s]).ToArray()));
            var i = states.Length - 1;
            ++states[i];
            while (i >= 0 && states[i] == chars[numbers[i]].Length)
            {
                states[i] = 0;
                --i;
                if (i >= 0)
                {
                    ++states[i];
                }
            }
            if (i < 0) return results;
        }
    }
}
```

### **TypeScript**

```ts
const map = {
    '2': ['a', 'b', 'c'],
    '3': ['d', 'e', 'f'],
    '4': ['g', 'h', 'i'],
    '5': ['j', 'k', 'l'],
    '6': ['m', 'n', 'o'],
    '7': ['p', 'q', 'r', 's'],
    '8': ['t', 'u', 'v'],
    '9': ['w', 'x', 'y', 'z'],
};

function letterCombinations(digits: string): string[] {
    const n = digits.length;
    if (n === 0) {
        return [];
    }
    const res = [];
    const dfs = (i: number, str: string) => {
        if (i === n) {
            res.push(str);
            return;
        }
        for (const c of map[digits[i]]) {
            dfs(i + 1, str + c);
        }
    };
    dfs(0, '');
    return res;
}
```

```ts
const map = {
    '2': ['a', 'b', 'c'],
    '3': ['d', 'e', 'f'],
    '4': ['g', 'h', 'i'],
    '5': ['j', 'k', 'l'],
    '6': ['m', 'n', 'o'],
    '7': ['p', 'q', 'r', 's'],
    '8': ['t', 'u', 'v'],
    '9': ['w', 'x', 'y', 'z'],
};

function letterCombinations(digits: string): string[] {
    const n = digits.length;
    if (n === 0) {
        return [];
    }
    const dfs = (i: number, ss: string[]) => {
        if (i === n) {
            return ss;
        }
        const t = [];
        for (const c of map[digits[i]]) {
            for (const s of ss) {
                t.push(s + c);
            }
        }
        return dfs(i + 1, t);
    };
    return dfs(1, map[digits[0]]);
}
```

### **Rust**

```rust
use std::collections::HashMap;

impl Solution {
    fn dfs(
        i: usize,
        s: &mut String,
        cs: &Vec<char>,
        map: &HashMap<char, String>,
        res: &mut Vec<String>,
    ) {
        if i == cs.len() {
            res.push(s.clone());
            return;
        }
        for c in map.get(&cs[i]).unwrap().chars() {
            s.push(c);
            Self::dfs(i + 1, s, cs, map, res);
            s.pop();
        }
    }

    pub fn letter_combinations(digits: String) -> Vec<String> {
        let mut res = vec![];
        if digits.is_empty() {
            return res;
        }

        let mut map = HashMap::new();
        map.insert('2', String::from("abc"));
        map.insert('3', String::from("def"));
        map.insert('4', String::from("ghi"));
        map.insert('5', String::from("jkl"));
        map.insert('6', String::from("mno"));
        map.insert('7', String::from("pqrs"));
        map.insert('8', String::from("tuv"));
        map.insert('9', String::from("wxyz"));

        Self::dfs(
            0,
            &mut String::new(),
            &digits.chars().collect(),
            &map,
            &mut res,
        );
        res
    }
}
```

### **...**

```

```

<!-- tabs:end -->
