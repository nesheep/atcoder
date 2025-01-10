use std::iter;

use itertools::Itertools;
use proconio::input;

fn sum(x: i64) -> i64 {
    x * (x + 1) / 2
}

fn cost(x: i64, y: i64) -> i64 {
    sum(x) - sum(x - y)
}

fn main() {
    input! {
        n: usize,
        m: usize,
        x: [i64; m],
        a: [i64; m],
    }

    let x0 = x.iter().min().unwrap().clone();

    let it = x
        .into_iter()
        .zip(a.into_iter())
        .chain(iter::once((n as i64 + 1, 0)))
        .sorted_by(|a, b| Ord::cmp(&a.0, &b.0));

    let result = it
        .clone()
        .zip(it.skip(1))
        .try_fold((0, 0), |(c, r), ((x1, a1), (x2, _))| {
            let q = r + a1 - 1;
            let d = x2 - x1 - 1;
            let next_c = c + r + cost(q, d);
            let next_r = q - d;
            if next_r < 0 {
                None
            } else {
                Some((next_c, next_r))
            }
        });

    let ans = match result {
        Some((c, r)) => match (x0, r) {
            (1, 0) => c,
            _ => -1,
        },
        None => -1,
    };

    println!("{ans}");
}
