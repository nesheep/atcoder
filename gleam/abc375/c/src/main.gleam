import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/order
import gleam/pair
import gleam/result
import gleam/string

@external(javascript, "./js_ffi.mjs", "read_all")
fn read_all() -> String

fn parse_input(in: String) -> #(Int, List(String)) {
  let assert [first, ..rest] = in |> string.trim |> string.split("\n")
  let n = first |> int.parse |> result.unwrap(0)
  #(n, rest)
}

fn operation_cnt(p: #(Int, Int), n: Int) -> Int {
  let #(i, j) = p
  int.min(int.min(i, n - 1 - i), int.min(j, n - 1 - j)) + 1
}

fn operate(p: #(Int, Int), n: Int, cnt: Int) -> #(Int, Int) {
  let #(i, j) = p
  case cnt % 4 {
    0 -> #(i, j)
    1 -> #(j, n - 1 - i)
    2 -> #(n - 1 - i, n - 1 - j)
    3 -> #(n - 1 - j, i)
    _ -> #(-1, -1)
  }
}

pub fn main() {
  let in = read_all()
  let #(n, a) = in |> parse_input

  let ans =
    a
    |> iterator.from_list
    |> iterator.map(string.to_graphemes)
    |> iterator.map(iterator.from_list)
    |> iterator.map(iterator.index)
    |> iterator.index
    |> iterator.map(fn(e) {
      let #(l, i) = e
      l
      |> iterator.map(fn(f) {
        let #(c, j) = f
        #(#(i, j), c)
      })
    })
    |> iterator.flatten
    |> iterator.map(fn(e) {
      let #(p, c) = e
      let cnt = operation_cnt(p, n)
      #(operate(p, n, cnt), c)
    })
    |> iterator.to_list
    |> list.sort(fn(e, f) {
      let #(#(i1, j1), _) = e
      let #(#(i2, j2), _) = f
      case int.compare(i1, i2) {
        order.Eq -> int.compare(j1, j2)
        o -> o
      }
    })
    |> iterator.from_list
    |> iterator.map(pair.second)
    |> iterator.sized_chunk(n)
    |> iterator.map(string.join(_, ""))
    |> iterator.to_list
    |> string.join("\n")

  ans |> io.println
}
