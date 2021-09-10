# [399. Evaluate Division](https://leetcode.com/problems/evaluate-division)

[中文文档](/solution/0300-0399/0399.Evaluate%20Division/README.md)

## Description

<p>You are given an array of variable pairs <code>equations</code> and an array of real numbers <code>values</code>, where <code>equations[i] = [A<sub>i</sub>, B<sub>i</sub>]</code> and <code>values[i]</code> represent the equation <code>A<sub>i</sub> / B<sub>i</sub> = values[i]</code>. Each <code>A<sub>i</sub></code> or <code>B<sub>i</sub></code> is a string that represents a single variable.</p>

<p>You are also given some <code>queries</code>, where <code>queries[j] = [C<sub>j</sub>, D<sub>j</sub>]</code> represents the <code>j<sup>th</sup></code> query where you must find the answer for <code>C<sub>j</sub> / D<sub>j</sub> = ?</code>.</p>

<p>Return <em>the answers to all queries</em>. If a single answer cannot be determined, return <code>-1.0</code>.</p>

<p><strong>Note:</strong> The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> equations = [[&quot;a&quot;,&quot;b&quot;],[&quot;b&quot;,&quot;c&quot;]], values = [2.0,3.0], queries = [[&quot;a&quot;,&quot;c&quot;],[&quot;b&quot;,&quot;a&quot;],[&quot;a&quot;,&quot;e&quot;],[&quot;a&quot;,&quot;a&quot;],[&quot;x&quot;,&quot;x&quot;]]
<strong>Output:</strong> [6.00000,0.50000,-1.00000,1.00000,-1.00000]
<strong>Explanation:</strong> 
Given: <em>a / b = 2.0</em>, <em>b / c = 3.0</em>
queries are: <em>a / c = ?</em>, <em>b / a = ?</em>, <em>a / e = ?</em>, <em>a / a = ?</em>, <em>x / x = ?</em>
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> equations = [[&quot;a&quot;,&quot;b&quot;],[&quot;b&quot;,&quot;c&quot;],[&quot;bc&quot;,&quot;cd&quot;]], values = [1.5,2.5,5.0], queries = [[&quot;a&quot;,&quot;c&quot;],[&quot;c&quot;,&quot;b&quot;],[&quot;bc&quot;,&quot;cd&quot;],[&quot;cd&quot;,&quot;bc&quot;]]
<strong>Output:</strong> [3.75000,0.40000,5.00000,0.20000]
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> equations = [[&quot;a&quot;,&quot;b&quot;]], values = [0.5], queries = [[&quot;a&quot;,&quot;b&quot;],[&quot;b&quot;,&quot;a&quot;],[&quot;a&quot;,&quot;c&quot;],[&quot;x&quot;,&quot;y&quot;]]
<strong>Output:</strong> [0.50000,2.00000,-1.00000,-1.00000]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= equations.length &lt;= 20</code></li>
	<li><code>equations[i].length == 2</code></li>
	<li><code>1 &lt;= A<sub>i</sub>.length, B<sub>i</sub>.length &lt;= 5</code></li>
	<li><code>values.length == equations.length</code></li>
	<li><code>0.0 &lt; values[i] &lt;= 20.0</code></li>
	<li><code>1 &lt;= queries.length &lt;= 20</code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>1 &lt;= C<sub>j</sub>.length, D<sub>j</sub>.length &lt;= 5</code></li>
	<li><code>A<sub>i</sub>, B<sub>i</sub>, C<sub>j</sub>, D<sub>j</sub></code> consist of lower case English letters and digits.</li>
</ul>

## Solutions

Union find.

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        w = collections.defaultdict(lambda: 1)
        p = collections.defaultdict()
        for a, b in equations:
            p[a] = a
            p[b] = b
        
        def find(x):
            if p[x] != x:
                origin = p[x]
                p[x] = find(p[x])
                w[x] *= w[origin]
            return p[x]

        for i, e in enumerate(equations):
            pa, pb = find(e[0]), find(e[1])
            if pa == pb:
                continue
            p[pa] = pb
            w[pa] = w[e[1]] * values[i] / w[e[0]]
        
        return [-1 if c not in p or d not in p or find(c) != find(d) else w[c] / w[d] for c, d in queries]
