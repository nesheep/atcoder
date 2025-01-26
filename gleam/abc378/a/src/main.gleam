import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{None, Some}
import gleam/result
import gleam/string
import stdin

fn parse_input(in: String) -> List(Int) {
  let assert Ok(a) =
    in
    |> string.trim
    |> string.split(" ")
    |> list.map(int.parse)
    |> result.all
  a
}

pub fn main() {
  let a = stdin.read_all() |> parse_input

  let ans =
    list.fold(a, dict.new(), fn(d, k) {
      dict.upsert(d, k, fn(x) {
        case x {
          Some(v) -> v + 1
          None -> 1
        }
      })
    })
    |> dict.to_list
    |> list.map(fn(x) { x.1 / 2 })
    |> int.sum

  ans |> int.to_string |> io.println
}
