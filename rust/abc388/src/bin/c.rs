use std::collections::HashMap;

use itertools::Itertools;
use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [usize; n],
    }

    let mut m = HashMap::new();
    a.iter().for_each(|&x| {
        m.insert(x * 2, 0usize);
        *m.entry(x).or_insert(0) += 1;
    });

    let mut v = m
        .into_iter()
        .sorted_by(|(x, _), (y, _)| y.cmp(x))
        .collect::<Vec<_>>();

    for i in 1..v.len() {
        v[i].1 += v[i - 1].1;
    }

    let m = v.into_iter().collect::<HashMap<_, _>>();

    let ans = a
        .iter()
        .fold(0, |acc, &x| acc + m.get(&(x * 2)).unwrap_or(&0));

    println!("{ans}");
}
