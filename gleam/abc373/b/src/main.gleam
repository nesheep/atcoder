import gleam/dict
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

fn get_index(d: dict.Dict(String, Int), c: Int) -> Int {
  let assert Ok(a) = string.utf_codepoint(c)
  d |> dict.get(string.from_utf_codepoints([a])) |> result.unwrap(0)
}

pub fn main() {
  let in = stdin.stdin()
  let s = in |> read_string

  let d =
    s
    |> string.split("")
    |> list.index_map(fn(v, i) { #(v, i) })
    |> dict.from_list

  let start = 65
  iterator.range(start + 1, 90)
  |> iterator.fold(#(get_index(d, start), 0), fn(acc, c) {
    let #(pc, pa) = acc
    let i = get_index(d, c)
    #(i, pa + int.absolute_value(i - pc))
  })
  |> pair.second
  |> int.to_string
  |> io.println
}
