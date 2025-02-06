import gleam/io
import gleam/string
import stdin

pub fn main() {
  let in = stdin.values()
  let #(s, _) = stdin.take_string(in)

  let assert Ok(a) = string.first(s)
  let ans = a <> "UPC"

  io.println(ans)
}
