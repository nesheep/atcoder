use proconio::input;

fn main() {
    input! { s: String }
    let ans = s.ends_with("san");
    println!("{}", if ans { "Yes" } else { "No" });
}
