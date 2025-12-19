#include <cfloat>
#include <fstream>
#include <iostream>
#include <queue>
#include <sstream>
#include <string>
#include <unordered_map>
#include <vector>
#include <z3++.h>

using namespace std;
using namespace z3;

struct PairHash {
  size_t operator()(const pair<vector<bool>, int> &k) const noexcept {
    size_t h = std::hash<int>{}(k.second);
    for (bool b : k.first) {
      h ^= std::hash<bool>{}(b) + 0x9e3779b97f4a7c15ULL + (h << 6) + (h >> 2);
    }
    return h;
  }
};

struct VecHash {
  size_t operator()(const vector<int> &v) const noexcept {
    size_t h = 1469598103934665603ull; // FNV offset
    for (int x : v) {
      h ^= static_cast<size_t>(x);
      h *= 1099511628211ull; // FNV prime
    }
    return h;
  }
};

namespace Day10 {

class Machine {
public:
  vector<bool> goalstate;
  vector<vector<int>> buttons;
  vector<int> joltage;
  unordered_map<pair<vector<bool>, int>, int, PairHash> memo;
  unordered_map<vector<int>, int, VecHash> memop2;
  int bound = 10;

  void printState() {
    cout << "[";
    for (auto v : goalstate) {
      cout << v << " ";
    }
    cout << "]";
  }

  void printState(vector<bool> state) {
    cout << "[";
    for (auto v : state) {
      cout << v << " ";
    }
    cout << "]";
  }

  void printState(vector<int> state) {
    cout << "[";
    for (auto v : state) {
      cout << v << " ";
    }
    cout << "]";
  }

  void printJoltage() {
    cout << "[";
    for (auto v : joltage) {
      cout << v << " ";
    }
    cout << "]";
  }

  void printButton() {
    cout << " ";
    for (const auto &b : this->buttons) {
      cout << "(";
      for (size_t i = 0; i < b.size(); ++i) {
        cout << b[i];
        if (i + 1 < b.size())
          cout << ",";
      }
      cout << ") ";
    }
  }

  void printMemo() {
    cout << "Memo: \n";
    for (auto k : this->memo) {
      cout << "{";
      cout << "Curr State: ";
      printState(k.first.first);
      cout << "Last Press: " << k.first.second;
      cout << ": " << memo[k.first];
      cout << "}\n";
    }
  }

  void print() {
    this->printState();
    cout << " ";
    this->printButton();
    cout << " ";
    this->printJoltage();
  }

  vector<bool> pressButton(vector<bool> curr, int pressIdx) {
    auto nextstate = curr;
    for (auto i : this->buttons[pressIdx]) {
      nextstate[i] = !nextstate[i];
    }
    return nextstate;
  }

  int backtrackPresses(vector<bool> curr, int n, int lastPressIdx) {
    // cout << "------ " << n << "-----\n";
    // cout.flush();
    // cout << "curr state: ";
    // printState(curr);
    // cout << "- last press: " << lastPressIdx;
    // cout << "\n";
    // printMemo();

    if (curr == this->goalstate) {
      // cout << "Found!\n";
      return n;
    }

    if (memo.count({curr, n})) {
      return memo[{curr, n}];
    }

    if (n == this->bound)
      return INT_MAX;

    int minv = INT_MAX;
    for (int i = 0; i < this->buttons.size(); i++) {
      if (i == lastPressIdx)
        continue;

      minv = min(minv, backtrackPresses(pressButton(curr, i), n + 1, i));
    }

    memo[{curr, n}] = minv;
    return minv;
  }

  int getButtonPresses() {
    vector<bool> init_s(this->goalstate.size(), false);
    return backtrackPresses(init_s, 0, -1);
  }

  // -------  Part 2 -------
  vector<int> pressButton(const vector<int> &curr, int pressIdx) {
    auto nextstate = curr;
    for (auto i : this->buttons[pressIdx]) {
      nextstate[i] = nextstate[i] + 1;
    }
    return nextstate;
  }

  bool isWithinBounds(const vector<int> &curr) {
    for (int i = 0; i < curr.size(); i++) {
      if (curr[i] > this->joltage[i]) {
        return false;
      }
    }

    return true;
  }

  int backtrackJoltage(vector<int> curr) {
    if (curr == this->joltage) {
      return 0;
    }

    if (memop2.count(curr)) {
      return memop2[curr];
    }

    int minv = INT_MAX;

    for (int i = 0; i < this->buttons.size(); i++) {
      auto next = pressButton(curr, i);

      if (!isWithinBounds(next))
        continue;

      int sub = backtrackJoltage(next);
      if (sub != INT_MAX) {
        minv = min(minv, sub + 1);
      }
    }

    memop2[curr] = minv;
    return minv;
  }

  int getJoltagePresses() {
    memop2.clear();
    vector<int> init_s(this->joltage.size(), 0);
    return backtrackJoltage(init_s);
  }

