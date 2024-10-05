import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/pair
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn parse_int(v: String) -> Int {
  v |> int.parse |> result.unwrap(0)
}

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}

fn read_int_string_lines(
  in: iterator.Iterator(String),
  n: Int,
) -> List(#(Int, String)) {
  iterator.repeatedly(fn() { in |> read_string })
  |> iterator.take(n)
  |> iterator.map(fn(v) {
    let assert [a, s] = v |> string.split(" ")
    #(a |> int.parse |> result.unwrap(0), s)
  })
  |> iterator.to_list
}

fn find_init_pos(cmds: List(#(Int, String)), s: String) -> Int {
  cmds
  |> list.find(fn(c) { c.1 == s })
  |> result.unwrap(#(0, s))
  |> pair.first
}

pub fn main() {
  let in = stdin.stdin()
  let n = in |> read_int
  let cmds = in |> read_int_string_lines(n)

  let init_pos = #(find_init_pos(cmds, "L"), find_init_pos(cmds, "R"))

  let ans =
    cmds
    |> list.fold(#(init_pos, 0), fn(acc, cmd) {
      let #(#(pl, pr), e) = acc
      let #(a, s) = cmd
      case s {
        "L" -> #(#(a, pr), e + int.absolute_value(a - pl))
        "R" -> #(#(pl, a), e + int.absolute_value(a - pr))
        _ -> acc
      }
    })
    |> pair.second

  ans |> int.to_string |> io.println
}
