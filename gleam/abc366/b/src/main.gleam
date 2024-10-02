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

pub fn main() {
  let in = stdin.stdin()
  let n = in |> read_int
  let s = in |> read_string_lines(n)

  let m = s |> list.map(string.length) |> list.fold(0, int.max)

  let ans =
    s
    |> iterator.from_list
    |> iterator.map(string.pad_right(_, m, "*"))
    |> iterator.map(string.to_graphemes)
    |> iterator.to_list
    |> list.transpose
    |> iterator.from_list
    |> iterator.map(fn(s) { s |> list.drop_while(fn(x) { x == "*" }) })
    |> iterator.map(list.reverse)
    |> iterator.map(string.join(_, ""))
    |> iterator.to_list

  ans |> list.map(io.println)
}
