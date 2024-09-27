import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()
  let assert [n, t, a] = in |> read_ints

  case int.max(t, a) >= n / 2 + 1 {
    True -> "Yes"
    _ -> "No"
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
