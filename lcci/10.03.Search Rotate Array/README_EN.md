# [10.03. Search Rotate Array](https://leetcode.cn/problems/search-rotate-array-lcci)

[中文文档](/lcci/10.03.Search%20Rotate%20Array/README.md)

## Description

<p>Given a sorted array of n integers that has been rotated an unknown number of times, write code to find an element in the array. You may assume that the array was originally sorted in increasing order. If there are more than one target elements in the array, return the smallest index.</p>
<p><strong>Example1:</strong></p>
<pre>

<strong> Input</strong>: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5

<strong> Output</strong>: 8 (the index of 5 in the array)

</pre>
<p><strong>Example2:</strong></p>
<pre>

<strong> Input</strong>: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11

<strong> Output</strong>: -1 (not found)

</pre>
<p><strong>Note:</strong></p>
<ol>
	<li><code>1 &lt;= arr.length &lt;= 1000000</code></li>
</ol>

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class Solution:
    def search(self, arr: List[int], target: int) -> int:
        l, r = 0, len(arr) - 1
        while arr[l] == arr[r]:
            r -= 1
        while l < r:
            mid = (l + r) >> 1
            if arr[mid] > arr[r]:
                if arr[l] <= target <= arr[mid]:
                    r = mid
                else:
                    l = mid + 1
            elif arr[mid] < arr[r]:
                if arr[mid] < target <= arr[r]:
                    l = mid + 1
                else:
                    r = mid
            else:
                r -= 1
        return l if arr[l] == target else -1
```

### **Java**

```java
class Solution {
    public int search(int[] arr, int target) {
        int l = 0, r = arr.length - 1;
        while (arr[l] == arr[r]) {
            --r;
        }
        while (l < r) {
            int mid = (l + r) >> 1;
            if (arr[mid] > arr[r]) {
                if (arr[l] <= target && target <= arr[mid]) {
                    r = mid;
                } else {
                    l = mid + 1;
                }
            } else if (arr[mid] < arr[r]) {
                if (arr[mid] < target && target <= arr[r]) {
                    l = mid + 1;
                } else {
                    r = mid;
                }
            } else {
                --r;
            }
        }
        return arr[l] == target ? l : -1;
    }
}
```

### **C++**

```cpp
class Solution {
public:
    int search(vector<int>& arr, int target) {
        int l = 0, r = arr.size() - 1;
        while (arr[l] == arr[r]) {
            --r;
        }
        while (l < r) {
            int mid = (l + r) >> 1;
            if (arr[mid] > arr[r]) {
                if (arr[l] <= target && target <= arr[mid]) {
                    r = mid;
                } else {
                    l = mid + 1;
                }
            } else if (arr[mid] < arr[r]) {
                if (arr[mid] < target && target <= arr[r]) {
                    l = mid + 1;
                } else {
                    r = mid;
                }
            } else {
                --r;
            }
        }
        return arr[l] == target ? l : -1;
    }
};
```

### **Go**

```go
func search(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for arr[l] == arr[r] {
		r--
	}
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] > arr[r] {
			if arr[l] <= target && target <= arr[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else if arr[mid] < arr[r] {
			if arr[mid] < target && target <= arr[r] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			r--
		}
	}
	if arr[l] == target {
		return l
	}
	return -1
}
```

### **TypeScript**

```ts
function search(arr: number[], target: number): number {
    let [l, r] = [0, arr.length - 1];
    while (arr[l] === arr[r]) {
        --r;
    }
    while (l < r) {
        const mid = (l + r) >> 1;
        if (arr[mid] > arr[r]) {
            if (arr[l] <= target && target <= arr[mid]) {
                r = mid;
            } else {
                l = mid + 1;
            }
        } else if (arr[mid] < arr[r]) {
            if (arr[mid] < target && target <= arr[r]) {
                l = mid + 1;
            } else {
                r = mid;
            }
        } else {
            --r;
        }
    }
    return arr[l] === target ? l : -1;
}
```

### **...**

```

```

<!-- tabs:end -->
