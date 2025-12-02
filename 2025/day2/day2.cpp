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
    Entry newentry;
    int dash = tok.find('-');
    newentry.start = std::stoll(tok.substr(0, dash));
    newentry.end = std::stoll(tok.substr(dash + 1));
    res.push_back(newentry);
  }

  return res;
};
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

    for (auto i : input)
      std::cout << "Start: " << i.start << " " << "end: " << i.end << "\n";

    // auto a = Day1::getPasscode(input);
    // auto b = Day1::getPasscodeWithZero(input);
    // std::cout << "Task 1 = " << a << "\n";
    // std::cout << "Task 2 = " << b << std::endl;
    return 0;
  } catch (const std::exception &e) {
    std::cerr << "Error: " << e.what() << std::endl;
    return 1;
  }

  return 0;
}
