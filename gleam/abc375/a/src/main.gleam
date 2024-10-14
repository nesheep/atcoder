import gleam/bit_array
import gleam/int
import gleam/io
import gleam/iterator
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
  let n = in |> read_int
  let s = in |> read_string

  let b = bit_array.from_string(s)

  let ans =
    iterator.range(0, n - 3)
    |> iterator.fold(0, fn(acc, i) {
      case b |> bit_array.slice(i, 3) {
        Ok(<<"#.#">>) -> acc + 1
        _ -> acc
      }
    })

  ans |> int.to_string |> io.println
}
