#include "iostream"
#include "vector"
#include <fstream>
#include <string>

class Solution {
public:
  long long max;
  std::vector<std::string> parse(std::ifstream &file) {
    std::vector<std::string> res;
    std::string line;

    while (std::getline(file, line)) {
      if (line.empty())
        continue;
      // long long in = std::stoll(line);
      res.push_back(line);
    }

    return res;
  };

  // Task 1:
  // 987654321111111
  // Assumption: all the lengths are the same
  // [a, b, c],
  // [d, b, c],
  // [e, b, c],
  //
  int task1(std::vector<std::string> input) {
    int sum = 0;
    for (auto s : input) {
      int max = 0;
      for (int i = 0; i < s.size() - 1; i++) {
        for (int j = i + 1; j < s.size(); j++) {
          std::string v_s = "";
          v_s += s[i];
          v_s += s[j];
          int v = std::stoi(v_s);
          if (v > max)
            max = v;
        }
      }

      sum += max;
    }

    return sum;
  }

  // Task 2:
  // LOLL!!!!!
  long long task2(std::vector<std::string> input) {
    long long sum = 0;
    for (auto s : input) {
      std::cout << s << "\n";
      long long max = 0;
      for (int i = 0; i < s.size() - 11; i++) {
        for (int j = i + 1; j < s.size() - 10; j++) {
          for (int k = j + 1; k < s.size() - 9; k++) {
            for (int l = k + 1; l < s.size() - 8; l++) {
              for (int m = l + 1; m < s.size() - 7; m++) {
                for (int n = m + 1; n < s.size() - 6; n++) {
                  for (int o = n + 1; o < s.size() - 5; o++) {
                    for (int p = o + 1; p < s.size() - 4; p++) {
                      for (int q = p + 1; q < s.size() - 3; q++) {
                        for (int r = q + 1; r < s.size() - 2; r++) {
                          for (int t = r + 1; t < s.size() - 1; t++) {
                            for (int u = t + 1; u < s.size(); u++) {
                              std::string v_s = "";
                              v_s += s[i];
                              v_s += s[j];
                              v_s += s[k];
                              v_s += s[l];
                              v_s += s[m];
                              v_s += s[n];
                              v_s += s[o];
                              v_s += s[p];
                              v_s += s[q];
                              v_s += s[r];
                              v_s += s[t];
                              v_s += s[u];
                              long long v = std::stoll(v_s);
                              if (v > max)
                                max = v;
                            }
                          }
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }

      sum += max;
    }

    return sum;
  }

  void dfs(int i, const std::string &s, std::string &curr, std::string &max) {
    if (curr.size() >= 12) {
      if (curr > max)
        max = curr;
      return;
    };

    if (i == s.size())
      return;
    if (curr.size() + (s.size() - i) < 12)
      return;

    curr.push_back(s[i]); // take s[i]
    dfs(i + 1, s, curr, max);
    curr.pop_back();

    // skip s[i]
    dfs(i + 1, s, curr, max);
  }

  long long task2B(std::vector<std::string> input) {
    long long sum = 0;
    for (const auto &s : input) {
      std::cout << s << "\n";
      std::string max = "";
      std::string str = "";
      dfs(0, s, str, max);
      sum += std::stoll(max);
    }
    return sum;
  }

  long long task2C(std::vector<std::string> input) {
    long long sum = 0;
    for (const auto s : input) {
      int lastpos = 0;
      std::string max = "";

      for (int limit = 12; limit > 0; limit--) {
        char max_c = s[lastpos];
        int max_i = lastpos;

        for (int i = lastpos + 1; i <= s.size() - limit; i++) {
          if (s[i] > max_c) {
            max_c = s[i];
            max_i = i;
          }
        }

        max += max_c;
        lastpos = max_i + 1;
      }

      std::cout << max << "\n";
      sum += std::stoll(max);
    }
    return sum;
  }
};

int main() {
  Solution sol;
  std::ifstream in("input.txt");
  // std::ifstream in("test.txt");
  if (!in)
    throw std::runtime_error("not found file");

  auto res = sol.parse(in);

  // for (auto i : res)
  //   std::cout << i.size() << "\n";

  // auto t1 = sol.task1(res);
  // std::cout << t1 << "\n";

  auto t2 = sol.task2C(res);
  // auto t2 = sol.task2(std::vector<std::string>());
  std::cout << t2 << "\n";

  return 0;
}
