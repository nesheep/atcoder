import gleam/io
import gleam/list
import gleam/string
import stdin

pub fn main() {
  let in = stdin.values()
  let #(_, in) = stdin.take_int(in)
  let #(c1, in) = stdin.take_string(in)
  let #(c2, in) = stdin.take_string(in)
  let #(s, _) = stdin.take_string(in)

  let ans =
    s
    |> string.to_graphemes
    |> list.map(fn(c) {
      case c == c1 {
        False -> c2
        True -> c
      }
    })
    |> string.join("")

  io.println(ans)
}
