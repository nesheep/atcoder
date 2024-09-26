import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()

  let s = in |> iterator.first
  use s <- result.map(s)

  let assert [a, b] =
    s
    |> string.trim
    |> string.split(" ")
    |> list.map(fn(v) { v |> int.parse |> result.unwrap(0) })

  let ans = case int.max(a, b) - int.min(a, b) {
    0 -> 1
    d if d % 2 == 1 -> 2
    _ -> 3
  }

  io.println(ans |> int.to_string)
}
