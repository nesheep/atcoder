use std::collections::HashSet;

use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        s: Chars,
    }

    let ans = s
        .iter()
        .zip(s.iter().skip(1))
        .filter(|(a, b)| a == b)
        .map(|(a, _)| a)
        .collect::<HashSet<_>>()
        .len()
        * 2
        == s.len();

    println!("{}", if ans { "Yes" } else { "No" });
}
