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

pub fn main() {
  let in = stdin.stdin()
  let assert [s, t] = in |> read_strings

  let ans =
    iterator.range(1, string.length(s) - 1)
    |> iterator.map(fn(w) {
      iterator.range(0, w - 1)
      |> iterator.map(fn(c) {
        s
        |> string.to_graphemes
        |> iterator.from_list
        |> iterator.drop(c)
        |> iterator.sized_chunk(w)
        |> iterator.map(fn(u) { u |> list.first |> result.unwrap("") })
        |> iterator.to_list
        |> string.join("")
      })
    })
    |> iterator.flatten
    |> iterator.any(fn(u) { u == t })

  case ans {
    True -> "Yes"
    _ -> "No"
  }
  |> io.println
}
