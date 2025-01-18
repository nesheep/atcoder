use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { ss: [Chars; 12] }

    let ans = ss
        .iter()
        .enumerate()
        .filter(|(i, s)| s.len() == i + 1)
        .count();

    println!("{ans}");
}
