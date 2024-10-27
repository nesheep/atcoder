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
  let s = in |> read_string

  let ans = ["A", "B", "C"] |> list.all(string.contains(s, _))

  case ans {
    True -> "Yes"
    False -> "No"
  }
  |> io.println
}
