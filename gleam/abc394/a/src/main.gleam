import gleam/io
import gleam/list
import gleam/string
import stdin

pub fn main() {
  let in = stdin.values()
  let #(s, _) = stdin.take_string(in)

  let ans =
    s
    |> string.to_graphemes
    |> list.filter(fn(c) { c == "2" })
    |> string.join("")

  io.println(ans)
}
