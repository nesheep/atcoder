import gleam/int
import gleam/io
import gleam/list
import stdin

fn bool_to_string(b: Bool) -> String {
  case b {
    True -> "Yes"
    False -> "No"
  }
}

pub fn main() {
  let in = stdin.values()
  let #(a, in) = stdin.take_int(in)
  let #(b, in) = stdin.take_int(in)
  let #(c, _) = stdin.take_int(in)

  let assert [f, ..r] = [a, b, c] |> list.sort(int.compare) |> list.reverse
  let ans = f == int.sum(r) || f * 2 == int.sum(r)

  ans |> bool_to_string |> io.println
}
