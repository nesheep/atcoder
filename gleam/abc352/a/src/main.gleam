import gleam/io
import stdin

fn bool_to_string(b: Bool) -> String {
  case b {
    False -> "No"
    True -> "Yes"
  }
}

pub fn main() {
  let in = stdin.values()
  let #(_, in) = stdin.take_int(in)
  let #(x, in) = stdin.take_int(in)
  let #(y, in) = stdin.take_int(in)
  let #(z, _) = stdin.take_int(in)

  let ans = x < z == z < y

  ans |> bool_to_string |> io.println
}
