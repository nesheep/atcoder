use std::iter;

use itertools::Itertools;
use proconio::input;

fn cost((x1, y1): &(f64, f64), (x2, y2): &(f64, f64)) -> f64 {
    ((x1 - x2).powf(2.0) + (y1 - y2).powf(2.0)).sqrt()
}

fn main() {
    input! {
        n: usize,
        ps: [(f64, f64); n],
    }

    let o = (0.0, 0.0);

    let ans = iter::once(&o)
        .chain(ps.iter())
        .chain(iter::once(&o))
        .tuple_windows::<(_, _)>()
        .fold(0.0, |acc, (a, b)| acc + cost(a, b));

    println!("{ans:.20}");
}
