import gleam/io
import gleam/set
import stdin

fn bool_to_string(b: Bool) -> String {
  case b {
    False -> "No"
    True -> "Yes"
  }
}

pub fn main() {
  let in = stdin.values()
  let #(l, _) = stdin.take_list(stdin.take_int, 4)(in)
  let ans = l |> set.from_list |> set.size == 2
  ans |> bool_to_string |> io.println
}
