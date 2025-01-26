import gleam/io
import gleam/list
import gleam/pair
import gleam/string
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(_, in) = stdin.take_int(in)
  let #(d, in) = stdin.take_int(in)
  let #(s, _) = stdin.take_string(in)

  let ans =
    s
    |> string.to_graphemes
    |> list.reverse
    |> yielder.from_list
    |> yielder.fold(#(0, ""), fn(acc, c) {
      let #(cnt, t) = acc
      case cnt < d, c == "@" {
        True, True -> #(cnt + 1, "." <> t)
        False, True -> #(cnt, "@" <> t)
        _, _ -> #(cnt, "." <> t)
      }
    })
    |> pair.second

  io.println(ans)
}
