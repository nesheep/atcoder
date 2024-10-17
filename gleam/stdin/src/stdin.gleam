import gleam/io

@external(javascript, "./js_ffi.mjs", "read_all")
fn read_all() -> String

pub fn main() {
  read_all() |> io.print
}
