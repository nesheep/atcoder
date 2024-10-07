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
  case in |> read_string {
    "AtCoder Land" -> "Yes"
    _ -> "No"
  }
  |> io.println
}
