import gleam/int
import gleam/io
import gleam/list
import gleam/string
import stdin

fn loop(s: List(String), acc: Int) -> Int {
  case s {
    [] -> acc
    ["0", "0", ..rest] -> loop(rest, acc + 1)
    [_, ..rest] -> loop(rest, acc + 1)
  }
}

pub fn main() {
  let in = stdin.values()
  let #(s, _) = stdin.take_string(in)
  let ans = s |> string.to_graphemes |> list.reverse |> loop(0)
  ans |> int.to_string |> io.println
}
