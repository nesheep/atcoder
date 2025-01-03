use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { n: Chars }

    let ans = n
        .iter()
        .map(|&x| x.to_digit(10).unwrap())
        .rev()
        .enumerate()
        .map(|(i, x)| x as i64 * 1 << i)
        .sum::<i64>();

    println!("{ans}");
}
