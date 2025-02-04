import gleam/dict
import gleam/int
import gleam/io
import gleam/pair
import gleam/result
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let #(q, in) = stdin.take_int(in)
  let #(a, in) = stdin.take_list(stdin.take_int, n)(in)
  let take_range = stdin.take_tuple2(stdin.take_int, stdin.take_int)
  let #(ranges, _) = stdin.take_list(take_range, q)(in)

  let b =
    yielder.single(0)
    |> yielder.append(yielder.from_list(a))
    |> yielder.scan(0, int.add)
    |> yielder.index
    |> yielder.map(pair.swap)
    |> yielder.to_list
    |> dict.from_list

  let ans =
    ranges
    |> yielder.from_list
    |> yielder.map(fn(l_r) {
      let #(l, r) = l_r
      use l <- result.try(dict.get(b, l - 1))
      use r <- result.try(dict.get(b, r))
      Ok(r - l)
    })
    |> yielder.map(result.unwrap(_, 0))

  ans
  |> yielder.map(int.to_string)
  |> yielder.each(io.println)
}
