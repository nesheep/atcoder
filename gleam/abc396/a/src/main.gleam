import gleam/io
import gleam/list
import gleam/set
import stdin

fn parse_input(in: List(String)) -> List(Int) {
  let #(n, in) = stdin.take_int(in)
  let #(a, _) = stdin.take_list(stdin.take_int, n)(in)
  a
}

fn solve(a: List(Int)) -> Bool {
  a
  |> list.window(3)
  |> list.any(fn(v) { v |> set.from_list |> set.size == 1 })
}

fn bool_to_string(b: Bool) -> String {
  case b {
    False -> "No"
    True -> "Yes"
  }
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> bool_to_string
  |> io.println
}
