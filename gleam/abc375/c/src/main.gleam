import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/iterator
import gleam/result
import gleam/string

@external(javascript, "./js_ffi.mjs", "read_all")
fn read_all() -> String

fn parse_input(in: String) -> #(Int, List(String)) {
  let assert [first, ..rest] = in |> string.trim |> string.split("\n")
  let n = first |> int.parse |> result.unwrap(0)
  #(n, rest)
}

fn grid_dict(a: List(String)) -> Dict(#(Int, Int), String) {
  a
  |> iterator.from_list
  |> iterator.index
  |> iterator.map(fn(s) {
    let #(l, i) = s
    l
    |> string.to_graphemes
    |> iterator.from_list
    |> iterator.index
    |> iterator.map(fn(t) {
      let #(c, j) = t
      #(#(i, j), c)
    })
  })
  |> iterator.flatten
  |> iterator.to_list
  |> dict.from_list
}

fn operation_cnt(p: #(Int, Int), n: Int) -> Int {
  let #(i, j) = p
  int.min(int.min(i, n - 1 - i), int.min(j, n - 1 - j)) + 1
}

fn operate(p: #(Int, Int), n: Int, cnt: Int) -> #(Int, Int) {
  let #(i, j) = p
  case cnt % 4 {
    0 -> #(i, j)
    1 -> #(n - 1 - j, i)
    2 -> #(n - 1 - i, n - 1 - j)
    3 -> #(j, n - 1 - i)
    _ -> #(-1, -1)
  }
}

pub fn main() {
  let in = read_all()
  let #(n, a) = in |> parse_input

  let d = grid_dict(a)

  let ans =
    iterator.range(0, n - 1)
    |> iterator.map(fn(i) {
      iterator.range(0, n - 1)
      |> iterator.map(fn(j) {
        let p = #(i, j)
        p
        |> operation_cnt(n)
        |> operate(p, n, _)
        |> dict.get(d, _)
        |> result.unwrap("?")
      })
      |> iterator.to_list
      |> string.join("")
    })
    |> iterator.to_list
    |> string.join("\n")

  ans |> io.println
}
