import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/pair
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()
  let n = in |> read_int

  let ans =
    iterator.repeatedly(fn() { in |> read_string })
    |> iterator.take(n - 1)
    |> iterator.to_list
    |> to_pairs
    |> list.any(fn(p) { pair.first(p) == "sweet" && pair.second(p) == "sweet" })

  case ans {
    True -> "No"
    _ -> "Yes"
  }
  |> io.println
}

fn to_pairs(l: List(a)) -> List(#(a, a)) {
  list.zip(l, l |> list.rest |> result.unwrap([]))
}

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn parse_int(v: String) -> Int {
  v |> int.parse |> result.unwrap(0)
}

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}
