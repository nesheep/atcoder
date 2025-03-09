import gleam/int
import gleam/io
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/pair
import gleam/yielder
import stdin

type Query {
  Query1(Int)
  Query2
}

fn take_query(in: List(String)) -> #(Query, List(String)) {
  let #(t, in) = stdin.take_int(in)
  case t {
    1 -> {
      let #(x, in) = stdin.take_int(in)
      #(Query1(x), in)
    }
    2 -> #(Query2, in)
    _ -> panic
  }
}

fn parse_input(in: List(String)) -> List(Query) {
  let #(n, in) = stdin.take_int(in)
  let #(queries, _) = stdin.take_list(take_query, n)(in)
  queries
}

fn run_query(
  acc: #(List(Int), Option(Int)),
  q: Query,
) -> #(List(Int), Option(Int)) {
  let #(acc, _) = acc
  case q {
    Query1(x) -> #([x, ..acc], None)
    Query2 ->
      case acc {
        [f, ..rest] -> #(rest, Some(f))
        [] -> #([], Some(0))
      }
  }
}

fn solve(queries: List(Query)) -> List(Int) {
  queries
  |> yielder.from_list
  |> yielder.scan(#([], None), run_query)
  |> yielder.map(pair.second)
  |> yielder.filter(option.is_some)
  |> yielder.map(option.unwrap(_, 0))
  |> yielder.to_list
}

fn print_int(i: Int) -> Nil {
  i |> int.to_string |> io.println
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> list.each(print_int)
}
