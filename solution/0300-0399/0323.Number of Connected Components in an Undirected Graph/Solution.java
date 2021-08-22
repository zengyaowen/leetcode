class Solution {
    private int[] p;
    
    public int countComponents(int n, int[][] edges) {
        p = new int[n];
        for (int i = 0; i < n; ++i) {
            p[i] = i;
        }
        for (int[] e : edges) {
            int a = e[0], b = e[1];
            p[find(b)] = find(a);
        }

        int cnt = 0;
        boolean[] visit = new boolean[n];
        for (int i = 0; i < n; ++i) {
            if (!visit[find(i)]) {
                ++cnt;
                visit[find(i)] = true;
            }
        }
        return cnt;
    }

    private int find(int x) {
        if (p[x] != x) {
            p[x] = find(p[x]);
        }
        return p[x];
    }
}