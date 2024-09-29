import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/option.{None, Some}
import gleam/pair
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

pub fn main() {
  let in = stdin.stdin()
  let s = in |> read_string

  s
  |> string.to_graphemes
  |> list.index_map(fn(v, i) { #(v, i) })
  |> list.sort(fn(a, b) { string.compare(pair.first(a), pair.first(b)) })
  |> list.map(fn(a) { pair.second(a) })
  |> list.fold(#(None, 0), fn(acc, a) {
    case acc {
      #(None, _) -> #(Some(a), 0)
      #(Some(p), sum) -> #(Some(a), sum + int.absolute_value(a - p))
    }
  })
  |> pair.second
  |> int.to_string
  |> io.println
}
