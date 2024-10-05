import gleam/int
import gleam/io
import gleam/iterator.{Done, Next}
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

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}

pub fn main() {
  let in = stdin.stdin()
  let m = in |> read_int

  let pow3s =
    iterator.iterate(#(0, 1), fn(p) { #(p.0 + 1, p.1 * 3) })
    |> iterator.take(11)
    |> iterator.to_list
    |> list.reverse

  let ans =
    iterator.unfold(m, fn(v) {
      case v {
        0 -> Done
        v ->
          pow3s
          |> list.find(fn(p) { p.1 <= v })
          |> result.unwrap(#(0, 1))
          |> fn(p) { Next(p.0, v - p.1) }
      }
    })
    |> iterator.to_list

  ans |> list.length |> int.to_string |> io.println
  ans |> list.map(int.to_string) |> string.join(" ") |> io.println
}
