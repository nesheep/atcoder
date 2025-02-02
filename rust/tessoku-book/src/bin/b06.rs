use std::cmp::Ordering::*;

use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [usize; n],
        q: usize,
        questions: [(usize, usize); q],
    }

    let mut b = vec![0; n + 1];
    for i in 0..n {
        b[i + 1] = b[i] + a[i];
    }

    let ans = questions
        .iter()
        .map(|&(l, r)| ((b[r] - b[l - 1]) * 2).cmp(&(r - l + 1)));

    ans.for_each(|a| {
        println!(
            "{}",
            match a {
                Greater => "win",
                Equal => "draw",
                Less => "lose",
            }
        )
    });
}
