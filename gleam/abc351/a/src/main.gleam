import gleam/int
import gleam/io
import stdin

fn parse_input(in: List(String)) -> #(List(Int), List(Int)) {
  let #(a, in) = stdin.take_list(stdin.take_int, 9)(in)
  let #(b, _) = stdin.take_list(stdin.take_int, 8)(in)
  #(a, b)
}

fn solve(ab: #(List(Int), List(Int))) -> Int {
  let #(a, b) = ab
  int.sum(a) - int.sum(b) + 1
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> int.to_string
  |> io.println
}
