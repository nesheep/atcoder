import gleam/int
import gleam/list
import gleam/string

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

pub fn take_string(in: List(String)) -> #(String, List(String)) {
  let assert [v, ..rest] = in
  #(v, rest)
}

pub fn take_int(in: List(String)) -> #(Int, List(String)) {
  let #(v, rest) = take_string(in)
  let assert Ok(v) = int.parse(v)
  #(v, rest)
}
