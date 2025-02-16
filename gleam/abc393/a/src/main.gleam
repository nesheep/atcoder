import gleam/int
import gleam/io
import stdin

pub fn main() {
  let in = stdin.values()
  let #(s1, in) = stdin.take_string(in)
  let #(s2, _) = stdin.take_string(in)

  let ans = case s1, s2 {
    "sick", "sick" -> 1
    "sick", "fine" -> 2
    "fine", "sick" -> 3
    _, _ -> 4
  }

  ans |> int.to_string |> io.println
}
