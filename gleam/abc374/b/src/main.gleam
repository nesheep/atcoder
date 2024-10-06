import gleam/int
import gleam/io
import gleam/iterator
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

pub fn main() {
  let in = stdin.stdin()
  let s = in |> read_string
  let t = in |> read_string

  let u = s |> string.pad_right(string.length(t), ".") |> string.to_graphemes
  let v = t |> string.pad_right(string.length(s), ".") |> string.to_graphemes

  let ans =
    iterator.zip(iterator.from_list(u), iterator.from_list(v))
    |> iterator.index
    |> iterator.find(fn(x) { x.0.0 != x.0.1 })
    |> result.map_error(fn(_) { 0 })
    |> result.map(fn(x) { x.1 + 1 })
    |> result.unwrap_both

  ans |> int.to_string |> io.println
}
