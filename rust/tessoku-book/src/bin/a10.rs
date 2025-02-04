use std::iter;

use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [usize; n],
        d: usize,
        ranges: [(usize, usize); d],
    }

    let max_from_l = iter::once(0)
        .chain(a.iter().scan(0, |acc, &v| {
            *acc = acc.clone().max(v);
            Some(*acc)
        }))
        .chain(iter::once(0))
        .collect::<Vec<_>>();

    let max_from_r = {
        let mut t = iter::once(0)
            .chain(a.iter().rev().scan(0, |acc, &v| {
                *acc = acc.clone().max(v);
                Some(*acc)
            }))
            .chain(iter::once(0))
            .collect::<Vec<_>>();
        t.reverse();
        t
    };

    let ans = ranges
        .iter()
        .map(|&(l, r)| max_from_l[l - 1].max(max_from_r[r + 1]));

    ans.for_each(|a| println!("{a}"));
}
