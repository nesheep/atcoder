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

  s
  |> string.to_graphemes
  |> list.filter(fn(c) { c != "." })
  |> string.join("")
  |> io.println
}
