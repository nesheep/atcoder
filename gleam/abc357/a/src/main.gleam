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

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_strings |> list.map(parse_int)
}

pub fn main() {
  let in = stdin.stdin()
  let assert [n, m] = in |> read_ints
  let h = in |> read_ints

  let ans =
    h
    |> iterator.from_list
    |> iterator.scan(0, int.add)
    |> iterator.index
    |> iterator.find(fn(x) { x.0 > m })
    |> result.map(pair.second)
    |> result.map_error(fn(_) { n })
    |> result.unwrap_both

  ans |> int.to_string |> io.println
}
