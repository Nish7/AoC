#include <algorithm>
#include <fstream>
#include <iostream>
#include <ranges>
#include <vector>

using namespace std;

struct Entry {
  int x;
  int y;
};

using Segment = pair<Entry, Entry>;

namespace Day9 {
void print(Entry e) { cout << "[" << e.x << "," << e.y << "]"; }

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
    res.push_back(e);
  }

  return res;
}

long long getLargestArea(vector<Entry> &input) {
  long long max = 0;
  for (int i = 0; i < input.size() - 1; i++) {
    Entry vali = input[i];
    for (int j = i + 1; j < input.size(); j++) {
      Entry valj = input[j];
      long long area = (vali.x - valj.x + 1) * (vali.y - valj.y + 1);
      if (area > max)
        max = area;
    }
  }

  return max;
}

pair<int, int> getMax(const vector<Entry> &points) {
  int maxwidth = points[0].x;
  int maxheight = points[0].y;

  for (const auto &p : points) {
    if (p.x > maxwidth)
      maxwidth = p.x;
    if (p.y > maxheight)
      maxheight = p.y;
  }

  return {maxwidth, maxheight};
}

vector<vector<char>> getGrid(vector<Entry> &points) {
  const auto [rows, cols] = getMax(points);
  vector<vector<char>> grid(rows + 1, vector<char>(cols + 1, '.'));
  for (const auto &p : points)
    grid[p.x][p.y] = '#';
  return grid;
}

vector<Segment> getRectEdges(Entry a, Entry b) {
  int x1 = min(a.x, b.x);
  int x2 = max(a.x, b.x);
  int y1 = min(a.y, b.y);
  int y2 = max(a.y, b.y);
  return {
      {Entry{x1, y1}, Entry{x2, y1}}, // bottom seg
      {Entry{x2, y1}, Entry{x2, y2}}, // right seg
      {Entry{x2, y2}, Entry{x1, y2}}, // top seg
      {Entry{x1, y2}, Entry{x1, y1}}, // left seg
  };
}

vector<Segment> getPolygonEdges(vector<Entry> &points) {
  // Assumption: points in the input are already "boundary-aligned"
  vector<Segment> res;
  for (int i = 0; i < points.size(); i++) {
    int j = (i + 1) % points.size();
    res.push_back(Segment{points[i], points[j]});
  }

  return res;
}

long long area(Entry a, Entry b) {
    return (long long)(max(a.x,b.x) - min(a.x,b.x)) *
           (long long)(max(a.y,b.y) - min(a.y,b.y));
}

int orientation(Entry a, Entry b, Entry c) {
  long long v = (long long)(b.x - a.x) * (c.y - a.y) -
                (long long)(b.y - a.y) * (c.x - a.x);

  if (v > 0)
    return 1;
  if (v < 0)
    return -1;
  return 0;
}

bool isSegmentIntersect(Segment a, Segment b) {
    int o1 = orientation(a.first, a.second, b.first);
    int o2 = orientation(a.first, a.second, b.second);
    int o3 = orientation(b.first, b.second, a.first);
    int o4 = orientation(b.first, b.second, a.second);

    return o1 * o2 < 0 && o3 * o4 < 0;
}

long long getInsideLargestArea(vector<Entry> &points) {
  vector<pair<long long, pair<Entry, Entry>>> areas;
  for (int i = 0; i < points.size() - 1; i++) {
    for (int j = i + 1; j < points.size(); j++) {
      areas.push_back({area(points[i], points[j]), {points[i], points[j]}});
    }
  }

  // get larger area points first
  sort(areas.begin(), areas.end(),
       [](const auto &a, const auto &b) { return a.first > b.first; });

  auto polyedges = getPolygonEdges(points);

  // for (auto p : polyedges) {
  //     cout << "---";
  //     print(p.first) ;
  //     cout << "--";
  //     print(p.second) ;
  //     cout << "\n";
  // }

  for (auto a : areas) {
    bool inters = false;
    auto rectedges = getRectEdges(a.second.first, a.second.second);
    cout << " area:" << a.first << "\n";
    cout << "p1:";
    print(a.second.first);
    cout << " --- p2:";
    print(a.second.second);
    cout << "rectage segments:\n";
    for (auto p : polyedges) {
      for (auto r : rectedges) {
        cout << "---";
        print(r.first);
        cout << "--";
        print(r.second);
        cout << "with ";
        print(p.first);
        cout << "--";
        print(p.second);
        cout << "\n";
        cout << "\n";
        if (isSegmentIntersect(p, r)) {
          inters = true;
          break;
        }
      }

      if (inters)
        break;
    }

    if (!inters)
      return a.first;
  }

  return -1;
}

} // namespace Day9

int main() {
  ifstream input("test.txt");
  // ifstream input("input.txt");
  // auto vec = Day8::parseInput(test);
  auto points = Day9::parseInput(input);
  // cout << Day9::getLargestArea(vec) << endl;
  cout << Day9::getInsideLargestArea(points) << endl;
  //
  //
  // for (auto c : points) {
  //   cout << c.x << " " << c.y << "\n";
  // }

  return 0;
}
