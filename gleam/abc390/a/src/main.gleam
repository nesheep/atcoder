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
  let #(a, _) = stdin.take_list(stdin.take_int, 5)(in)

  let ans = case a {
    [2, 1, 3, 4, 5] -> True
    [1, 3, 2, 4, 5] -> True
    [1, 2, 4, 3, 5] -> True
    [1, 2, 3, 5, 4] -> True
    _ -> False
  }

  ans |> bool_to_string |> io.println
}
