import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/string
import stdin

pub fn main() {
  let in = stdin.stdin()

  let assert [n, k] = in |> read_ints
  let a = in |> read_ints

  let #(f, s) = a |> list.split(n - k)
  let ans = s |> list.append(f)

  ans
  |> list.map(fn(v) { v |> int.to_string })
  |> string.join(" ")
  |> io.println
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in
  |> iterator.first
  |> result.unwrap("")
  |> string.trim
  |> string.split(" ")
  |> list.map(fn(v) { v |> int.parse |> result.unwrap(0) })
}
