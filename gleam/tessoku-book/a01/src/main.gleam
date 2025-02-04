import gleam/int
import gleam/io
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, _) = stdin.take_int(in)
  let ans = n * n
  ans |> int.to_string |> io.println
}
