use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { mut s: Chars }
    s.sort();
    let ans = s.iter().collect::<String>() == "ABC";
    println!("{}", if ans { "Yes" } else { "No" });
}
