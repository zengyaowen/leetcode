class Solution:
    def isBipartite(self, graph: List[List[int]]) -> bool:
        n = len(graph)
        p = list(range(n))

        def find(x):
            if p[x] != x:
                p[x] = find(p[x])
            return p[x]

        for u, g in enumerate(graph):
            for v in g:
                if find(u) == find(v):
                    return False
                p[find(v)] = find(g[0])
        return True
