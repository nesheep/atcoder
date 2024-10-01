import gleam/dict
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
  let assert [h, _] = in |> read_ints
  let assert [si, sj] = in |> read_ints
  let c =
    iterator.repeatedly(fn() { in |> read_string })
    |> iterator.take(h)
    |> iterator.to_list
    |> list.index_map(fn(l, i) {
      l
      |> string.to_graphemes
      |> list.index_map(fn(s, j) { #(#(i + 1, j + 1), s == ".") })
    })
    |> list.flatten
    |> dict.from_list
  let x = in |> read_string

  let #(ei, ej) =
    x
    |> string.to_graphemes
    |> list.fold(#(si, sj), fn(t, d) {
      let #(ti, tj) = t
      let #(di, dj) = case d {
        "L" -> #(0, -1)
        "R" -> #(0, 1)
        "U" -> #(-1, 0)
        "D" -> #(1, 0)
        _ -> #(0, 0)
      }
      let next = #(ti + di, tj + dj)
      case c |> dict.get(next) |> result.unwrap(False) {
        True -> next
        False -> t
      }
    })

  let ans = int.to_string(ei) <> " " <> int.to_string(ej)
  ans |> io.println
}
