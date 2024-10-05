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

fn read_int_string_lines(
  in: iterator.Iterator(String),
  n: Int,
) -> List(#(Int, String)) {
  iterator.repeatedly(fn() { in |> read_string })
  |> iterator.take(n)
  |> iterator.map(fn(v) {
    let assert [a, s] = v |> string.split(" ")
    #(a |> int.parse |> result.unwrap(0), s)
  })
  |> iterator.to_list
}

pub fn main() {
  let in = stdin.stdin()
  let assert [_, m] = in |> read_ints
  let cmds = in |> read_int_string_lines(m)

  cmds
  |> list.fold(dict.new(), fn(flgs, c) {
    let #(a, b) = c
    let ans = case b {
      "M" -> {
        case flgs |> dict.get(a) {
          Ok(_) -> False
          _ -> True
        }
      }
      _ -> False
    }
    case ans {
      True -> "Yes"
      False -> "No"
    }
    |> io.println
    case ans {
      True -> flgs |> dict.insert(a, False)
      False -> flgs
    }
  })
}
