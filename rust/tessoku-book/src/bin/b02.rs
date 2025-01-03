use proconio::input;

fn main() {
    input! { a: i64, b: i64 }
    let ans = (a..=b).any(|v| 100 % v == 0);
    println!("{}", if ans { "Yes" } else { "No" });
}
