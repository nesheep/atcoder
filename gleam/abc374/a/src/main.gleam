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

  case s |> string.ends_with("san") {
    True -> "Yes"
    False -> "No"
  }
  |> io.println
}
