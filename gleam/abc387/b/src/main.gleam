import gleam/int
import gleam/io
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(x, _) = stdin.take_int(in)

  let ans =
    {
      use x <- yielder.flat_map(yielder.range(1, 9))
      use y <- yielder.map(yielder.range(1, 9))
      x * y
    }
    |> yielder.filter(fn(v) { v != x })
    |> yielder.fold(0, int.add)

  ans |> int.to_string |> io.println
}
