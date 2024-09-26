import gleam/io
import gleam/iterator
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()

  let s = in |> iterator.first
  use s <- result.map(s)

  let ans = case s |> string.trim {
    "1 0" -> "Yes"
    "0 1" -> "No"
    _ -> "Invalid"
  }

  io.println(ans)
}
