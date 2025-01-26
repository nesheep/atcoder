import gleam/int
import gleam/io
import gleam/option.{None, Some}
import gleam/string
import gleam/yielder
import stdin

fn distance(a: #(Int, Int), b: #(Int, Int)) -> Int {
  int.absolute_value(a.0 - b.0) + int.absolute_value(a.1 - b.1)
}

pub fn main() {
  let in = stdin.values()
  let #(h, in) = stdin.take_int(in)
  let #(_, in) = stdin.take_int(in)
  let #(d, in) = stdin.take_int(in)
  let #(s, _) = stdin.take_list(stdin.take_string, h)(in)

  let assert Some(t) =
    s
    |> yielder.from_list
    |> yielder.index
    |> yielder.flat_map(fn(r_x) {
      let #(r, x) = r_x
      r
      |> string.to_graphemes
      |> yielder.from_list
      |> yielder.index
      |> yielder.map(fn(c_y) {
        let #(c, y) = c_y
        case c {
          "." -> Some(#(x, y))
          _ -> None
        }
      })
    })
    |> yielder.filter(option.is_some)
    |> yielder.to_list
    |> option.all

  let ans =
    t
    |> yielder.from_list
    |> yielder.flat_map(fn(a) {
      t
      |> yielder.from_list
      |> yielder.filter(fn(b) { a != b })
      |> yielder.map(fn(b) {
        t
        |> yielder.from_list
        |> yielder.filter(fn(c) { distance(c, a) <= d || distance(c, b) <= d })
        |> yielder.length
      })
    })
    |> yielder.fold(0, int.max)

  ans |> int.to_string |> io.println
}
