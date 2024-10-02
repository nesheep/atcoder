import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/pair
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn read_strings(in: iterator.Iterator(String)) -> List(String) {
  in |> read_string |> string.split(" ")
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

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_int
  let a = in |> read_ints

  a
  |> list.index_map(fn(x, i) { #(x, i + 1) })
  |> list.sort(fn(x, y) { int.compare(x.0, y.0) })
  |> list.map(pair.second)
  |> list.reverse
  |> list.drop(1)
  |> list.first
  |> result.unwrap(0)
  |> int.to_string
  |> io.println
}
