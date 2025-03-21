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
  let assert [n, l, r] = in |> read_ints

  let a = case l == 1 {
    True -> []
    False -> list.range(1, l - 1)
  }

  let b = list.range(r, l)

  let c = case r == n {
    True -> []
    False -> list.range(r + 1, n)
  }

  let ans = a |> list.append(b) |> list.append(c)
  ans |> list.map(int.to_string) |> string.join(" ") |> io.println
}
