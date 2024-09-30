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

fn square_of_distance(a: #(Int, Int), b: #(Int, Int)) -> Int {
  let dx = pair.first(b) - pair.first(a)
  let dy = pair.second(b) - pair.second(a)
  dx * dx + dy * dy
}

pub fn main() {
  let in = stdin.stdin()
  let assert [xa, ya] = in |> read_ints
  let assert [xb, yb] = in |> read_ints
  let assert [xc, yc] = in |> read_ints

  let ab2 = square_of_distance(#(xa, ya), #(xb, yb))
  let bc2 = square_of_distance(#(xb, yb), #(xc, yc))
  let ac2 = square_of_distance(#(xa, ya), #(xc, yc))

  let assert [max, ..rest] =
    [ab2, bc2, ac2]
    |> list.sort(int.compare)
    |> list.reverse

  case max == int.sum(rest) {
    True -> "Yes"
    False -> "No"
  }
  |> io.println
}
