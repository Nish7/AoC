#include <cfloat>
#include <fstream>
#include <iostream>
#include <map>
#include <queue>
#include <ranges>
#include <set>
#include <unordered_map>
#include <vector>

using namespace std;

namespace Day11 {
map<string, vector<string>> parseInput(ifstream &file) {
  map<string, vector<string>> res;
  string line;

  while (getline(file, line)) {
    int semi = line.find(':');
    string key = line.substr(0, semi);
    string vstring = line.substr(semi + 1);

    auto view = vstring | views::split(' ');
    vector<string> values;
    for (const auto &v : view) {
      auto s = string(v.begin(), v.end());
      if (s == "")
        continue;
      values.push_back(s);
    }

    res[key] = values;
  }

  return res;
}

void dfspaths(string k, set<string> &path, map<string, vector<string>> &mp,
              int &cnt) {
  if (k == "out") {
    for (auto c : path) {
      cout << c << ",";
    }
    cout << "\n";

    if (path.contains("dac") && path.contains("fft")) {
      cout << cnt << endl;
      cnt++;
    }

    return;
  }

  for (auto nbr : mp[k]) {
    if (path.contains(nbr))
      continue;
    path.insert(nbr);
    dfspaths(nbr, path, mp, cnt);
    path.erase(nbr);
  }
}

int getPathsWithDACFFT(map<string, vector<string>> mp) {
  int cnt = 0;
  set<string> path;
  path.insert("svr");
  dfspaths("svr", path, mp, cnt);
  return cnt;
}

int getPaths(map<string, vector<string>> mp) {
  queue<string> q;
  q.push("you");
  int paths = 0;

  while (!q.empty()) {
    auto v = q.front();
    q.pop();

    for (auto nbr : mp[v]) {
      if (nbr == "out") {
        paths++;
        continue;
      }

      q.push(nbr);
    }
  }

  return paths;
}

vector<string> topoSort(map<string, vector<string>> &m) {
  vector<string> res;
  unordered_map<string, long long> indegree;
  for (auto k : m)
    indegree[k.first] = 0;

  for (auto k : m) {
    for (auto n : m[k.first]) {
      if (!indegree.count(n))
        indegree[n] = 0;
      indegree[n]++;
    }
  }

  queue<string> q;
  for (auto k : indegree) {
    if (k.second == 0)
      q.push(k.first);
  }

  while (!q.empty()) {
    auto v = q.front();
    q.pop();

    res.push_back(v);
    for (auto nbr : m[v]) {
      indegree[nbr]--;
      if (indegree[nbr] == 0) {
        q.push(nbr);
      }
    }
  }

  return res;
}

long long getPathsDP(map<string, vector<string>> &m, string start, string end) {
  auto topo = topoSort(m);
  unordered_map<string, long long> count;
  int stidx = -1;
  for (int i = 0; i < topo.size(); i++) {
    auto k = topo[i];
    if (k == start) {
      count[k] = 1;
      stidx = i;
      continue;
    }
    count[k] = 0;
  }

  if (stidx == -1)
    return 0;

  for (int i = stidx; i < topo.size(); i++) {
    for (auto nbr : m[topo[i]]) {
      count[nbr] += count[topo[i]];
    }
  }

  return count[end];
}
}; // namespace Day11

int main() {
  // ifstream input("test.txt");
  // ifstream input("test2.txt");
  ifstream input("input.txt");

  auto inputmap = Day11::parseInput(input);

  /*for (auto k : inputmap) {*/
  /*  cout << k.first << " ";*/
  /*  cout << k.second.size() << " ";*/
  /*  cout << "[";*/
  /*  for (auto v : k.second) {*/
  /*    cout << v << " ";*/
  /*  }*/
  /*  cout << "]" << endl;*/
  /*}*
   */

  // cout << Day11::getPaths(inputmap) << endl;
  /*cout << Day11::getPathsWithDACFFT(inputmap) << endl;*/

  /*auto v = Day11::topoSort(inputmap);*/
  auto p1 = Day11::getPathsDP(inputmap, "svr", "fft");
  auto p2 = Day11::getPathsDP(inputmap, "fft", "dac");
  auto p3 = Day11::getPathsDP(inputmap, "dac", "out");
  long long a = p1 * p2 * p3;
  cout << p1 << " " << p2 << " " << p3 << endl;
  cout << a << endl;
  return 0;
}
