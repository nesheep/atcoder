import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

pub fn main() {
  let in = stdin.stdin()
  iterator.repeatedly(fn() { in |> read_string })
  |> iterator.take(12)
  |> iterator.to_list
  |> list.index_map(fn(s, i) { string.length(s) == i + 1 })
  |> list.count(fn(v) { v })
  |> int.to_string
  |> io.println
}