  int getJoltagePressesBFS() {
    const vector<int> target = joltage;
    vector<int> start(target.size(), 0);

    auto within = [&](const vector<int> &s) {
      for (size_t i = 0; i < s.size(); i++)
        if (s[i] > target[i])
          return false;
      return true;
    };

    unordered_map<vector<int>, int, VecHash> dist;
    queue<vector<int>> q;

    dist[start] = 0;
    q.push(start);

    while (!q.empty()) {
      auto cur = q.front();
      q.pop();
      int d = dist[cur];
      if (cur == target)
        return d;

      for (int b = 0; b < (int)buttons.size(); b++) {
        auto nxt = cur;
        for (int idx : buttons[b])
          nxt[idx]++;

        if (!within(nxt))
          continue;
        if (dist.find(nxt) != dist.end())
          continue;

        dist[nxt] = d + 1;
        q.push(std::move(nxt));
      }
    }
    return INT_MAX; // unreachable
  }

  vector<z3::expr> buildEquations(z3::context &ctx, optimize &opt,
                                  vector<expr> &x) {
    // T = {3,5,4,7} -> target vector
    // (3) (1,3) (2) (2,3) (0,2) (0,1) --> buttons presses
    // Systems of linear equations:
    // (+ 0 x4 x5) = 3
    // (+ 0 x1 x5) = 5
    // (+ 0 x2 x3 x4) = 4
    // (+ 0 x0 x1 x3) = 7
    vector<z3::expr> counters;
    // init the counters with expr (0)
    for (int i = 0; i < this->joltage.size(); i++) {
      counters.push_back(ctx.int_val(0));
    }

    for (int j = 0; j < this->buttons.size(); j++) {
      auto b = buttons[j];
      for (int i = 0; i < b.size(); i++) {
        counters[b[i]] = counters[b[i]] + x[j];
      }
    }

    for (int i = 0; i < this->joltage.size(); i++) {
      opt.add(counters[i] == this->joltage[i]);
    }
    
    return counters;
  }

  int getJoltagePressesSAT() {
    z3::context ctx;
    optimize opt(ctx);

    // creates the variable
    vector<z3::expr> x;
    for (int i = 0; i < buttons.size(); i++) {
      x.push_back(ctx.int_const(("x" + to_string(i)).c_str()));
      opt.add(x[i] >= 0);
    }

    // get the equations
    auto eqs = buildEquations(ctx, opt, x);

    // minimize
    expr total = ctx.int_val(0);
    for (int i = 0; i < buttons.size(); i++) {
      total = total + x[i];
    }

    opt.minimize(total);

    // eval
    int res = -1;
    if (opt.check() == z3::sat) {
      model m = opt.get_model();
      res = m.eval(total).get_numeral_int();
    }

    return res;
  }
};

Machine parseMachine(string s) {
  size_t i = 0;
  string state;
  vector<vector<int>> buttons;
  vector<int> joltage;

  while (i < s.size()) {
    if (isspace(s[i])) {
      i++;
      continue;
    }

    if (s[i] == '[') {
      size_t j = s.find(']', i);
      state = s.substr(i + 1, j - i - 1);
      i = j + 1;
    } else if (s[i] == '(') {
      size_t j = s.find(')', i);
      string inside = s.substr(i + 1, j - i - 1);

      vector<int> nums;
      stringstream ss(inside);
      string token;
      while (getline(ss, token, ',')) {
        nums.push_back(stoi(token));
      }

      buttons.push_back(nums);
      i = j + 1;
    } else if (s[i] == '{') {
      size_t j = s.find('}', i);
      string inside = s.substr(i + 1, j - i - 1);

      stringstream ss(inside);
      string token;
      while (getline(ss, token, ',')) {
        joltage.push_back(stoi(token));
      }

      i = j + 1;
    } else {
      i++;
    }
  }

  // parse state into
  vector<bool> statebool;
  for (auto c : state) {
    if (c == '.') {
      statebool.push_back(false);
    } else if (c == '#') {
      statebool.push_back(true);
    }
  }

  return Machine{statebool, buttons, joltage};
}

vector<Machine> getMachines(ifstream &ifile) {
  vector<Machine> machines;
  string line;
  while (getline(ifile, line)) {
    machines.push_back(parseMachine(line));
  }

  return machines;
}

int getTotalPresses(vector<Machine> machines) {
  int s = 0;
  for (auto m : machines) {
    // m.print();
    auto c = m.getButtonPresses();
    s += c;
    // cout << "\n";
  }

  return s;
}

int getTotalJoltage(vector<Machine> machines) {
  int s = 0;
  for (auto m : machines)
    s += m.getJoltagePressesSAT();
  return s;
}

}; // namespace Day10

int main() {
  ifstream input("test.txt");
  // ifstream input("input.txt");
  auto machines = Day10::getMachines(input);
  // cout << Day10::getTotalPresses(machines);
  cout << Day10::getTotalJoltage(machines) << endl;

  return 0;
}
