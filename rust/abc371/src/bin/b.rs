use std::collections::HashSet;

use proconio::input;

fn main() {
    input! {
        n: usize,
        m: usize,
        children: [(usize, char); m],
    }

    let ans = children
        .iter()
        .scan(HashSet::with_capacity(n), |s, &(a, b)| {
            if b == 'M' && !s.contains(&a) {
                s.insert(a);
                Some(true)
            } else {
                Some(false)
            }
        });

    ans.for_each(|a| println!("{}", if a { "Yes" } else { "No" }));
}
