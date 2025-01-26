import gleam/io
import gleam/set
import gleam/string
import gleam/yielder
import stdin

fn bool_to_string(b: Bool) -> String {
  case b {
    True -> "Yes"
    False -> "No"
  }
}

pub fn main() {
  let s = stdin.read_all() |> string.trim

  let y =
    s
    |> string.to_graphemes
    |> yielder.from_list

  let ans =
    {
      yielder.zip(y, yielder.drop(y, 1))
      |> yielder.filter(fn(v) { v.0 == v.1 })
      |> yielder.map(fn(v) { v.0 })
      |> yielder.to_list
      |> set.from_list
      |> set.size
    }
    * 2
    == string.length(s)

  ans |> bool_to_string |> io.println
}
