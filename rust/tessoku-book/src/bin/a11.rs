use std::cmp::Ordering::*;

use proconio::input;

fn main() {
    input! {
        n: usize,
        x: usize,
        a: [usize; n],
    }

    let ans = {
        let mut l = 0;
        let mut r = n - 1;
        loop {
            let m = (l + r) / 2;
            match Ord::cmp(&x, &a[m]) {
                Less => r = m - 1,
                Equal => break m,
                Greater => l = m + 1,
            }
        }
    };

    println!("{}", ans + 1);
}
