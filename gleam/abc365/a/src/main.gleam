import gleam/int
import gleam/io
import gleam/iterator
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()
  let y = in |> read_int

  case y % 4, y % 100, y % 400 {
    0, 0, 0 -> 366
    0, 0, _ -> 365
    0, _, _ -> 366
    _, _, _ -> 365
  }
  |> int.to_string
  |> io.println
}

fn read_int(in: iterator.Iterator(String)) -> Int {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> int.parse
  |> result.unwrap(0)
}
