use std::collections::HashMap;

use itertools::Itertools;
use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [i64; n],
    }

    let mut m = HashMap::new();
    let mut b = vec![0; n];

    for i in 0..n {
        b[i] = match m.get(&a[i]) {
            Some(v) => *v,
            None => -1,
        };

        m.insert(a[i], i as i64 + 1);
    }

    let ans = b.iter().map(|v| v.to_string()).join(" ");
    println!("{ans}");
}
