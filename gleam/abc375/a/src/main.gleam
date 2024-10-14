import gleam/int
import gleam/io
import gleam/iterator
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

fn do_count(s: String, acc: Int) -> Int {
  case s {
    "#.#" <> rest -> do_count("#" <> rest, acc + 1)
    "#." <> rest | ".." <> rest -> do_count(rest, acc)
    "#" <> rest | "." <> rest -> do_count(rest, acc)
    _ -> acc
  }
}

pub fn main() {
  let in = stdin.stdin()
  let _ = in |> read_int
  let s = in |> read_string

  let ans = s |> do_count(0)

  ans |> int.to_string |> io.println
}
