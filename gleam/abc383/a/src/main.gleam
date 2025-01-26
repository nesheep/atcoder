import gleam/int
import gleam/io
import gleam/list
import gleam/pair
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let take_record = stdin.take_tuple2(stdin.take_int, stdin.take_int)
  let #(l, _) = stdin.take_list(take_record, n)(in)

  let ans =
    list.fold(l, #(0, 0), fn(acc_prev, t_v) {
      let #(acc, prev) = acc_prev
      let #(t, v) = t_v
      #(int.max(acc - t + prev, 0) + v, t)
    })
    |> pair.first

  ans |> int.to_string |> io.println
}
