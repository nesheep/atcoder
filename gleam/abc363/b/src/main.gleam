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
  let assert [n, t, p] = in |> read_ints
  let l = in |> read_ints

  let ans =
    l
    |> list.sort(int.compare)
    |> list.drop(n - p)
    |> list.first
    |> result.unwrap(0)
    |> fn(v) { t - v }
    |> int.max(0)

  ans |> int.to_string |> io.println
}
