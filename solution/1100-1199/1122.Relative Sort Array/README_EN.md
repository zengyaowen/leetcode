# [1122. Relative Sort Array](https://leetcode.com/problems/relative-sort-array)

[中文文档](/solution/1100-1199/1122.Relative%20Sort%20Array/README.md)

## Description

<p>Given two arrays <code>arr1</code> and <code>arr2</code>, the elements of <code>arr2</code> are distinct, and all elements in <code>arr2</code> are also in <code>arr1</code>.</p>

<p>Sort the elements of <code>arr1</code> such that the relative ordering of items in <code>arr1</code> are the same as in <code>arr2</code>. Elements that do not appear in <code>arr2</code> should be placed at the end of <code>arr1</code> in <strong>ascending</strong> order.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
<strong>Output:</strong> [2,2,2,1,4,3,3,9,6,7,19]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
<strong>Output:</strong> [22,28,8,6,17,44]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= arr1.length, arr2.length &lt;= 1000</code></li>
	<li><code>0 &lt;= arr1[i], arr2[i] &lt;= 1000</code></li>
	<li>All the elements of <code>arr2</code> are <strong>distinct</strong>.</li>
	<li>Each&nbsp;<code>arr2[i]</code> is in <code>arr1</code>.</li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        pos = {x: i for i, x in enumerate(arr2)}
        return sorted(arr1, key=lambda x: pos.get(x, 1000 + x))
```

### **Java**

```java
class Solution {
    public int[] relativeSortArray(int[] arr1, int[] arr2) {
        Map<Integer, Integer> pos = new HashMap<>(arr2.length);
        for (int i = 0; i < arr2.length; ++i) {
            pos.put(arr2[i], i);
        }
        int[][] arr = new int[arr1.length][0];
        for (int i = 0; i < arr.length; ++i) {
            arr[i] = new int[] {arr1[i], pos.getOrDefault(arr1[i], arr2.length + arr1[i])};
        }
        Arrays.sort(arr, (a, b) -> a[1] - b[1]);
        for (int i = 0; i < arr.length; ++i) {
            arr1[i] = arr[i][0];
        }
        return arr1;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    vector<int> relativeSortArray(vector<int>& arr1, vector<int>& arr2) {
        unordered_map<int, int> pos;
        for (int i = 0; i < arr2.size(); ++i) {
            pos[arr2[i]] = i;
        }
        vector<pair<int, int>> arr;
        for (int i = 0; i < arr1.size(); ++i) {
            int j = pos.count(arr1[i]) ? pos[arr1[i]] : arr2.size();
            arr.emplace_back(j, arr1[i]);
        }
        sort(arr.begin(), arr.end());
        for (int i = 0; i < arr1.size(); ++i) {
            arr1[i] = arr[i].second;
        }
        return arr1;
    }
};
```

### **Go**

```go
func relativeSortArray(arr1 []int, arr2 []int) []int {
	pos := map[int]int{}
	for i, x := range arr2 {
		pos[x] = i
	}
	arr := make([][2]int, len(arr1))
	for i, x := range arr1 {
		if p, ok := pos[x]; ok {
			arr[i] = [2]int{p, x}
		} else {
			arr[i] = [2]int{len(arr2), x}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1]
	})
	for i, x := range arr {
		arr1[i] = x[1]
	}
	return arr1
}
```

### **TypeScript**

```ts
function relativeSortArray(arr1: number[], arr2: number[]): number[] {
    const pos: Map<number, number> = new Map();
    for (let i = 0; i < arr2.length; ++i) {
        pos.set(arr2[i], i);
    }
    const arr: number[][] = [];
    for (const x of arr1) {
        const j = pos.get(x) ?? arr2.length;
        arr.push([j, x]);
    }
    arr.sort((a, b) => a[0] - b[0] || a[1] - b[1]);
    return arr.map(a => a[1]);
}
```

### **...**

```

```

<!-- tabs:end -->
