use std::collections::HashSet;

use proconio::input;

fn main() {
    input! {
        n: i64,
        m: usize,
        pieces: [(i64, i64); m],
    }

    let ds = [
        (2, 1),
        (1, 2),
        (-1, 2),
        (-2, 1),
        (-2, -1),
        (-1, -2),
        (1, -2),
        (2, -1),
    ];

    let m = pieces
        .iter()
        .flat_map(|p| ds.iter().map(move |d| (p.0 + d.0, p.1 + d.1)))
        .filter(|p| p.0 >= 1 && p.0 <= n && p.1 >= 1 && p.1 <= n)
        .chain(pieces.iter().cloned())
        .collect::<HashSet<_>>();

    let ans = n * n - m.len() as i64;
    println!("{ans}");
}
