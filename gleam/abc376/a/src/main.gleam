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
  let assert [_, c] = in |> read_ints
  let assert [first, ..rest] = in |> read_ints

  let ans =
    rest
    |> list.fold(#(1, first), fn(acc, x) {
      let #(total, prev) = acc
      case x - prev >= c {
        True -> #(total + 1, x)
        False -> #(total, prev)
      }
    })
    |> pair.first

  ans |> int.to_string |> io.println
}
