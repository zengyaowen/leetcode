# [869. Reordered Power of 2](https://leetcode.com/problems/reordered-power-of-2)

[中文文档](/solution/0800-0899/0869.Reordered%20Power%20of%202/README.md)

## Description

<p>You are given an integer <code>n</code>. We reorder the digits in any order (including the original order) such that the leading digit is not zero.</p>

<p>Return <code>true</code> <em>if and only if we can do this so that the resulting number is a power of two</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 1
<strong>Output:</strong> true
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 10
<strong>Output:</strong> false
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>9</sup></code></li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def reorderedPowerOf2(self, n: int) -> bool:
        def convert(n):
            counter = [0] * 10
            while n > 0:
                counter[n % 10] += 1
                n //= 10
            return counter

        i, s = 1, convert(n)
        while i <= 10 ** 9:
            if convert(i) == s:
                return True
            i <<= 1
        return False
```

### **Java**

```java
class Solution {
    public boolean reorderedPowerOf2(int n) {
        String s = convert(n);
        for (int i = 1; i <= Math.pow(10, 9); i <<= 1) {
            if (s.equals(convert(i))) {
                return true;
            }
        }
        return false;
    }

    private String convert(int n) {
        char[] counter = new char[10];
        while (n > 0) {
            ++counter[n % 10];
            n /= 10;
        }
        return new String(counter);
    }
}
```

### **C++**

```cpp
class Solution {
public:
    bool reorderedPowerOf2(int n) {
        vector<int> s = convert(n);
        for (int i = 1; i <= pow(10, 9); i <<= 1)
            if (s == convert(i)) return true;
        return false;
    }

    vector<int> convert(int n) {
        vector<int> counter(10);
        while (n)
        {
            ++counter[n % 10];
            n /= 10;
        }
        return counter;
    }
};
```

### **Go**

```go
func reorderedPowerOf2(n int) bool {
	convert := func(n int) []byte {
		counter := make([]byte, 10)
		for n > 0 {
			counter[n%10]++
			n /= 10
		}
		return counter
	}

	s := convert(n)
	for i := 1; i <= 1e9; i <<= 1 {
		if bytes.Equal(s, convert(i)) {
			return true
		}
	}
	return false
}
```

### **...**

```

```

<!-- tabs:end -->
