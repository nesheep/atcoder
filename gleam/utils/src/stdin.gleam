import gleam/float
import gleam/int
import gleam/list
import gleam/string
import gleam/yielder

@external(javascript, "./js_ffi.mjs", "read_all")
pub fn read_all() -> String

pub fn lines() -> List(String) {
  read_all()
  |> string.trim
  |> string.split("\n")
}

pub fn values() -> List(String) {
  lines()
  |> list.flat_map(fn(l) { string.split(l, " ") })
}

type TakeFn(a) =
  fn(List(String)) -> #(a, List(String))

pub fn take_string(in: List(String)) -> #(String, List(String)) {
  let assert [v, ..rest] = in
  #(v, rest)
}

pub fn take_int(in: List(String)) -> #(Int, List(String)) {
  let #(v, rest) = take_string(in)
  let assert Ok(v) = int.parse(v)
  #(v, rest)
}

pub fn take_float(in: List(String)) -> #(Float, List(String)) {
  let #(v, rest) = take_string(in)
  let assert Ok(v) = float.parse(v)
  #(v, rest)
}

pub fn take_tuple2(f0: TakeFn(a), f1: TakeFn(b)) -> TakeFn(#(a, b)) {
  fn(in: List(String)) -> #(#(a, b), List(String)) {
    let #(v0, in) = f0(in)
    let #(v1, in) = f1(in)
    #(#(v0, v1), in)
  }
}

pub fn take_list(f: TakeFn(a), n: Int) -> TakeFn(List(a)) {
  fn(in: List(String)) -> #(List(a), List(String)) {
    let assert Ok(#(l, in)) =
      yielder.iterate(#([], in), fn(x) {
        let #(l, in) = x
        let #(v, in) = f(in)
        #([v, ..l], in)
      })
      |> yielder.take(n)
      |> yielder.last
    let #(v, in) = f(in)
    let l = [v, ..l]
    #(list.reverse(l), in)
  }
}
