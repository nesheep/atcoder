import gleam/int
import gleam/io
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.values()
  let #(s, _) = stdin.take_string(in)

  let #(a, b) = case string.to_graphemes(s) {
    [a, "x", b] -> #(
      a |> int.parse |> result.unwrap(0),
      b |> int.parse |> result.unwrap(0),
    )
    _ -> #(0, 0)
  }
  let ans = a * b

  ans |> int.to_string |> io.println
}
