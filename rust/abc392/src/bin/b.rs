use itertools::Itertools;
use proconio::input;

fn main() {
    input! {
        n: usize,
        m: usize,
        a: [usize; m],
    }

    let ans = (1..=n)
        .filter(|&x| a.iter().all(|&y| x != y))
        .collect::<Vec<_>>();

    println!("{}", ans.len());
    println!("{}", ans.iter().join(" "));
}
