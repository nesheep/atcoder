import gleam/io
import gleam/string
import stdin

pub fn main() {
  let n = stdin.read_all() |> string.trim
  let assert [a, b, c] = n |> string.to_graphemes
  let bca = [b, c, a] |> string.join("")
  let cab = [c, a, b] |> string.join("")
  [bca, cab] |> string.join(" ") |> io.println
}
