import gleam/int
import gleam/io
import stdin

fn parse_input(in: List(String)) -> Int {
  let #(a, _) = stdin.take_int(in)
  a
}

fn solve(a: Int) -> Int {
  case 400 % a {
    0 -> 400 / a
    _ -> -1
  }
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> int.to_string
  |> io.println
}
