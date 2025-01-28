import gleam/int
import gleam/io
import gleam/yielder
import stdin

pub fn main() {
  let in = stdin.values()
  let #(n, in) = stdin.take_int(in)
  let #(r, in) = stdin.take_int(in)
  let take_record = stdin.take_tuple2(stdin.take_int, stdin.take_int)
  let #(contests, _) = stdin.take_list(take_record, n)(in)

  let ans =
    contests
    |> yielder.from_list
    |> yielder.fold(r, fn(t, contest) {
      let #(d, a) = contest
      case d {
        1 if t >= 1600 && t < 2800 -> t + a
        2 if t >= 1200 && t < 2400 -> t + a
        _ -> t
      }
    })

  ans |> int.to_string |> io.println
}
