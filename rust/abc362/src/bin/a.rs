use std::cmp;

use proconio::input;

fn main() {
    input! {
        r: i64,
        g: i64,
        b: i64,
        c: String,
    }

    let ans = match c.as_str() {
        "Red" => cmp::min(g, b),
        "Green" => cmp::min(r, b),
        "Blue" => cmp::min(r, g),
        _ => unreachable!(),
    };

    println!("{ans}");
}
