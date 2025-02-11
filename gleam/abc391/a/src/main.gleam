import gleam/io
import stdin

pub fn main() {
  let in = stdin.values()
  let #(d, _) = stdin.take_string(in)

  let ans = case d {
    "N" -> "S"
    "E" -> "W"
    "W" -> "E"
    "S" -> "N"
    "NE" -> "SW"
    "NW" -> "SE"
    "SE" -> "NW"
    "SW" -> "NE"
    _ -> ""
  }

  io.println(ans)
}
