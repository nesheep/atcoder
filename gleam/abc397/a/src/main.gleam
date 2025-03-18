import gleam/int
import gleam/io
import stdin

fn parse_input(in: List(String)) -> Float {
  let #(x, _) = stdin.take_float(in)
  x
}

fn solve(x: Float) -> Int {
  case x {
    x if x >=. 38.0 -> 1
    x if x >=. 37.5 -> 2
    _ -> 3
  }
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> int.to_string
  |> io.println
}
