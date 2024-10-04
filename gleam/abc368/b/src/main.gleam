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

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_strings |> list.map(parse_int)
}

fn loop(a: List(Int), acc: Int) -> Int {
  let cnt = a |> list.count(fn(v) { v > 0 })
  case cnt <= 1 {
    True -> acc
    False -> {
      let assert [a1, a2, ..rest] = a |> list.sort(int.compare) |> list.reverse
      loop([a1 - 1, a2 - 1, ..rest], acc + 1)
    }
  }
}

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_int
  let a = in |> read_ints
  let ans = loop(a, 0)
  ans |> int.to_string |> io.println
}
