import gleam/float
import gleam/int
import gleam/io
import gleam/iterator.{type Iterator}
import gleam/list
import gleam/pair
import gleam/result
import gleam/string

@external(javascript, "./js_ffi.mjs", "read_all")
fn do_read_all() -> String

pub fn read_all() -> String {
  do_read_all()
}

fn parse_positions(in: String) -> Iterator(#(Int, Int)) {
  in
  |> string.split("\n")
  |> iterator.from_list
  |> iterator.drop(1)
  |> iterator.take_while(fn(l) { l != "" })
  |> iterator.map(fn(l) {
    let assert [x, y] =
      l
      |> string.split(" ")
      |> list.map(fn(s) { s |> int.parse |> result.unwrap(0) })
    #(x, y)
  })
}

fn moving_cost(p: #(Int, Int), q: #(Int, Int)) -> Float {
  let ac = p.0 - q.0
  let bd = p.1 - q.1
  { ac * ac + bd * bd }
  |> int.to_float
  |> float.square_root
  |> result.unwrap(0.0)
}

pub fn main() {
  let in = read_all()
  let ps = in |> parse_positions

  let ans =
    ps
    |> iterator.append(iterator.single(#(0, 0)))
    |> iterator.fold(#(0.0, #(0, 0)), fn(acc, q) {
      let #(cost, p) = acc
      #(cost +. moving_cost(p, q), q)
    })
    |> pair.first

  ans |> float.to_string |> io.println
}
