import gleam/float
import gleam/int
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_int
  let _ = in |> read_ints
  let _ = in |> read_float
}

fn read_int(in: iterator.Iterator(String)) -> Int {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> int.parse
  |> result.unwrap(0)
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> string.split(" ")
  |> list.map(fn(v) { v |> int.parse |> result.unwrap(0) })
}

fn read_float(in: iterator.Iterator(String)) -> Float {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> float.parse
  |> result.unwrap(0.0)
}
