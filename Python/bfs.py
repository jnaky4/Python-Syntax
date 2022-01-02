from collections import defaultdict


class Graph:
    def __init__(self):
        self.neighbors = defaultdict(list)

    def add_edge(self, u, v):
        self.neighbors[u].append(v)

    def bfs_recursion(self, v):
        visited = [False] * (max(self.neighbors) + 1)


g = Graph()
g.add_edge(0, 1)
g.add_edge(0, 2)
g.add_edge(1, 2)
g.add_edge(2, 0)
g.add_edge(2, 3)
g.add_edge(3, 3)

g.bfs_recursion(2)