use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        _n: usize,
        c1: char,
        c2: char,
        s: Chars,
    }

    let ans = s
        .iter()
        .map(|&c| if c == c1 { c1 } else { c2 })
        .collect::<String>();

    println!("{ans}");
}
