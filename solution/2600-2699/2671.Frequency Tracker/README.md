# [2671. 频率跟踪器](https://leetcode.cn/problems/frequency-tracker)

[English Version](/solution/2600-2699/2671.Frequency%20Tracker/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>请你设计并实现一个能够对其中的值进行跟踪的数据结构，并支持对频率相关查询进行应答。</p>

<p>实现 <code>FrequencyTracker</code> 类：</p>

<ul>
	<li><code>FrequencyTracker()</code>：使用一个空数组初始化 <code>FrequencyTracker</code> 对象。</li>
	<li><code>void add(int number)</code>：添加一个 <code>number</code> 到数据结构中。</li>
	<li><code>void deleteOne(int number)</code>：从数据结构中删除一个 <code>number</code> 。数据结构 <strong>可能不包含</strong> <code>number</code> ，在这种情况下不删除任何内容。</li>
	<li><code>bool hasFrequency(int frequency)</code>: 如果数据结构中存在出现 <code>frequency</code> 次的数字，则返回 <code>true</code>，否则返回 <code>false</code>。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入</strong>
["FrequencyTracker", "add", "add", "hasFrequency"]
[[], [3], [3], [2]]
<strong>输出</strong>
[null, null, null, true]

<strong>解释</strong>
FrequencyTracker frequencyTracker = new FrequencyTracker();
frequencyTracker.add(3); // 数据结构现在包含 [3]
frequencyTracker.add(3); // 数据结构现在包含 [3, 3]
frequencyTracker.hasFrequency(2); // 返回 true ，因为 3 出现 2 次
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入</strong>
["FrequencyTracker", "add", "deleteOne", "hasFrequency"]
[[], [1], [1], [1]]
<strong>输出</strong>
[null, null, null, false]

<strong>解释</strong>
FrequencyTracker frequencyTracker = new FrequencyTracker();
frequencyTracker.add(1); // 数据结构现在包含 [1]
frequencyTracker.deleteOne(1); // 数据结构现在为空 []
frequencyTracker.hasFrequency(1); // 返回 false ，因为数据结构为空
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入</strong>
["FrequencyTracker", "hasFrequency", "add", "hasFrequency"]
[[], [2], [3], [1]]
<strong>输出</strong>
[null, false, null, true]

<strong>解释</strong>
FrequencyTracker frequencyTracker = new FrequencyTracker();
frequencyTracker.hasFrequency(2); // 返回 false ，因为数据结构为空
frequencyTracker.add(3); // 数据结构现在包含 [3]
frequencyTracker.hasFrequency(1); // 返回 true ，因为 3 出现 1 次
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= number &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= frequency &lt;= 10<sup>5</sup></code></li>
	<li>最多调用 <code>add</code>、<code>deleteOne</code> 和 <code>hasFrequency</code> <strong>共计</strong> <code>2 *&nbsp;10<sup>5</sup></code> 次</li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class FrequencyTracker:
    def __init__(self):
        self.cnt = defaultdict(int)
        self.freq = defaultdict(int)

    def add(self, number: int) -> None:
        if self.freq[self.cnt[number]] > 0:
            self.freq[self.cnt[number]] -= 1
        self.cnt[number] += 1
        self.freq[self.cnt[number]] += 1

    def deleteOne(self, number: int) -> None:
        if self.cnt[number] == 0:
            return
        self.freq[self.cnt[number]] -= 1
        self.cnt[number] -= 1
        self.freq[self.cnt[number]] += 1

    def hasFrequency(self, frequency: int) -> bool:
        return self.freq[frequency] > 0


# Your FrequencyTracker object will be instantiated and called as such:
# obj = FrequencyTracker()
# obj.add(number)
# obj.deleteOne(number)
# param_3 = obj.hasFrequency(frequency)
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class FrequencyTracker {
    private Map<Integer, Integer> cnt = new HashMap<>();
    private Map<Integer, Integer> freq = new HashMap<>();

    public FrequencyTracker() {
    }

    public void add(int number) {
        int f = cnt.getOrDefault(number, 0);
        if (freq.getOrDefault(f, 0) > 0) {
            freq.merge(f, -1, Integer::sum);
        }
        cnt.merge(number, 1, Integer::sum);
        freq.merge(f + 1, 1, Integer::sum);
    }

    public void deleteOne(int number) {
        int f = cnt.getOrDefault(number, 0);
        if (f == 0) {
            return;
        }
        freq.merge(f, -1, Integer::sum);
        cnt.merge(number, -1, Integer::sum);
        freq.merge(f - 1, 1, Integer::sum);
    }

    public boolean hasFrequency(int frequency) {
        return freq.getOrDefault(frequency, 0) > 0;
    }
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * FrequencyTracker obj = new FrequencyTracker();
 * obj.add(number);
 * obj.deleteOne(number);
 * boolean param_3 = obj.hasFrequency(frequency);
 */
```

### **C++**

```cpp
class FrequencyTracker {
public:
    FrequencyTracker() {

    }

    void add(int number) {
        int f =  cnt[number];
        if (f > 0) {
            freq[f]--;
        }
        cnt[number]++;
        freq[f + 1]++;
    }

    void deleteOne(int number) {
        int f = cnt[number];
        if (f == 0) {
            return;
        }
        freq[f]--;
        cnt[number]--;
        freq[f - 1]++;
    }

    bool hasFrequency(int frequency) {
        return freq[frequency] > 0;
    }

private:
    unordered_map<int, int> cnt;
    unordered_map<int, int> freq;
};

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * FrequencyTracker* obj = new FrequencyTracker();
 * obj->add(number);
 * obj->deleteOne(number);
 * bool param_3 = obj->hasFrequency(frequency);
 */
```

### **Go**

```go
type FrequencyTracker struct {
	cnt  map[int]int
	freq map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (this *FrequencyTracker) Add(number int) {
	f := this.cnt[number]
	if f > 0 {
		this.freq[f]--
	}
	this.cnt[number]++
	this.freq[f+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	f := this.cnt[number]
	if f == 0 {
		return
	}
	this.freq[f]--
	this.cnt[number]--
	this.freq[f-1]++
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freq[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
```

### **TypeScript**

```ts
class FrequencyTracker {
    private cnt: Map<number, number>;
    private freq: Map<number, number>;

    constructor() {
        this.cnt = new Map();
        this.freq = new Map();
    }

    add(number: number): void {
        const f = this.cnt.get(number) || 0;
        if (f > 0) {
            this.freq.set(f, (this.freq.get(f) || 0) - 1);
        }
        this.cnt.set(number, f + 1);
        this.freq.set(f + 1, (this.freq.get(f + 1) || 0) + 1);
    }

    deleteOne(number: number): void {
        const f = this.cnt.get(number) || 0;
        if (f === 0) {
            return;
        }
        this.freq.set(f, (this.freq.get(f) || 0) - 1);
        this.cnt.set(number, f - 1);
        this.freq.set(f - 1, (this.freq.get(f - 1) || 0) + 1);
    }

    hasFrequency(frequency: number): boolean {
        return (this.freq.get(frequency) || 0) > 0;
    }
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * var obj = new FrequencyTracker()
 * obj.add(number)
 * obj.deleteOne(number)
 * var param_3 = obj.hasFrequency(frequency)
 */
```

### **...**

```

```

<!-- tabs:end -->
