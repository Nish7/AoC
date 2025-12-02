#include <cstdlib>
#include <fstream>
#include <iostream>
#include <sstream>
#include <future>
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
  long long sum = 0;
  for (int i = e.start; i <= e.end; i++) {
    std::string i_s = std::to_string(i);
    int len = i_s.size();
    if (len % 2 != 0)
      continue;
    len = len / 2;
    std::string a = i_s.substr(0, len);
    std::string b = i_s.substr(len);
    if (a == b) {
      // std::cout << a << " " << b << "\n";
      // std::cout << "same\n";
      sum += i;
    }
  }

  return sum;
}

long long getSumOfInvalids(std::vector<Entry> in) { 
  long long sum = 0;
  std::vector<std::future<long long>> futs;
  futs.reserve(in.size());

  for (const auto &e : in)
    futs.push_back(std::async(std::launch::async, getSum, e));

  for (auto &f : futs)
    sum += f.get();

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
