use std::collections::HashSet;

use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [i64; n],
    }

    let ans = a
        .windows(3)
        .any(|v| v.iter().collect::<HashSet<_>>().len() == 1);

    println!("{}", if ans { "Yes" } else { "No" });
}
