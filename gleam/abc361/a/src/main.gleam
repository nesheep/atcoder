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

fn insert(a: List(a), k: Int, x: a) -> List(a) {
  let #(f, s) = a |> list.split(k)
  list.append(f, [x, ..s])
}

pub fn main() {
  let in = stdin.stdin()
  let assert [_, k, x] = in |> read_ints
  let a = in |> read_ints

  a
  |> insert(k, x)
  |> list.map(fn(v) { v |> int.to_string })
  |> string.join(" ")
  |> io.println
}
