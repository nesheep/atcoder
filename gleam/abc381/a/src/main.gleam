import gleam/int
import gleam/io
import gleam/string
import gleam/yielder
import stdin

fn parse_input(in: String) -> #(Int, String) {
  let assert [n, s] = in |> string.trim |> string.split("\n")
  let assert Ok(n) = int.parse(n)
  #(n, s)
}

fn bool_to_string(b: Bool) -> String {
  case b {
    True -> "Yes"
    False -> "No"
  }
}

pub fn main() {
  let #(n, s) = stdin.read_all() |> parse_input

  let ans = case n % 2 {
    0 -> False
    _ ->
      yielder.range(0, n - 1)
      |> yielder.map(fn(i) {
        case { n + 1 } / 2 - 1 {
          m if i < m -> "1"
          m if i > m -> "2"
          _ -> "/"
        }
      })
      |> yielder.zip(s |> string.to_graphemes |> yielder.from_list)
      |> yielder.all(fn(v) { v.0 == v.1 })
  }

  ans |> bool_to_string |> io.println
}
