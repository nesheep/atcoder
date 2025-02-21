use proconio::input;

fn main() {
    input! {
        n: i64,
        t: i64,
        a: i64,
    }

    let ans = n < t * 2 || n < a * 2;

    println!("{}", if ans { "Yes" } else { "No" });
}
