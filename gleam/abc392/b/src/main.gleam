import gleam/int
import gleam/io
import gleam/list
import gleam/string
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let #(m, in) = stdin.take_int(in)
  let #(a, _) = stdin.take_list(stdin.take_int, m)(in)

  let ans =
    list.range(1, n)
    |> list.filter(fn(x) { list.all(a, fn(y) { x != y }) })

  ans |> list.length |> int.to_string |> io.println
  ans |> list.map(int.to_string) |> string.join(" ") |> io.println
}
