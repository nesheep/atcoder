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
    list.range(1, string.length(s) - 1)
    |> list.map(fn(w) {
      list.range(0, w - 1)
      |> list.map(fn(c) {
        s
        |> string.drop_left(c)
        |> string.to_graphemes
        |> list.sized_chunk(w)
        |> list.map(fn(u) { u |> list.first |> result.unwrap("") })
        |> string.join("")
      })
    })
    |> list.flatten
    |> list.any(fn(u) { u == t })

  case ans {
    True -> "Yes"
    _ -> "No"
  }
  |> io.println
}
