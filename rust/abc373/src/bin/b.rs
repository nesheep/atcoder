use itertools::Itertools;
use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { s: Chars }

    let ans = s
        .iter()
        .enumerate()
        .sorted_by_key(|x| x.1)
        .map(|x| x.0 as i64)
        .tuple_windows()
        .fold(0, |acc, (a, b)| acc + (b - a).abs());

    println!("{ans}");
}
