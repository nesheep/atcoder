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

fn parse_int(v: String) -> Int {
  v |> int.parse |> result.unwrap(0)
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_string |> string.split(" ") |> list.map(parse_int)
}

pub fn main() {
  let in = stdin.stdin()
  let assert [r, g, b] = in |> read_ints
  let c = in |> read_string

  case c {
    "Red" -> int.min(g, b)
    "Green" -> int.min(r, b)
    _ -> int.min(r, g)
  }
  |> int.to_string
  |> io.println
}
