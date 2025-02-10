import gleam/int
import gleam/io
import gleam/option.{None, Some}
import gleam/pair
import gleam/result
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(x, _) = stdin.take_int(in)

  let ans =
    yielder.range(1, 50)
    |> yielder.scan(Some(#(1, 0)), fn(acc, i) {
      case acc {
        None -> None
        Some(prev) ->
          case prev.0 * i {
            n if n <= x -> Some(#(n, i))
            _ -> None
          }
      }
    })
    |> yielder.take_while(option.is_some)
    |> yielder.last
    |> result.map(option.to_result(_, Nil))
    |> result.flatten
    |> result.unwrap(#(0, 0))
    |> pair.second

  ans |> int.to_string |> io.println
}
