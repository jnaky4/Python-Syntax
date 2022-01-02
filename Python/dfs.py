from collections import defaultdict


class Graph:

    def __init__(self):
        self.neighbors = defaultdict(list)

    def add_edge(self, u, v):
        self.neighbors[u].append(v)

    def dfs_recursion(self, v, visited):
        print(v, end=' ')
        visited.add(v)
        for neighbor in self.neighbors:
            if neighbor not in visited:
                self.dfs_recursion(neighbor, visited)

    def dfs(self, v):
        visited = set()
        self.dfs_recursion(v, visited)


g = Graph()

g.add_edge(0, 1)
g.add_edge(0, 2)
g.add_edge(1, 2)
g.add_edge(2, 0)
g.add_edge(2, 3)
g.add_edge(3, 3)

g.dfs(2)
