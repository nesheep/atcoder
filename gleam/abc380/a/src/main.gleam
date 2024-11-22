import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn count(s: List(a), v: a) -> Int {
  list.count(s, fn(c) { c == v })
}

pub fn main() {
  let in = stdin.stdin()
  let s = in |> read_string |> string.to_graphemes

  let cnt1 = count(s, "1")
  let cnt2 = count(s, "2")
  let cnt3 = count(s, "3")

  case cnt1, cnt2, cnt3 {
    1, 2, 3 -> "Yes"
    _, _, _ -> "No"
  }
  |> io.println
}
