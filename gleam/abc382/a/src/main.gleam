import gleam/int
import gleam/io
import gleam/list
import gleam/string
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let #(d, in) = stdin.take_int(in)
  let #(s, _) = stdin.take_string(in)

  let c = s |> string.to_graphemes |> list.count(fn(c) { c == "@" })
  let ans = n - c + d

  ans |> int.to_string |> io.println
}
