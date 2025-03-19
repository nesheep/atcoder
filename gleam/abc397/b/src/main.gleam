import gleam/int
import gleam/io
import gleam/list
import gleam/string
import stdin

fn parse_input(in: List(String)) -> String {
  let #(s, _) = stdin.take_string(in)
  s
}

fn solve(s: String) -> Int {
  s
  |> string.to_graphemes
  |> list.fold(#(0, 0), fn(acc, c) {
    let #(i, a) = acc
    case i % 2, c {
      0, "i" | 1, "o" -> #(i + 1, a)
      _, _ -> #(i + 2, a + 1)
    }
  })
  |> fn(acc) {
    let #(i, a) = acc
    case i % 2 == 0 {
      True -> a
      False -> a + 1
    }
  }
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> int.to_string
  |> io.println
}
