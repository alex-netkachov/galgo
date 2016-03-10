#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

class Edge
{
public:
    Edge(int node1, int node2, int weight) : node1(node1), node2(node2), weight(weight)
    {
    }

    int getNode1()
    {
        return node1;
    }

    int getNode2()
    {
        return node2;
    }

    int getOtherNode(int node)
    {
        return node1 == node ? node2 : node1;
    }

    int getWeight()
    {
        return weight;
    }

private:
    int node1;
    int node2;
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

    vector<int> getEdges()
    {
        return edges;
    }

    void addEdge(int id)
    {
        edges.push_back(id);
    }

private:
    int id;
    vector<int> edges;
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

    vector<Edge> edges;

    for (auto i = 0; i < m; i++) {
        int node1, node2, weight, id;
        cin >> node1;
        cin >> node2;
        cin >> weight;

        node1--;
        node2--;

        Edge e(node1, node2, weight);
        id = edges.size();
        edges.push_back(e);

        nodes[node1].addEdge(id);
        nodes[node2].addEdge(id);
    }

    int start;
    cin >> start;

    start--;

    vector<Edge> selectedEdges;

    bool visitedNodes[nodes.size()];
    for (auto i = 0; i < nodes.size(); i++) {
        visitedNodes[i] = false;
    }

    visitedNodes[start] = true;

    vector<Edge> queue;
    for (auto eid : nodes[start].getEdges())
    {
        Edge e(start, edges[eid].getOtherNode(start), edges[eid].getWeight());
        queue.push_back(e);
    }

    for (auto i = 0; i < queue.size(); i++)
    {
        auto min = 1000000;
        auto minIdx = -1;
        for (auto j = 0; j < queue.size(); j++)
        {
            auto e = queue[j];

            if (!visitedNodes[e.getNode2()] && e.getWeight() < min)
            {
                min = e.getWeight();
                minIdx = j;
            }
        }

        if (minIdx == -1)
        {
            break;
        }

        {
            auto e = queue[minIdx];
            Edge eg(e.getNode1(), e.getNode2(), e.getWeight());
            selectedEdges.push_back(eg);
            visitedNodes[e.getNode2()] = true;

            for (auto eid : nodes[e.getNode2()].getEdges())
            {
                Edge eg(e.getNode2(), edges[eid].getOtherNode(e.getNode2()), edges[eid].getWeight());
                if (!visitedNodes[eg.getNode2()]) {
                    queue.push_back(eg);
                }
            }
        }
    }

    auto weight = 0;
    for (auto e : selectedEdges)
    {
        weight += e.getWeight();
    }

    cout << weight << endl;

    return 0;
}
