import gleam/io
import gleam/string
import stdin

fn parse_input(in: List(String)) -> Int {
  let #(n, _) = stdin.take_int(in)
  n
}

fn solve(n: Int) -> String {
  let a = string.repeat("-", { n - 1 } / 2)
  let b = case n % 2 {
    0 -> "=="
    _ -> "="
  }
  a <> b <> a
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> io.println
}
