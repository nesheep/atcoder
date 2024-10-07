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

fn count(list: List(Int)) -> Int {
  do_count(list, 0)
}

fn do_count(list: List(Int), acc: Int) -> Int {
  case list {
    [x, y, z, ..rest] -> {
      let cnt = case x == z {
        True -> acc + 1
        False -> acc
      }
      do_count([y, z, ..rest], cnt)
    }
    _ -> acc
  }
}

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_int
  let a = in |> read_ints

  a |> count |> int.to_string |> io.println
}
