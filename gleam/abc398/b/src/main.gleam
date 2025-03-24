import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/pair
import gleam/result
import stdin

fn parse_input(in: List(String)) -> List(Int) {
  in |> list.map(fn(x) { x |> int.parse |> result.unwrap(0) })
}

fn increment(x: Option(Int)) -> Int {
  case x {
    Some(x) -> x + 1
    None -> 1
  }
}

fn count(a: List(Int)) -> dict.Dict(Int, Int) {
  list.fold(a, dict.new(), fn(d, v) { dict.upsert(d, v, increment) })
}

fn judge(l: List(Int)) -> Bool {
  case l {
    [x, y, ..] -> x >= 3 && y >= 2
    _ -> False
  }
}

fn solve(a: List(Int)) -> Bool {
  a
  |> count
  |> dict.to_list
  |> list.map(pair.second)
  |> list.sort(fn(x, y) { int.compare(y, x) })
  |> judge
}

fn output(b: Bool) -> String {
  case b {
    False -> "No"
    True -> "Yes"
  }
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> output
  |> io.println
}
