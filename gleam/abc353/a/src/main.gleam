import gleam/int
import gleam/io
import gleam/list
import gleam/pair
import gleam/result
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let assert #([first, ..rest], _) = stdin.take_list(stdin.take_int, n)(in)

  let ans =
    rest
    |> list.index_map(fn(v, i) { #(v, i + 2) })
    |> list.find(fn(x) { pair.first(x) > first })
    |> result.map(pair.second)
    |> result.unwrap(-1)

  ans |> int.to_string |> io.println
}
