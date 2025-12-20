#include <fstream>
#include <iostream>
#include <ranges>
#include <vector>

using namespace std;

struct Entry {
  int width;
  int height;
  vector<int> shapes;
};

namespace Day12 {
vector<Entry> parseinput(ifstream &file) {
  string line;
  vector<Entry> res;

  while (getline(file, line)) {
    int colonidx = line.find(":");
    string wh = line.substr(0, colonidx);
    int xidx = wh.find("x");
    int w = stoi(line.substr(0, xidx));
    int h = stoi(line.substr(xidx + 1));
    string linenumber = line.substr(colonidx + 2);
    auto sv = linenumber | ranges::views::split(' ');
    vector<int> shapes;
    for (auto t : sv) {
      shapes.push_back(stoi(string(t.begin(), t.end())));
    }

    res.emplace_back(Entry{w, h, shapes});
  }

  return res;
}
int getPossibleRegions(vector<Entry> input) {
  int s = 0;
  for (auto i : input) {
    int area = i.width * i.height;
    int total = 0;
    for (auto v : i.shapes)
      total += v;
    if (area >= total * 9)
      s++;
  }

  return s;
}
}; // namespace Day12

int main() {
  ifstream input("input.txt");
  auto vec = Day12::parseinput(input);

  // for (auto c : vec) {
  //   cout << c.width << " " << c.height << " [";
  //   for (auto f : c.shapes) {
  //     cout << f << " ";
  //   }
  //   cout << "]" << endl;
  // }
  //
  cout << Day12::getPossibleRegions(vec) << endl;

  return 0;
}
