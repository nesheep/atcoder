import gleam/float
import gleam/int
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn read_strings(in: iterator.Iterator(String)) -> List(String) {
  in |> read_string |> string.split(" ")
}

fn read_string_lines(in: iterator.Iterator(String), n: Int) -> List(String) {
  iterator.repeatedly(fn() { in |> read_string })
  |> iterator.take(n)
  |> iterator.to_list
}

fn parse_int(v: String) -> Int {
  v |> int.parse |> result.unwrap(0)
}

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_strings |> list.map(parse_int)
}

fn parse_float(v: String) -> Float {
  v |> float.parse |> result.unwrap(0.0)
}

fn read_float(in: iterator.Iterator(String)) -> Float {
  in |> read_string |> parse_float
}

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_string
  let _ = in |> read_strings
  let _ = in |> read_string_lines(0)
  let _ = in |> read_int
  let _ = in |> read_ints
  let _ = in |> read_float

  let _ = [] |> to_pairs
}

fn to_pairs(l: List(a)) -> List(#(a, a)) {
  list.zip(l, l |> list.rest |> result.unwrap([]))
}
