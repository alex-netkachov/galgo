#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

class Edge
{
public:
    Edge(int src, int dst, int weight) : src(src), dst(dst), weight(weight)
    {
    }

    int getSrc()
    {
        return src;
    }

    int getDst()
    {
        return dst;
    }

private:
    int src;
    int dst;
    int weight;
};

class Node
{
public:
    Node(int id) : id(id)
    {
    }

    int getId()
    {
        return id;
    }

    vector<Edge> getEdges()
    {
        return edges;
    }

    void addEdge(int dst, int weight)
    {
        Edge e(id, dst, weight);
        edges.push_back(e);
    }

private:
    int id;
    vector<Edge> edges;
};


int main()
{
    int n, m;
    cin >> n;
    cin >> m;

    vector<Node> nodes;
    for (auto i = 0; i < n; i++) {
        Node node(i);
        nodes.push_back(node);
    }

    vector<Edge> edge;
    for (auto i = 0; i < m; i++) {
        int src, dst, weight;
        cin >> src;
        cin >> dst;
        cin >> weight;

        src--;
        dst--;

        nodes[src].addEdge(dst, weight);
        nodes[dst].addEdge(src, weight);
    }

    int start;
    cin >> start;

    bool visited[nodes.size()];
    visited[start] = true;
    vector<int> queue;
    queue.push_back(start);
    for (auto i = 0; i < queue.size(); i++)
    {
        for (auto e : nodes[queue[i]].getEdges())
        {
            if (visited[e.getDst()])
            {
                continue;
            }

            visited[e.getDst()] = false;
            queue.push_back(e.getDst());
        }
    }

    cout << m << " " << n << endl;

    return 0;
}
