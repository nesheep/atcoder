import gleam/float
import gleam/int
import gleam/io
import gleam/iterator.{type Iterator}
import gleam/list
import gleam/pair
import gleam/result
import gleam/string
import stdin

fn read_string(in: iterator.Iterator(String)) -> String {
  in |> iterator.first |> result.unwrap("") |> string.trim
}

fn read_strings(in: iterator.Iterator(String)) -> List(String) {
  in |> read_string |> string.split(" ")
}

fn parse_int(v: String) -> Int {
  v |> int.parse |> result.unwrap(0)
}

fn read_int(in: iterator.Iterator(String)) -> Int {
  in |> read_string |> parse_int
}

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_strings |> list.map(parse_int)
}

fn read_positions(
  in: iterator.Iterator(String),
  n: Int,
) -> Iterator(#(Float, Float)) {
  iterator.repeatedly(fn() {
    let assert [x, y] = in |> read_ints |> list.map(int.to_float)
    #(x, y)
  })
  |> iterator.take(n)
}

fn moving_cost(p: #(Float, Float), q: #(Float, Float)) -> Float {
  let ac = p.0 -. q.0
  let bd = p.1 -. q.1
  float.square_root(ac *. ac +. bd *. bd) |> result.unwrap(0.0)
}

pub fn main() {
  let in = stdin.stdin()
  let n = in |> read_int
  let ps = in |> read_positions(n)

  let ans =
    ps
    |> iterator.append(iterator.single(#(0.0, 0.0)))
    |> iterator.fold(#(0.0, #(0.0, 0.0)), fn(acc, q) {
      let #(cost, p) = acc
      #(cost +. moving_cost(p, q), q)
    })
    |> pair.first

  ans |> float.to_string |> io.println
}
