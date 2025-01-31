import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/pair
import gleam/set
import gleam/string
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(h, in) = stdin.take_int(in)
  let #(_, in) = stdin.take_int(in)
  let #(x, in) = stdin.take_int(in)
  let #(y, in) = stdin.take_int(in)
  let #(s, in) = stdin.take_list(stdin.take_string, h)(in)
  let #(t, _) = stdin.take_string(in)

  let d =
    s
    |> yielder.from_list
    |> yielder.map(string.to_graphemes)
    |> yielder.map(yielder.from_list)
    |> yielder.map(yielder.index)
    |> yielder.index
    |> yielder.flat_map(fn(l_x) {
      let #(l, x) = l_x
      yielder.map(l, fn(a_y) {
        let #(a, y) = a_y
        #(#(x + 1, y + 1), a)
      })
    })
    |> yielder.to_list
    |> dict.from_list

  let ans =
    t
    |> string.to_graphemes
    |> yielder.from_list
    |> yielder.fold(#(#(x, y), set.new()), fn(acc, a) {
      let #(#(x, y), c) = acc
      let next = case a {
        "U" -> #(x - 1, y)
        "D" -> #(x + 1, y)
        "L" -> #(x, y - 1)
        "R" -> #(x, y + 1)
        _ -> #(x, y)
      }
      case dict.get(d, next) {
        Ok(".") -> #(next, c)
        Ok("@") -> #(next, set.insert(c, next))
        _ -> acc
      }
    })

  [
    ans |> pair.first |> pair.first,
    ans |> pair.first |> pair.second,
    ans |> pair.second |> set.size,
  ]
  |> list.map(int.to_string)
  |> string.join(" ")
  |> io.println
}
