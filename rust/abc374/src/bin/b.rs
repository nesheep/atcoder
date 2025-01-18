use std::iter;

use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        s: Chars,
        t: Chars,
    }

    let l = s.len().max(t.len());

    let si = s.iter().chain(iter::repeat(&'.')).take(l);
    let ti = t.iter().chain(iter::repeat(&'.')).take(l);

    let ans = si
        .zip(ti)
        .position(|(&s, &t)| s != t)
        .map(|i| i + 1)
        .unwrap_or(0);

    println!("{ans}");
}
