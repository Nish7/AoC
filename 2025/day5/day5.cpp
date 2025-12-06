#include <algorithm>
#include <fstream>
#include <iostream>
#include <set>
#include <string>
#include <sys/resource.h>
#include <vector>
using namespace std;

namespace Day5 {
using ll = long long;
using RangeVec = vector<pair<ll, ll>>;
using IngredientVec = vector<ll>;

RangeVec mergeIntervals(RangeVec &intervals) {
  sort(intervals.begin(), intervals.end(),
       [](pair<ll, ll> &a, pair<ll, ll> &b) { return a.first < b.first; });

  RangeVec ranges;
  ranges.push_back(intervals[0]);

  for (int i = 1; i < intervals.size(); i++) {
    auto [_, back_prev] = ranges.back();
    auto [start_i, end_i] = intervals[i];

    if (start_i <= back_prev) {
      ranges.back().second = max(end_i, back_prev);
    } else {
      ranges.emplace_back(start_i, end_i);
    }
  }

  return ranges;
}

IngredientVec parseIngredients(ifstream &ifile) {
  vector<ll> r;
  string line;
  while (getline(ifile, line))
    r.push_back(stoll(line));
  return r;
}

RangeVec parseRanges(ifstream &rfile) {
  RangeVec r;
  string line;

  while (getline(rfile, line)) {
    auto pos = line.find('-');
    auto a = stoll(line.substr(0, pos));
    auto b = stoll(line.substr(pos + 1));
    r.emplace_back(a, b);
  }

  return r;
}

// naive try. but loops are too big for this.
int getFreshUnOp(IngredientVec ingr, RangeVec rangeVec) {
  int freshThings = 0;
  std::set<ll> numset;
  rangeVec = Day5::mergeIntervals(rangeVec);

  for (auto r : rangeVec) {
    const auto [s, e] = r;
    for (auto i = s; i <= e; i++)
      numset.insert(i);
  }

  for (const auto e : ingr) {
    if (numset.contains(e))
      freshThings++;
  }

  return freshThings;
}

int getFresh(IngredientVec ingr, RangeVec rangeVec) {
  int freshThings = 0;
  rangeVec = mergeIntervals(rangeVec);

  for (const auto v : ingr) {
    for (const auto r : rangeVec) {
      const auto [s, e] = r;
      if ((v > s && v < e) || v == s || v == e) {
        freshThings++;
        break;
      } else if (s > v) {
        break;
      };
    }
  }

  return freshThings;
}

long long getTotalFreshCount(RangeVec ranges) {
  long long count = 0;
  for (auto r : mergeIntervals(ranges))
    count += r.second - r.first + 1;
  return count;
}

} // namespace Day5

int main() {
  ifstream ingredientsfile("test_ingredient.txt");
  ifstream rangesfile("test_ranges.txt");

  // ifstream ingredientsfile("input_ing.txt");
  // ifstream rangesfile("input_range.txt");

  auto ig = Day5::parseIngredients(ingredientsfile);
  auto ranges = Day5::parseRanges(rangesfile);

  // for (auto e : ig) cout << e << "\n";
  // cout << "ranges\n";
  // for (auto e : ranges) cout << e.first <<  " " << e.second << "\n";

  // auto p1 = Day5::getFresh(ig, ranges);
  // cout << p1 << endl;

  auto p2 = Day5::getTotalFreshCount(ranges);
  cout << p2 << endl;
}
