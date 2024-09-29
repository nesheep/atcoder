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
  let assert [a, b, c, d, e, f] = in |> read_ints
  let assert [g, h, i, j, k, l] = in |> read_ints

  let c1 = a < j && g < d
  let c2 = b < k && h < e
  let c3 = c < l && i < f
  let ans = c1 && c2 && c3

  case ans {
    True -> "Yes"
    _ -> "No"
  }
  |> io.println
}
