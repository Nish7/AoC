#include <fstream>
#include <iostream>
#include <ranges>
#include <vector>

using namespace std;

struct Entry {
  long long x;
  long long y;
};


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

long long getLargestArea(vector<Entry>& input) {
    long long max = 0;
    for (int i = 0; i < input.size() - 1; i++ ){
        Entry vali = input[i];
        for (int j = i + 1; j < input.size(); j++){
           Entry valj = input[j];
           long long area = (vali.x - valj.x + 1) * (vali.y - valj.y + 1);
           if (area > max) max = area;
        }
    }
    
    return max;
}


} // namespace Day8

int main() {
  // ifstream input("test.txt");
  ifstream input("input.txt");
  // auto vec = Day8::parseInput(test);
  auto vec = Day9::parseInput(input);
  cout << Day9::getLargestArea(vec) << endl;

  // for (auto c : vec){
  //         cout << c.x << " " << c.y << "\n";
  // }
  

  return 0;
}
