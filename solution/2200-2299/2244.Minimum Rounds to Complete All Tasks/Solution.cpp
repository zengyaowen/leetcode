class Solution {
public:
    int minimumRounds(vector<int>& tasks) {
        unordered_map<int, int> cnt;
        for (int& t : tasks) ++cnt[t];
        int ans = 0;
        for (auto& [_, v] : cnt)
        {
            if (v == 1) return -1;
            ans += v / 3 + (v % 3 == 0 ? 0 : 1);
        }
        return ans;
    }
};