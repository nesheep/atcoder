import gleam/io
import gleam/iterator
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()

  let s = in |> iterator.first
  use s <- result.map(s)

  let ans = case s |> string.trim |> string.split(" ") {
    [">", "<", _] | ["<", ">", _] -> "A"
    ["<", _, "<"] | [">", _, ">"] -> "B"
    _ -> "C"
  }

  io.println(ans)
}
