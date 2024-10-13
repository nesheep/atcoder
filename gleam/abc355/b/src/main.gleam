import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
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

fn read_ints(in: iterator.Iterator(String)) -> List(Int) {
  in |> read_strings |> list.map(parse_int)
}

type Element {
  A(value: Int)
  B(value: Int)
}

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_ints
  let a = in |> read_ints |> list.map(fn(v) { A(v) })
  let b = in |> read_ints |> list.map(fn(v) { B(v) })

  let c =
    list.append(a, b)
    |> list.sort(fn(x, y) { int.compare(x.value, y.value) })

  let ans =
    c
    |> list.window_by_2
    |> list.any(fn(v) {
      case v {
        #(A(_), A(_)) -> True
        _ -> False
      }
    })

  case ans {
    True -> "Yes"
    False -> "No"
  }
  |> io.println
}
