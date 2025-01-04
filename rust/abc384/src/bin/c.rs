use std::char;
use std::cmp::Ordering;

use itertools::Itertools;
use proconio::input;

fn main() {
    input! {
        a: usize,
        b: usize,
        c: usize,
        d: usize,
        e: usize,
    }

    let v = vec![a, b, c, d, e];

    let ans = (1..1 << 5)
        .map(|x| {
            let sum = (0..5)
                .map(|i| if x & (1 << i) != 0 { v[i] } else { 0 })
                .sum::<usize>();
            let name = (0..5)
                .map(|i| {
                    if x & (1 << i) != 0 {
                        char::from_u32('A' as u32 + i)
                    } else {
                        None
                    }
                })
                .filter(Option::is_some)
                .map(Option::unwrap)
                .collect::<String>();
            (name, sum)
        })
        .sorted_by(|(aname, asum), (bname, bsum)| match bsum.cmp(asum) {
            Ordering::Equal => aname.cmp(bname),
            o => o,
        });

    ans.for_each(|(name, _)| println!("{name}"));
}
