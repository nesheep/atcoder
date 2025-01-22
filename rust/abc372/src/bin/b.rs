use std::iter;

use itertools::Itertools;
use proconio::input;

fn main() {
    input! { mut m: i64 }

    let ans = iter::from_fn(move || {
        (0..11)
            .map(|x| 3i64.pow(x))
            .enumerate()
            .rev()
            .find_map(|(i, x)| {
                if x <= m {
                    m -= x;
                    Some(i)
                } else {
                    None
                }
            })
    })
    .collect::<Vec<_>>();

    println!("{}", ans.len());
    println!("{}", ans.iter().map(ToString::to_string).join(" "));
}
