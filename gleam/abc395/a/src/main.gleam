import gleam/io
import gleam/list
import gleam/pair
import stdin

fn parse_input(in: List(String)) -> List(Int) {
  let #(n, in) = stdin.take_int(in)
  let #(a, _) = stdin.take_list(stdin.take_int, n)(in)
  a
}

fn solve(a: List(Int)) -> Bool {
  list.window_by_2(a)
  |> list.all(fn(x) { pair.first(x) < pair.second(x) })
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
