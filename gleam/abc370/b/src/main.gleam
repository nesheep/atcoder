import gleam/dict.{type Dict}
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

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_strings |> list.map(parse_int)
}

fn read_elems(in: iterator.Iterator(String), n: Int) -> Dict(#(Int, Int), Int) {
  iterator.range(1, n)
  |> iterator.map(fn(i) {
    in
    |> read_ints
    |> iterator.from_list
    |> iterator.index
    |> iterator.map(fn(v) { #(#(i, v.1 + 1), v.0) })
  })
  |> iterator.flatten
  |> iterator.to_list
  |> dict.from_list
}

pub fn main() {
  let in = stdin.stdin()
  let n = in |> read_int
  let a = in |> read_elems(n)

  let ans =
    iterator.range(1, n)
    |> iterator.fold(1, fn(acc, v) {
      case acc > v {
        True -> a |> dict.get(#(acc, v)) |> result.unwrap(0)
        False -> a |> dict.get(#(v, acc)) |> result.unwrap(0)
      }
    })

  ans |> int.to_string |> io.println
}
