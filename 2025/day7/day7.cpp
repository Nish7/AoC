#include <chrono>
#include <cstddef>
#include <fstream>
#include <iostream>
#include <map>
#include <queue>
#include <set>
#include <stdlib.h>
#include <thread>
#include <utility>
#include <vector>
using namespace std;

class Day7 {
public:
  int rows;
  int cols;
  vector<vector<char>> grid;

  void printGrid(vector<vector<char>> grid) {
    std::cout << "\033[2J\033[1;1H";
    cout << flush;
    for (auto g : grid) {
      for (auto c : g)
        cout << c << "";
      cout << "\n";
    }

    // std::this_thread::sleep_for(std::chrono::milliseconds(100));
  }

  vector<vector<char>> parseGrid(ifstream &ifile) {
    vector<vector<char>> res;
    string line;
    while (getline(ifile, line)) {
      vector<char> e;
      for (auto c : line)
        e.push_back(c);

      res.push_back(e);
    }

    return res;
  }

  pair<int, int> findStartingPoint(vector<vector<char>> grid) {
    pair<int, int> res;
    for (auto i = 0; i < grid.size(); i++) {
      for (auto j = 0; j < grid[0].size(); j++) {
        if (grid[i][j] == 'S') {
          res.first = i;
          res.second = j;
          return res;
        }
      }
    }

    return res;
  }

  int numberOfSplits(vector<vector<char>> grid) {
    auto [si, sj] = Day7::findStartingPoint(grid);
    int count = 0;

    set<pair<int, int>> visited_set;
    queue<pair<int, int>> q;

    q.emplace(si + 1, sj);

    while (!q.empty()) {
      auto [i, j] = q.front();
      grid[i][j] = '|';
      q.pop();
      // printGrid(grid);

      if (i + 1 < grid.size() && grid[i + 1][j] == '^') {
        count++;

        // split left
        if (i + 1 < grid.size() && j - 1 >= 0 && j - 1 < grid[0].size() &&
            !visited_set.contains({i + 1, j - 1})) {
          visited_set.insert({i + 1, j - 1});
          q.push({i + 1, j - 1});
        }

        // split right
        if (i + 1 < grid.size() && j + 1 >= 0 && j + 1 < grid[0].size() &&
            !visited_set.contains({i + 1, j + 1})) {
          visited_set.insert({i + 1, j + 1});
          q.push({i + 1, j + 1});
        }
      } else if (i + 1 < grid.size() && !visited_set.contains({i + 1, j})) {
        // go straight
        visited_set.insert({i + 1, j});
        q.push({i + 1, j});
      }
    }

    return count;
  }

  long long processBeam(int i, int j, map<pair<int, int>, long long> &memo) {
    if (i < 0 || j < 0 || i >= this->rows || j >= this->cols) {
      return 1;
    }

    if (memo[{i, j}])
      return memo[{i, j}];

    if (grid[i][j] == '^') {
      long long cnt =
          processBeam(i + 1, j - 1, memo) + processBeam(i + 1, j + 1, memo);
      memo[{i, j}] = cnt;
      return cnt;
    }

    this->grid[i][j] = '|';
    long long cnt = processBeam(i + 1, j, memo);
    memo[{i, j}] = cnt;
    return cnt;
  }

  long long getLifetime(vector<vector<char>> &grid) {
    this->grid = grid;
    this->rows = grid.size();
    this->cols = grid[0].size();

    map<pair<int, int>, long long> memo;

    auto [si, sj] = Day7::findStartingPoint(grid);
    return processBeam(si + 1, sj, memo);
  }
};

int main() {
  // ifstream test("test.txt");
  ifstream input("input.txt");
  Day7 day7;
  auto grid = day7.parseGrid(input);

  // day7.printGrid(grid);
  // auto p1 = day7.numberOfSplits(grid);
  cout << day7.getLifetime(grid);
  return 0;
}
