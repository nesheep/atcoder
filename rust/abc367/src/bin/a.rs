use proconio::input;

fn main() {
    input! {
        a: i64,
        b: i64,
        c: i64,
    }

    let ans = if b < c {
        a < b || c < a
    } else {
        c < a && a < b
    };

    println!("{}", if ans { "Yes" } else { "No" });
}
