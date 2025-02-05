import gleam/int
import gleam/io
import stdin

pub fn main() {
  let in = stdin.values()
  let #(a, in) = stdin.take_int(in)
  let #(b, _) = stdin.take_int(in)
  let ans = { a + b } * { a + b }
  ans |> int.to_string |> io.println
}
