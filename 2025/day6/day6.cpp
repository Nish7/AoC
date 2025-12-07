#include <cctype>
#include <fstream>
#include <iostream>
#include <ranges>
#include <string>
#include <vector>
using namespace std;

namespace Day6 {
vector<vector<int>> parseNumber(ifstream &file) {
  vector<vector<int>> r;
  string line;

  while (getline(file, line)) {
    string_view sv(line);
    auto split_sv = sv | ranges::views::split(' ');
    vector<int> re;
    for (const auto e : split_sv) {
      string s(e.begin(), e.end());
      if (s == "")
        continue;
      re.emplace_back(stoi(s));
    }

    r.push_back(re);
  }

  return r;
}

vector<vector<char>> parsegrid(ifstream &file) {
  vector<vector<char>> r;
  string line;

  while (getline(file, line)) {
    vector<char> re;
    for (const auto c : line) {
      re.push_back(c);
    }
    r.push_back(re);
  }

  return r;
}

enum Ops { PLUS, MUL };

vector<Ops> parseOps(ifstream &file) {
  vector<Ops> r;
  string line;
  getline(file, line);

  for (const auto e : line) {
    switch (e) {
    case ' ':
      continue;
    case '+':
      r.push_back(Ops::PLUS);
      break;
    case '*':
      r.push_back(Ops::MUL);
      break;
    default:
      break;
    }
  }

  return r;
}

vector<vector<int>> transposeVec(vector<vector<int>> vec) {
  // a(0,0) b(0,1) c(0,2)
  // d(1,0) e(1,1) f(1,2)
  //
  // a(0,0) d(1,0), 2,0
  // b(0,1) e
  // c(0,2) f
  //
  //
  vector<vector<int>> r(vec[0].size(), vector<int>(vec.size()));
  for (int cols = 0; cols < vec[0].size(); cols++) { // loop through the col
    for (int rows = 0; rows < vec.size(); rows++) {  // loop through the row
      r[cols][rows] = vec[rows][cols];
    }
  }

  return r;
}

long long doMaths(vector<vector<int>> num, vector<Ops> ops) {
  long long sum = 0;
  for (int i = 0; i < num.size(); i++) {
    auto op = ops[i];
    auto v = num[i];
    long long c = num[i][0];
    for (auto j = 1; j < v.size(); j++) {
      if (op == Ops::MUL) {
        c *= v[j];
      } else if (op == Ops::PLUS) {
        c += v[j];
      }
    }

    sum += c;
  }

  return sum;
}

long long byDigits(vector<vector<char>> charGrid, vector<Ops> ops) {
  vector<vector<long long>> res;
  vector<long long> newnenr;
  for (auto cols = 0; cols < charGrid[0].size(); cols++) {
    auto s = string(1, charGrid[0][cols]);
    bool allempty = s == " " ? true : false;
    string entry;

    for (auto rows = 0; rows < charGrid.size(); rows++) {
      const char v = charGrid[rows][cols];
      if (v == ' ') {
        continue;
      } else {
        allempty = false;
      }
      entry.push_back(v);
    }

    if (allempty) {
      res.push_back(newnenr);
      newnenr.clear();
    } else {
      newnenr.push_back(stoll(entry));
    }
  }
  
  res.push_back(newnenr);
  
  int opinx = 0;
  long long sum = 0;
  for (auto e : res) {
    auto op = ops[opinx];
    long long curr = op == Ops::MUL ? 1 : 0;
    for (auto v : e) {
      if (op == Ops::MUL) {
        curr *= v;
      } else if (op == Ops::PLUS) {
        curr += v;
      }
    }
    sum += curr;
    opinx++;
  }

  return sum;
}

} // namespace Day6

int main() {
  ifstream tinput("input.txt");
  ifstream opsinput("input_op.txt");

  // ifstream tinput("test.txt");
  // ifstream opsinput("test_op.txt");

  // auto vec = Day6::parseNumber(tinput);
  auto vec = Day6::parsegrid(tinput);
  auto ops = Day6::parseOps(opsinput);
  // auto transpo = Day6::transposeVec(vec);

  // for (auto e : vec) {
  //   for (auto er : e) {
  //     cout << er << " ";
  //   }
  //   cout << "\n";
  // }
  //
  // for (auto e : ops) {
  //   cout << e << " ";
  // }

  // cout << endl;
  //
  // cout << transpo.size() << " " << ops.size() << "\n" << endl;

  // auto p1 = Day6::doMaths(transpo, ops);
  // cout << p1 << endl;
  //

  auto p2 = Day6::byDigits(vec, ops);
  cout << p2 << endl;
}
