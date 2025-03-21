import gleam/int
import gleam/io
import stdin

fn parse_input(in: List(String)) -> String {
  let #(s, _) = stdin.take_string(in)
  s
}

fn solve(s: String) -> Bool {
  case s {
    "ABC316" -> False
    "ABC" <> n -> {
      let assert Ok(n) = int.parse(n)
      n != 0 && n < 350
    }
    _ -> panic
  }
}

fn output(b: Bool) -> String {
  case b {
    False -> "No"
    True -> "Yes"
  }
}

pub fn main() {
  stdin.values()
  |> parse_input
  |> solve
  |> output
  |> io.println
}
