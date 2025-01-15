use std::collections::HashMap;

use itertools::Itertools;
use proconio::input;

fn distance(m: &HashMap<usize, usize>, a: usize, b: usize, z: usize) -> Option<usize> {
    let mut c = a;
    let mut d = 0;
    while c != b {
        c = m[&c];
        if c == z {
            return None;
        }
        d += 1;
    }
    Some(d)
}

fn main() {
    input! {
        n: usize,
        q: usize,
        cmds: [(char, usize); q],
    }

    let rm = (1..n + 1)
        .circular_tuple_windows()
        .take(n)
        .collect::<HashMap<_, _>>();

    let lm = (1..n + 1)
        .rev()
        .circular_tuple_windows()
        .take(n)
        .collect::<HashMap<_, _>>();

    let ans = cmds
        .iter()
        .fold(((1, 2), 0), |((l, r), acc), &(h, t)| match h {
            'L' => match (distance(&rm, l, t, r), distance(&lm, l, t, r)) {
                (Some(d), _) => ((t, r), acc + d),
                (_, Some(d)) => ((t, r), acc + d),
                _ => unreachable!(),
            },
            'R' => match (distance(&rm, r, t, l), distance(&lm, r, t, l)) {
                (Some(d), _) => ((l, t), acc + d),
                (_, Some(d)) => ((l, t), acc + d),
                _ => unreachable!(),
            },
            _ => unreachable!(),
        })
        .1;

    println!("{ans}");
}
