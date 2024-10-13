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

fn read_x(in: iterator.Iterator(String), n: Int) -> List(List(Int)) {
  iterator.repeatedly(fn() { in |> read_ints })
  |> iterator.take(n)
  |> iterator.to_list
}

pub fn main() {
  let in = stdin.stdin()
  let assert [n, m] = in |> read_ints
  let a = in |> read_ints
  let x = in |> read_x(n)

  let sum =
    x
    |> list.fold(list.repeat(0, m), fn(acc, y) {
      list.zip(acc, y) |> list.map(fn(z) { z.0 + z.1 })
    })

  let ans =
    list.zip(a, sum)
    |> list.all(fn(z) { z.0 <= z.1 })

  case ans {
    True -> "Yes"
    False -> "No"
  }
  |> io.println
}
