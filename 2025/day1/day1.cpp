#include <fstream>
#include <iostream>
#include <stdexcept>
#include <vector>

enum class Dir { L, R };

struct Entry {
  Dir dir;
  int value;
};

namespace Day1 {
std::vector<Entry> parseFile(std::istream &in) {
  std::vector<Entry> res;
  std::string line;

  while (std::getline(in, line)) {
    if (line.empty())
      continue;

    Entry newentry;
    switch (line[0]) {
    case 'R':
      newentry.dir = Dir::R;
      break;
    case 'L':
      newentry.dir = Dir::L;
      break;
    default:
      throw std::runtime_error("bad dir");
    }
    newentry.value = std::stoi(line.substr(1));
    res.push_back(newentry);
  }

  return res;
}

// task 1
int getPasscode(std::vector<Entry> &input) {
  int zeroCnt = 0;
  int curr = 50;

  for (const auto i : input) {
    if (i.dir == Dir::L) {
      curr -= i.value;
    } else {
      curr += i.value;
    }
    curr = (curr % 100 + 100) % 100;
    if (curr == 0)
      zeroCnt++;
  }

  return zeroCnt;
}
};

int main() {
  try {
    std::ifstream in("input.txt");
    // std::ifstream in("test.txt");
    auto input = Day1::parseFile(in);

    // for (auto i : input)
    //   std::cout << "Dir: " << i.dir << " " << "value: " << i.value << "\n";

    auto zeroCnt = Day1::getPasscode(input);
    std::cout << "Task 1 = " << zeroCnt << std::endl;
    return 0;
  } catch (const std::exception& e) {
      std::cerr << "Error: " << e.what() << std::endl;
      return 1;
  }
  
  return 0;
};
