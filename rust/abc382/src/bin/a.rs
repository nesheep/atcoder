use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        n: usize,
        d: usize,
        s: Chars,
    }

    let c = s.iter().filter(|&&x| x == '@').count();
    let ans = n - c + d;

    println!("{ans}");
}
