use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        _n: usize,
        s: Chars,
    }

    let ans = s
        .windows(3)
        .filter(|&x| match x {
            ['#', '.', '#'] => true,
            _ => false,
        })
        .count();

    println!("{ans}");
}
