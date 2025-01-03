use proconio::input;

fn main() {
    input! {
        n: usize,
        x: i64,
        a: [i64; n],
    }

    let ans = if a.contains(&x) { "Yes" } else { "No" };
    println!("{ans}");
}
