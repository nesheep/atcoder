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

  let ans = s |> list.count(fn(x) { x == "Takahashi" })

  ans |> int.to_string |> io.println
}
