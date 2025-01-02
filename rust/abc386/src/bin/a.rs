use proconio::input;

use std::collections::HashSet;

fn main() {
    input! {
        a: i8,
        b: i8,
        c: i8,
        d: i8,
    }

    let s = HashSet::from([a, b, c, d]);
    let ans = if s.len() == 2 { "Yes" } else { "No" };
    println!("{ans}");
}
