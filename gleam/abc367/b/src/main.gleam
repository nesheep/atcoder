import gleam/float
import gleam/int
import gleam/io
import gleam/iterator
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()
  let x = in |> read_float

  case float.modulo(x, 1.0) |> result.unwrap(0.0) {
    0.0 -> x |> float.round |> int.to_string
    _ -> x |> float.to_string
  }
  |> io.println
}

fn read_float(in: iterator.Iterator(String)) -> Float {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> float.parse
  |> result.unwrap(0.0)
}
