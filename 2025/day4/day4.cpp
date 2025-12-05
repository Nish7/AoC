#include "fstream"
#include "vector"
#include <algorithm>
#include <fstream>
#include <iostream>
using namespace std; 

namespace Day4 {
typedef vector<vector<char>> GridType;

void printGrid(GridType &grid) {
  for (const auto &e : grid) {
    for (auto c : e) {
      cout << c << "";
    }

    cout << "\n";
  }
  cout << endl;
}

GridType parseInput(ifstream &i) {
  vector<vector<char>> res;
  string line;
  while (getline(i, line)) {
    vector<char> entry;
    for (auto c : line)
      entry.push_back(c);
    res.push_back(entry);
  }
  return res;
}

bool processRoll(int i, int j, GridType &grid) { //
  if (i < 0 || j < 0 || i >= grid.size() || j >= grid[0].size())
    return false;
  if (grid[i][j] == '@' || grid[i][j] == 'x')
    return true;
  return false;
}

int getRollsOfPaper(GridType &grid) {
  int sum = 0;
  for (int i = 0; i < grid.size(); i++) {
    for (int j = 0; j < grid[i].size(); j++) {
      auto c = grid[i][j];
      if (c == '@') {
        int adjacentRols =
            processRoll(i - 1, j, grid) + processRoll(i + 1, j, grid) +
            processRoll(i, j - 1, grid) + processRoll(i, j + 1, grid) +
            processRoll(i + 1, j + 1, grid) + processRoll(i - 1, j - 1, grid) +
            processRoll(i - 1, j + 1, grid) + processRoll(i + 1, j - 1, grid);

        if (adjacentRols < 4) {
          grid[i][j] = 'x';
          sum++;
        }
      }
    }
  }

  return sum;
}

int collectRollPaper(GridType &grid) {
  int sum = 0;
  int val;

  while ((val = getRollsOfPaper(grid)) != 0) {
    sum += val;
    for (auto &e : grid) ranges::replace(e, 'x', '.');
  }

  return sum;
}

} // namespace Day4

int main() {
  ifstream ifs("test.txt");
  // ifstream ifs("input.txt");
  auto grid = Day4::parseInput(ifs);

  // auto t1 = Day4::getRollsOfPaper(grid);
  // cout << t1 << endl;

  auto t2 = Day4::collectRollPaper(grid);
  cout << t2 << endl;
}
