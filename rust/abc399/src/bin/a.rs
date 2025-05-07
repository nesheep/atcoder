use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { _: usize, s: Chars, t: Chars }
    let ans = s.iter().zip(t.iter()).filter(|(s, t)| s != t).count();
    println!("{ans}");
}
