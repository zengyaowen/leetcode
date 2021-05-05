# [381. Insert Delete GetRandom O(1) - Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed)

[中文文档](/solution/0300-0399/0381.Insert%20Delete%20GetRandom%20O%281%29%20-%20Duplicates%20allowed/README.md)

## Description

<p>Implement the <code>RandomizedCollection</code> class:</p>

<ul>
	<li><code>RandomizedCollection()</code> Initializes the <code>RandomizedCollection</code> object.</li>
	<li><code>bool insert(int val)</code> Inserts an item <code>val</code> into the multiset if not present. Returns <code>true</code> if the item was not present, <code>false</code> otherwise.</li>
	<li><code>bool remove(int val)</code> Removes an item <code>val</code> from the multiset if present. Returns <code>true</code> if the item was present, <code>false</code> otherwise. Note that if <code>val</code> has multiple occurrences in the multiset, we only remove one of them.</li>
	<li><code>int getRandom()</code> Returns a random element from the current multiset of elements (it&#39;s guaranteed that at least one element exists when this method is called). The probability of each element being returned is <strong>linearly related</strong> to the number of same values the multiset contains.</li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input</strong>
[&quot;RandomizedCollection&quot;, &quot;insert&quot;, &quot;insert&quot;, &quot;insert&quot;, &quot;getRandom&quot;, &quot;remove&quot;, &quot;getRandom&quot;]
[[], [1], [1], [2], [], [1], []]
<strong>Output</strong>
[null, true, false, true, 2, true, 1]

<strong>Explanation</strong>
RandomizedCollection randomizedCollection = new RandomizedCollection();
randomizedCollection.insert(1);   // return True. Inserts 1 to the collection. Returns true as the collection did not contain 1.
randomizedCollection.insert(1);   // return False. Inserts another 1 to the collection. Returns false as the collection contained 1. Collection now contains [1,1].
randomizedCollection.insert(2);   // return True. Inserts 2 to the collection, returns true. Collection now contains [1,1,2].
randomizedCollection.getRandom(); // getRandom should return 1 with the probability 2/3, and returns 2 with the probability 1/3.
randomizedCollection.remove(1);   // return True. Removes 1 from the collection, returns true. Collection now contains [1,2].
randomizedCollection.getRandom(); // getRandom should return 1 and 2 both equally likely.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= val &lt;= 2<sup>31</sup> - 1</code></li>
	<li>At most <code>10<sup>5</sup></code> calls will be made to <code>insert</code>, <code>remove</code>, and <code>getRandom</code>.</li>
	<li>There will be <strong>at least one</strong> element in the data structure when <code>getRandom</code> is called.</li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you implement the functions of the class with each function works in <strong>average</strong> <code>O(1)</code> time?

## Solutions

<!-- tabs:start -->

### **Python3**

```python
class RandomizedCollection:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.m = {}
        self.l = []

    def insert(self, val: int) -> bool:
        """
        Inserts a value to the collection. Returns true if the collection did not already contain the specified element.
        """
        idx_set = self.m.get(val, set())
        idx_set.add(len(self.l))
        self.m[val] = idx_set
        self.l.append(val)
        return len(idx_set) == 1

    def remove(self, val: int) -> bool:
        """
        Removes a value from the collection. Returns true if the collection contained the specified element.
        """
        if val not in self.m:
            return False
        idx_set = self.m[val]
        idx = list(idx_set)[0]
        last_idx = len(self.l) - 1
        self.l[idx] = self.l[last_idx]
        idx_set.remove(idx)

        last_idx_set = self.m[self.l[last_idx]]
        if last_idx in last_idx_set:
            last_idx_set.remove(last_idx)
        if idx < last_idx:
            last_idx_set.add(idx)
        if not idx_set:
            self.m.pop(val)
        self.l.pop()
        return True

    def getRandom(self) -> int:
        """
        Get a random element from the collection.
        """
        return -1 if len(self.l) == 0 else random.choice(self.l)



# Your RandomizedCollection object will be instantiated and called as such:
# obj = RandomizedCollection()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
```

### **Java**

```java
class RandomizedCollection {
    private Map<Integer, Set<Integer>> m;
    private List<Integer> l;
    private Random rnd;

    /** Initialize your data structure here. */
    public RandomizedCollection() {
        m = new HashMap<>();
        l = new ArrayList<>();
        rnd = new Random();
    }

    /** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
    public boolean insert(int val) {
        Set<Integer> idxSet = m.getOrDefault(val, new HashSet<>());
        idxSet.add(l.size());
        m.put(val, idxSet);
        l.add(val);
        return idxSet.size() == 1;
    }

    /** Removes a value from the collection. Returns true if the collection contained the specified element. */
    public boolean remove(int val) {
        if (!m.containsKey(val)) {
            return false;
        }
        Set<Integer> idxSet = m.get(val);
        int idx = idxSet.iterator().next();
        int lastIdx = l.size() - 1;
        l.set(idx, l.get(lastIdx));
        idxSet.remove(idx);

        Set<Integer> lastIdxSet = m.get(l.get(lastIdx));
        lastIdxSet.remove(lastIdx);
        if (idx < lastIdx) {
            lastIdxSet.add(idx);
        }
        if (idxSet.isEmpty()) {
            m.remove(val);
        }
        l.remove(lastIdx);
        return true;
    }

    /** Get a random element from the collection. */
    public int getRandom() {
        int size = l.size();
        return size == 0 ? -1 : l.get(rnd.nextInt(size));
    }
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * RandomizedCollection obj = new RandomizedCollection();
 * boolean param_1 = obj.insert(val);
 * boolean param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */
```

### **...**

```

```

<!-- tabs:end -->
