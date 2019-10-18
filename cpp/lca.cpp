#include <iostream>
#include <algorithm>
#include <cstring>
#include <queue>
#include <vector>
using namespace std;

const int maxn = 10005;
int parents[maxn][20], depth[maxn];
int n, from[maxn], root = -1;
vector<int> G[maxn];

void init()
{
    memset(parents, -1, sizeof(parents));
    memset(from, -1, sizeof(from));
    memset(depth, -1, sizeof(depth));
}

void getData()
{
    cin >> n;
    int u, v;
    for (int i = 1; i < n; ++i)
    {
        cin >> u >> v;
        G[u].push_back(v);
        parents[v][0] = u;
        from[v] = 1;
    }
    for (int i = 1; i <= n; ++i)
    {
        if (from[i] == -1)
            root = i;
    }
}

void getDepth_dfs(int u)
{
    int len = G[u].size();
    for (int i = 0; i < len; ++i)
    {
        int v = G[u][i];
        depth[v] = depth[u] + 1;
        getDepth_dfs(v);
    }
}

void getDepth_bfs(int u)
{
    queue<int> Q;
    Q.push(u);
    while (!Q.empty())
    {
        int v = Q.front();
        Q.pop();
        for (int i = 0; i < G[v].size(); ++i)
        {
            depth[G[v][i]] = depth[v] + 1;
            Q.push(G[v][i]);
        }
    }
}

void getParents()
{
    for (int up = 1; (1 << up) <= n; ++up)
    {
        for (int i = 1; i <= n; ++i)
        {
            parents[i][up] = parents[parents[i][up - 1]][up - 1];
        }
    }
}

int Lca(int u, int v)
{
    if (depth[u] < depth[v])
        swap(u, v);
    int i = -1, j;
    while ((1 << (i + 1)) <= depth[u])
        ++i;
    for (j = i; j >= 0; --j)
    {
        if (depth[u] - (1 << j) >= depth[v])
        {
            u = parents[u][j];
        }
    }
    if (u == v)
        return u;
    for (j = i; j >= 0; --j)
    {
        if (parents[u][j] != parents[v][j])
        {
            u = parents[u][j];
            v = parents[v][j];
        }
    }
    return parents[u][0];
}

void questions()
{
    int q, u, v;
    cin >> q;
    for (int i = 0; i < q; ++i)
    {
        cin >> u >> v;
        int ans = Lca(u, v);
        cout << ans << endl;
        //cout << u << " 和 " << v << " 的最近公共祖先(LCA)是: " << ans << endl;
    }
}

int main()
{
    init();
    getData();
    depth[root] = 1;
    getDepth_dfs(root);
    //getDepth_bfs(root);
    getParents();
    questions();
}
