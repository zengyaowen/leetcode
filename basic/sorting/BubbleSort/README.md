# 冒泡排序

定义一个布尔变量 `hasChange`，用来标记每轮是否进行了交换。在每轮遍历开始时，将 `hasChange` 设置为 false。

若当轮没有发生交换，说明此时数组已经按照升序排列，`hasChange` 依然是为 false。此时外层循环直接退出，排序结束。

## 代码示例

<!-- tabs:start -->

### **Java**

```java
import java.util.Arrays;

public class BubbleSort {

    private static void bubbleSort(int[] nums) {
        boolean hasChange = true;
        for (int i = 0, n = nums.length; i < n - 1 && hasChange; ++i) {
            hasChange = false;
            for (int j = 0; j < n - i - 1; ++j) {
                if (nums[j] > nums[j + 1]) {
                    swap(nums, j, j + 1);
                    hasChange = true;
                }
            }
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int t = nums[i];
        nums[i] = nums[j];
        nums[j] = t;
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 7, 9, 5, 8};
        bubbleSort(nums);
        System.out.println(Arrays.toString(nums));
    }
}
```

### **JavaScript**

```js
function bubbleSort(inputArr) {
    for (let i = inputArr.length - 1; i > 0; i--) {
        let hasChange = false;
        for (let j = 0; j < i; j++) {
            if (inputArr[j] > inputArr[j + 1]) {
                const temp = inputArr[j];
                inputArr[j] = inputArr[j + 1];
                inputArr[j + 1] = temp;
                hasChange = true;
            }
        }

        if (!hasChange) {
            break;
        }
    }

    return inputArr;
}

const arr = [6, 3, 2, 1, 5];
console.log(bubbleSort(arr));
```

### **Go**

```go
package main

import "fmt"

func bubbleSort(nums []int) {
	hasChange := true
	for i, n := 0, len(nums); i < n-1 && hasChange; i++ {
		hasChange = false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				hasChange = true
			}
		}
	}
}

func main() {
	nums := []int{1, 2, 7, 9, 5, 8}
	bubbleSort(nums)
	fmt.Println(nums)
}
```

### **C++**

```cpp
#include <iostream>
#include <vector>
#include <string>

using namespace std;

/* 简单版本 */
void bubblesort(vector<int> &vec)
{
    for (int i = 0; i < vec.size() - 1; i++)
    {
        for (int j = 0; j < vec.size() - i - 1; j++)
        {
            if (vec[j] > vec[j + 1])
            {
                swap(vec[j], vec[j + 1]);
            }
        }
    }
}

/* 改进版本 */
void bubblesort1(vector<int> &vec)
{
    for (int i = 0; i < vec.size() - 1; i++)
    {
        bool exchange = false;
        for (int j = 0; j < vec.size() - i - 1; j++)
        {
            if (vec[j] > vec[j + 1])
            {
                swap(vec[j], vec[j + 1]);
                exchange = true;
            }
        }

        if (!exchange)
        {
            break;
        }
    }
}

void printvec(const vector<int> &vec, const string &strbegin = "", const string &strend = "")
{
    cout << strbegin << endl;
    for (auto val : vec)
    {
        cout << val << "\t";
    }

    cout << endl;
    cout << strend << endl;
}

int main(void)
{
    vector<int> vec = {9, 8, 7, 6, 5, 4, 3, 2, 1, 0};
    printvec(vec);

    bubblesort1(vec);

    printvec(vec, "after sort", "");
}
```

### **Rust**

```rust
fn bubble_sort(nums: &mut Vec<i32>) {
    let n = nums.len();
    for i in 0..n - 1 {
        for j in i..n {
            if nums[i] > nums[j] {
                let temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
            }
        }
    }
}

fn main() {
    let mut nums = vec![1, 2, 7, 9, 5, 8];
    bubble_sort(&mut nums);
    println!("{:?}", nums);
}
```

### **C#**

```cs
using static System.Console;
namespace Pro;
public class Program
{
    public static void Main()
    {
        int[] test = new int[] { 56, 876, 34, 23, 45, 501, 2, 3, 4, 6, 5, 7, 8, 9, 11, 10, 12, 23, 34 };
        BubbleSortNums(test);
        foreach (var item in test)
        {
            WriteLine(item);
        }
        ReadLine();
    }
    public static void BubbleSortNums(int[] nums)
    {
        int numchange = 0;
        for (int initial = 0; initial < nums.Length - numchange; initial++)
        {
            WriteLine($"{initial} start ");
            // 记录此值 用于迭代开始位置
            bool changelog = false;
            for (int second_sortnum = initial; second_sortnum < nums.Length - 1; second_sortnum++)
            {
                if (nums[second_sortnum] > nums[second_sortnum + 1])
                {
                    swap(ref nums[second_sortnum], ref nums[second_sortnum + 1]);
                    if (!changelog)
                    {
                        // 记录转换的位置，让initial开始位置从转换位置前开始
                        initial = ((second_sortnum - 2) > 0) ? (second_sortnum - 2) : -1;
                        numchange += 1;
                    }
                    changelog = true;
                }
            }
        }
    }
    private static void swap(ref int compare_left, ref int compare_right)
    {
        int temp = compare_left;
        compare_left = compare_right;
        compare_right = temp;
    }
}
```

<!-- tabs:end -->

## 算法分析

空间复杂度 O(1)、时间复杂度 O(n²)。

分情况讨论：

1. 给定的数组按照顺序已经排好：只需要进行 `n-1` 次比较，两两交换次数为 0，时间复杂度为 O(n)，这是最好的情况。
2. 给定的数组按照逆序排列：需要进行 `n*(n-1)/2` 次比较，时间复杂度为 O(n²)，这是最坏的情况。
3. 给定的数组杂乱无章。在这种情况下，平均时间复杂度 O(n²)。

因此，时间复杂度是 O(n²)，这是一种稳定的排序算法。

> 稳定是指，两个相等的数，在排序过后，相对位置保持不变。
