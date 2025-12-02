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

int getLenOfNum(long long n) { return std::log10(n) + 1; }

// 12 // 10 = 1
// 12 mod 10 2
//
// 1212 // 100 = 12
// 1212 mod 100 = 12
//
// 5555 // 100 = 12
// 5555 mod 100 = 12
//
// 5555 mod X = 12
// log2(5555,10)
long long getSum(const Entry &e) {
  std::cout << "range: " << e.start << " " << e.end  << "\n";
  long long sum = 0;
  int startlen = getLenOfNum(e.start);
  int endlen = getLenOfNum(e.end);

  if (startlen != endlen) {
      // return;
  }

  int pow = std::pow(10, startlen / 2);
  std::cout << startlen << " " << pow << "\n";
  for (int i = e.start; i <= e.end; i++) {
    long long first = i / pow;
    long long last = i % pow;
    if (first == last) {
        std::cout << "found: " << first << " " << last << "\n";
        sum += i;
    }
  }

  return sum;
}

long long getSumOfInvalids(std::vector<Entry> in) {
  long long sum = 0;

  for (const auto &e : in) {
    sum += getSum(e);
  }

  return sum;
}
} // namespace Day2

int main() {
  try {
    // std::ifstream in("input.txt");
    std::ifstream in("test.txt");

    if (!in) {
      std::cerr << "Error: file not found";
      return 1;
    }

    auto input = Day2::parseFile(in);
    // std::cout << "size" << input.size();
    // std::cout.flush();

    // for (auto i : input)
    //   std::cout << "Start: " << i.start << " " << "end: " << i.end << "\n";

    auto a = Day2::getSumOfInvalids(input);
    std::cout << "Task 1 = " << a << "\n";

    // auto b = Day1::getPasscodeWithZero(input);
    // std::cout << "Task 2 = " << b << std::endl;
    return 0;
  } catch (const std::exception &e) {
    std::cerr << "Error: " << e.what() << std::endl;
    return 1;
  }

  return 0;
}
