#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <queue>
#include <ranges>
#include <set>
#include <vector>

using namespace std;

struct Entry {
  long long x;
  long long y;
  long long z;

public:
};

bool operator<(const Entry &lhs, const Entry &rhs) {
  if (lhs.x != rhs.x)
    return lhs.x < rhs.x;
  if (lhs.y != rhs.y)
    return lhs.y < rhs.y;
  return lhs.z < rhs.z;
}

struct Node {
  long long distance;
  Entry a;
  Entry b;
};

struct Cmp {
  bool operator()(const Node &x, const Node &y) const {
    return x.distance > y.distance;
  }
};

namespace Day8 {
void print(Entry e) { cout << "[" << e.x << "," << e.y << "," << e.z << "]"; }

vector<Entry> parseInput(ifstream &ifile) {
  vector<Entry> res;
  string line;

  while (getline(ifile, line)) {
    string_view sv(line);
    auto split_v = sv | ranges::views::split(',');
    vector<int> t;
    for (const auto e : split_v) {
      t.push_back(stoi(string(e.begin(), e.end())));
    }

    Entry e;
    e.x = t[0];
    e.y = t[1];
    e.z = t[2];

    res.push_back(e);
  }

  return res;
}

int visitCircuit(Entry node, map<Entry, vector<Entry>> &mp,
                 set<Entry> &visited) {

  if (visited.contains(node))
    return 0;

  visited.insert(node);
  int num = 1;
  for (auto v : mp[node]) {
    num += visitCircuit(v, mp, visited);
  }

  return num;
}

pair<vector<int>, int> findSizeOfCircuits(int topKClosest,
                                          vector<Entry> &inputVec) {

  priority_queue<Node, vector<Node>, Cmp> pq;
  for (int i = 0; i < inputVec.size() - 1; i++) {
    const auto a = inputVec[i];
    for (int j = i + 1; j < inputVec.size(); j++) {
      const auto b = inputVec[j];
      long long dx = b.x - a.x;
      long long dy = b.y - a.y;
      long long dz = b.z - a.z;
      long long dist = dx * dx + dy * dy + dz * dz;
      pq.push({dist, a, b});
    }
  }

  map<Entry, vector<Entry>> adj;
  long long lastx = 0;
  for (int i = 0; i < topKClosest; i++) {
    const auto v = pq.top();
    lastx = v.b.x * v.a.x;

    pq.pop();
    if (!adj.count(v.b))
      adj[v.b] = {};
    if (!adj.count(v.a))
      adj[v.a] = {};
    adj[v.b].push_back(v.a);
    adj[v.a].push_back(v.b);
  }

  set<Entry> visited;
  vector<int> circuitsize;
  for (auto k : adj) {
    if (visited.contains(k.first))
      continue;

    circuitsize.push_back(visitCircuit(k.first, adj, visited));
  }

  sort(circuitsize.begin(), circuitsize.end(),
       [](int a, int b) { return a > b; });

  return {circuitsize, lastx};
}

int calculateTopK(int topKSize, vector<int> circuitsize) {
  int s = 1;
  for (auto i = 0; i < topKSize; i++) {
    if (circuitsize[i] == 0)
      continue;

    s *= circuitsize[i];
  }
  return s;
}

int returnLongest(vector<Entry> &input, int length) {
  int i = length;
  int j = -1;
  int ans = -1;

  while (length != j) {
    auto v = findSizeOfCircuits(i, input);
    j = v.first[0];
    ans = v.second;
    i++;
  }

  return ans;
}

} // namespace Day8

int main() {
  // ifstream input("test.txt");
  ifstream input("input.txt");
  // auto vec = Day8::parseInput(test);
  auto vec = Day8::parseInput(input);

  // auto v = Day8::findSizeOfCircuits(3919, vec);
  // for (auto c : v)
  //   cout << c << " ";
  // cout << Day8::calculateTopK(3, v) << endl;
  cout << Day8::returnLongest(vec, 1000);

  return 0;
}
