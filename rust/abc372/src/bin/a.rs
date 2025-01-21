use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { s: Chars }
    let ans = s.iter().filter(|&&c| c != '.').collect::<String>();
    println!("{ans}");
}
