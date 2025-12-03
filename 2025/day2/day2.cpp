#include <cmath>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>

struct Entry {
  long long start;
  long long end;
};

namespace Day2 {
std::vector<Entry> parseFile(std::istream &in) {
  std::vector<Entry> res;
  std::string line;
  std::getline(in, line);
  std::stringstream ss(line);
  std::string tok;

  while (std::getline(ss, tok, ',')) {
    auto dash = tok.find('-');
    Entry newentry{.start = std::stoll(tok.substr(0, dash)),
                   .end = std::stoll(tok.substr(dash + 1))};
    res.push_back(newentry);
  }

  return res;
};

bool repeatedSubstringPattern(std::string s) {
    std::string concatenated = s + s;
    return concatenated.substr(1, concatenated.length() - 2).find(s) != std::string::npos;
}

int getLenOfNum(long long n) { return std::log10(n) + 1; }

long long getSum(const Entry &e) {
  long long sum = 0;

  for (long long i = e.start; i <= e.end; ++i) {
    std::string s = std::to_string(i);
    int len = s.size();
    if ((len & 1) != 0)
      continue;
    int halflen = len / 2;

    bool ok = true;
    for (int j = 0; j < halflen; ++j) {
      if (s[j] != s[halflen + j]) {
        ok = false;
        break;
      }
    }

    if (ok)
      sum += i;
  }

  return sum;
}

long long getSumWithRepeated(const Entry &e) {
  long long sum = 0;

  for (long long i = e.start; i <= e.end; ++i) {
    std::string s = std::to_string(i);
    if (repeatedSubstringPattern(s)) sum += i;
  }

  return sum;
}

long long getSumOfInvalids(std::vector<Entry> in) {
  long long sum = 0;

  for (const auto &e : in) {
    sum += getSumWithRepeated(e);
  }

  return sum;
}

} // namespace Day2

int main() {
  try {
    std::ifstream in("input.txt");
    // std::ifstream in("test.txt");

    if (!in) {
      std::cerr << "Error: file not found";
      return 1;
    }

    auto input = Day2::parseFile(in);

    // for (auto i : input)
    //   std::cout << "Start: " << i.start << " " << "end: " << i.end << "\n";

    // auto a = Day2::getSumOfInvalids(input);
    // std::cout << "Task 1 = " << a << "\n";

    auto b = Day2::getSumOfInvalids(input);
    std::cout << "Task 2 = " << b << std::endl;
    return 0;
  } catch (const std::exception &e) {
    std::cerr << "Error: " << e.what() << std::endl;
    return 1;
  }

  return 0;
}
