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
    // cout << c << " ";
    for (auto j = 1; j < v.size(); j++) {
      // cout << v[j] << " ";
     if (op == Ops::MUL) {
        c *= v[j];
        // cout << "*";
      } else if (op == Ops::PLUS) {
        c += v[j];
        // cout << "+";
      }
    }
    
    // cout << " =" << c << " | ";
    sum += c;
    // cout << " =" << sum << "\n";
  }

  return sum;
}
} // namespace Day6

int main() {
  ifstream tinput("input.txt");
  ifstream opsinput("input_op.txt");

  // ifstream tinput("test.txt");
  // ifstream opsinput("test_op.txt");

  auto vec = Day6::parseNumber(tinput);
  auto ops = Day6::parseOps(opsinput);
  auto transpo = Day6::transposeVec(vec);

  // for (auto e : transpo) {
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
  //
  // cout << transpo.size() << " " << ops.size() << "\n" << endl;

  auto p1 = Day6::doMaths(transpo, ops);
  cout << p1 << endl;
}
