import gleam/io
import gleam/list
import stdin

fn bool_to_string(b: Bool) -> String {
  case b {
    False -> "No"
    True -> "Yes"
  }
}

pub fn main() {
  let in = stdin.values()
  let #(a1, in) = stdin.take_int(in)
  let #(a2, in) = stdin.take_int(in)
  let #(a3, _) = stdin.take_int(in)

  let ans =
    [#(a1 * a2, a3), #(a2 * a3, a1), #(a3 * a1, a2)]
    |> list.any(fn(x) { x.0 == x.1 })

  ans |> bool_to_string |> io.println
}
