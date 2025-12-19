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
  return (abs(a.x - b.x) + 1) * (abs(a.y - b.y) + 1);
}

// not used
int orientation(Entry a, Entry b, Entry c) {
  long long v = (long long)(b.x - a.x) * (c.y - a.y) -
                (long long)(b.y - a.y) * (c.x - a.x);

  if (v > 0)
    return 1;
  if (v < 0)
    return -1;
  return 0;
}

// not used
bool isSegmentIntersect(Segment a, Segment b) {
  int o1 = orientation(a.first, a.second, b.first);
  int o2 = orientation(a.first, a.second, b.second);
  int o3 = orientation(b.first, b.second, a.first);
  int o4 = orientation(b.first, b.second, a.second);

  return o1 * o2 < 0 && o3 * o4 < 0;
}

bool edgeoverlap(const Segment &edge, int innerminx, int innermaxx,
                 int innerminy, int innermaxy) {
  auto p1 = edge.first;
  auto p2 = edge.second;

  if (p1.y == p2.y) {
    if (p1.y < innerminy || p1.y > innermaxy)
      return false;

    int xleft = min(p1.x, p2.x);
    int xright = max(p1.x, p2.x);
    return max(xleft, innerminx) <= min(xright, innermaxx);
  } else {
    if (p1.x < innerminx || p1.x > innermaxx)
      return false;
    int ybottom = min(p1.y, p2.y);
    int ytop = max(p1.y, p2.y);
    return max(ybottom, innerminy) <= min(ytop, innermaxy);
  }
}

bool isValidRect(Entry a, Entry b, vector<Segment> &polyedges) {
  int maxx = max(a.x, b.x);
  int minx = min(a.x, b.x);
  int maxy = max(a.y, b.y);
  int miny = min(a.y, b.y);

  int innerminx = minx + 1;
  int innermaxx = maxx - 1;
  int innerminy = miny + 1;
  int innermaxy = maxy - 1;

  if (innermaxx < innerminx || innermaxy < innerminy)
    return true;

  for (const auto &edge : polyedges) {
    if (edgeoverlap(edge, innerminx, innermaxx, innerminy, innermaxy))
      return false;
  }

  return true;
}

long long getInsideLargestArea(vector<Entry> &points) {
  vector<pair<long long, pair<Entry, Entry>>> areas;
  for (int i = 0; i < points.size(); ++i) {
    for (int j = i + 1; j < points.size(); ++j) {
      areas.push_back({area(points[i], points[j]), {points[i], points[j]}});
    }
  }

  // get larger area points first
  sort(areas.begin(), areas.end(),
       [](const auto &a, const auto &b) { return a.first > b.first; });

  auto polyedges = getPolygonEdges(points);

  for (auto a : areas) {
    Entry p1 = a.second.first;
    Entry p2 = a.second.second;

    if (isValidRect(p1, p2, polyedges))
      return a.first;
  }

  return -1;
}

} // namespace Day9

int main() {
  // ifstream input("test.txt");
  ifstream input("input.txt");
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
