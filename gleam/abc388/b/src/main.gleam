import gleam/int
import gleam/io
import gleam/result
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let #(d, in) = stdin.take_int(in)
  let take_snake = stdin.take_tuple2(stdin.take_int, stdin.take_int)
  let #(snakes, _) = stdin.take_list(take_snake, n)(in)

  let ans =
    yielder.range(1, d)
    |> yielder.map(fn(k) {
      snakes
      |> yielder.from_list
      |> yielder.map(fn(snake) {
        let #(t, l) = snake
        t * { l + k }
      })
      |> yielder.reduce(int.max)
      |> result.unwrap(0)
    })

  ans |> yielder.map(int.to_string) |> yielder.each(io.println)
}
