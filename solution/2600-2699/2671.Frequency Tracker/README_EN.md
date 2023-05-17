# [2671. Frequency Tracker](https://leetcode.com/problems/frequency-tracker)

[中文文档](/solution/2600-2699/2671.Frequency%20Tracker/README.md)

## Description

<p>Design a data structure that keeps track of the values in it and answers some queries regarding their frequencies.</p>

<p>Implement the <code>FrequencyTracker</code> class.</p>

<ul>
	<li><code>FrequencyTracker()</code>: Initializes the <code>FrequencyTracker</code> object with an empty array initially.</li>
	<li><code>void add(int number)</code>: Adds <code>number</code> to the data structure.</li>
	<li><code>void deleteOne(int number)</code>: Deletes <strong>one</strong> occurence of <code>number</code> from the data structure. The data structure <strong>may not contain</strong> <code>number</code>, and in this case nothing is deleted.</li>
	<li><code>bool hasFrequency(int frequency)</code>: Returns <code>true</code> if there is a number in the data structure that occurs <code>frequency</code> number of times, otherwise, it returns <code>false</code>.</li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input</strong>
[&quot;FrequencyTracker&quot;, &quot;add&quot;, &quot;add&quot;, &quot;hasFrequency&quot;]
[[], [3], [3], [2]]
<strong>Output</strong>
[null, null, null, true]

<strong>Explanation</strong>
FrequencyTracker frequencyTracker = new FrequencyTracker();
frequencyTracker.add(3); // The data structure now contains [3]
frequencyTracker.add(3); // The data structure now contains [3, 3]
frequencyTracker.hasFrequency(2); // Returns true, because 3 occurs twice

</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input</strong>
[&quot;FrequencyTracker&quot;, &quot;add&quot;, &quot;deleteOne&quot;, &quot;hasFrequency&quot;]
[[], [1], [1], [1]]
<strong>Output</strong>
[null, null, null, false]

<strong>Explanation</strong>
FrequencyTracker frequencyTracker = new FrequencyTracker();
frequencyTracker.add(1); // The data structure now contains [1]
frequencyTracker.deleteOne(1); // The data structure becomes empty []
frequencyTracker.hasFrequency(1); // Returns false, because the data structure is empty

</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input</strong>
[&quot;FrequencyTracker&quot;, &quot;hasFrequency&quot;, &quot;add&quot;, &quot;hasFrequency&quot;]
[[], [2], [3], [1]]
<strong>Output</strong>
[null, false, null, true]

<strong>Explanation</strong>
FrequencyTracker frequencyTracker = new FrequencyTracker();
frequencyTracker.hasFrequency(2); // Returns false, because the data structure is empty
frequencyTracker.add(3); // The data structure now contains [3]
frequencyTracker.hasFrequency(1); // Returns true, because 3 occurs once

</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= number &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= frequency &lt;= 10<sup>5</sup></code></li>
	<li>At most, <code>2 *&nbsp;10<sup>5</sup></code>&nbsp;calls will be made to <code>add</code>, <code>deleteOne</code>, and <code>hasFrequency</code>&nbsp;in <strong>total</strong>.</li>
</ul>

## Solutions

<!-- tabs:start -->

### **Python3**

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
