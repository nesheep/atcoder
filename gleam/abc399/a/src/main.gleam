import gleam/int
import gleam/io
import gleam/list
import gleam/string
import stdin

fn parse_input(in: List(String)) -> #(String, String) {
  let #(_, in) = stdin.take_int(in)
  let #(s, in) = stdin.take_string(in)
  let #(t, _) = stdin.take_string(in)
  #(s, t)
}

fn solve(s_t: #(String, String)) -> Int {
  let #(s, t) = s_t
  string.to_graphemes(s)
  |> list.zip(string.to_graphemes(t))
  |> list.fold(0, fn(acc, x_y) {
    let #(x, y) = x_y
    case x == y {
      False -> acc + 1
      True -> acc
    }
  })
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> int.to_string
  |> io.println
}
