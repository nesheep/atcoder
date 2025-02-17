import gleam/int
import gleam/io
import gleam/pair
import gleam/result
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(h, _) = stdin.take_int(in)

  let ans =
    yielder.iterate(#(0, 1), fn(x) {
      let #(n, acc) = x
      #(n + acc, acc * 2)
    })
    |> yielder.map(pair.first)
    |> yielder.index
    |> yielder.find(fn(x) { pair.first(x) > h })
    |> result.map(pair.second)
    |> result.unwrap(0)

  ans |> int.to_string |> io.println
}
