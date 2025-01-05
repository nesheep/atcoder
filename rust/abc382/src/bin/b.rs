use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        _n: usize,
        d: usize,
        s: Chars,
    }

    let ans = s
        .iter()
        .rev()
        .collect::<String>()
        .replacen("@", ".", d)
        .chars()
        .rev()
        .collect::<String>();

    println!("{ans}")
}