```

### **Java**

```java
class Solution {
    private Map<String, String> p;
    private Map<String, Double> w;

    public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
        int n = equations.size();
        p = new HashMap<>(n << 1);
        w = new HashMap<>(n << 1);
        for (List<String> e : equations) {
            p.put(e.get(0), e.get(0));
            p.put(e.get(1), e.get(1));
            w.put(e.get(0), 1.0);
            w.put(e.get(1), 1.0);
        }
        for (int i = 0; i < n; ++i) {
            List<String> e = equations.get(i);
            String a = e.get(0), b = e.get(1);
            String pa = find(a), pb = find(b);
            if (Objects.equals(pa, pb)) {
                continue;
            }
            p.put(pa, pb);
            w.put(pa, w.get(b) * values[i] / w.get(a));
        }
        int m = queries.size();
        double[] res = new double[m];
        for (int i = 0; i < m; ++i) {
            String c = queries.get(i).get(0), d = queries.get(i).get(1);
            res[i] = !p.containsKey(c) || !p.containsKey(d) || !Objects.equals(find(c), find(d)) ? - 1.0 : w.get(c) / w.get(d);
        }
        return res;
    }

    private String find(String x) {
        if (!Objects.equals(p.get(x), x)) {
            String origin = p.get(x);
            p.put(x, find(p.get(x)));
            w.put(x, w.get(x) * w.get(origin));
        }
        return p.get(x);
    }
}
```

### **C++**

```cpp
class Solution {
public:
    unordered_map<string, string> p;
    unordered_map<string, double> w;

    vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
        int n = equations.size();
        for (auto e : equations)
        {
            p[e[0]] = e[0];
            p[e[1]] = e[1];
            w[e[0]] = 1.0;
            w[e[1]] = 1.0;
        }
        for (int i = 0; i < n; ++i)
        {
            vector<string> e = equations[i];
            string a = e[0], b = e[1];
            string pa = find(a), pb = find(b);
            if (pa == pb) continue;
            p[pa] = pb;
            w[pa] = w[b] * values[i] / w[a];
        }
        int m = queries.size();
        vector<double> res(m);
        for (int i = 0; i < m; ++i)
        {
            string c = queries[i][0], d = queries[i][1];
            res[i] = p.find(c) == p.end() || p.find(d) == p.end() || find(c) != find(d) ? -1.0 : w[c] / w[d];
        }
        return res;
    }

    string find(string x) {
        if (p[x] != x)
        {
            string origin = p[x];
            p[x] = find(p[x]);
            w[x] *= w[origin];
        }
        return p[x];
    }
};
```

### **Go**

```go
var p map[string]string
var w map[string]float64

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	p = make(map[string]string)
	w = make(map[string]float64)
	for _, e := range equations {
		p[e[0]] = e[0]
		p[e[1]] = e[1]
		w[e[0]] = 1.0
		w[e[1]] = 1.0
	}
	for i, e := range equations {
		a, b := e[0], e[1]
		pa, pb := find(a), find(b)
		if pa == pb {
			continue
		}
		p[pa] = pb
		w[pa] = w[b] * values[i] / w[a]
	}
	var res []float64
	for _, e := range queries {
		c, d := e[0], e[1]
		if p[c] == "" || p[d] == "" || find(c) != find(d) {
			res = append(res, -1.0)
		} else {
			res = append(res, w[c]/w[d])
		}
	}
	return res
}

func find(x string) string {
	if p[x] != x {
		origin := p[x]
		p[x] = find(p[x])
		w[x] *= w[origin]
	}
	return p[x]
}
```

### **...**

```

```

<!-- tabs:end -->
