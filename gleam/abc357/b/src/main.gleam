import gleam/io
import gleam/iterator
import gleam/list
import gleam/order
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

pub fn main() {
  let in = stdin.stdin()
  let s = in |> read_string

  let cnt =
    s
    |> string.to_graphemes
    |> list.count(fn(x) { string.compare(x, "a") == order.Lt })

  case cnt > string.length(s) / 2 {
    True -> string.uppercase(s)
    False -> string.lowercase(s)
  }
  |> io.println
}
