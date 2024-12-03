# Advent of Code - 2024
> https://adventofcode.com/2024/

## Day 1
Improvements:
- Instead of `os.ReadFile`, could have use `go:embed` to embed those file directly into the static assets
- Usage `fmt.Sscanf()` to parse the lines into the formatted strings to build up the lists 

## Day 3
Part Two Solution:
- Replace all new lines with the "" - Trim Newliens
- Replace all do() with \n do
- Replace all don't() with \n dont't()
- Remove all don't() lines
- Run the algo