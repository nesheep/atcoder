@external(javascript, "./js_ffi.mjs", "read_all")
fn do_read_all() -> String

pub fn read_all() -> String {
  do_read_all()
}
