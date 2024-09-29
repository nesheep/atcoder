import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/pair
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn to_pairs(l: List(a)) -> List(#(a, a)) {
  list.zip(l, l |> list.rest |> result.unwrap([]))
}

pub fn main() {
  stdin.stdin()
  |> read_string
  |> string.to_graphemes
  |> list.index_map(fn(v, i) { #(v, i) })
  |> list.sort(fn(a, b) { string.compare(pair.first(a), pair.first(b)) })
  |> list.map(fn(a) { pair.second(a) })
  |> to_pairs
  |> list.map(fn(p) { int.absolute_value(pair.first(p) - pair.second(p)) })
  |> int.sum
  |> int.to_string
  |> io.println
}
