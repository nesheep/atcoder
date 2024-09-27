import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()
  let assert [a, b, c] = in |> read_ints

  case b < c {
    True ->
      case a < b || a > c {
        True -> "Yes"
        _ -> "No"
      }
    _ ->
      case a < b && a > c {
        True -> "Yes"
        _ -> "No"
      }
  }
  |> io.println
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> string.split(" ")
  |> list.map(fn(v) { v |> int.parse |> result.unwrap(0) })
}
