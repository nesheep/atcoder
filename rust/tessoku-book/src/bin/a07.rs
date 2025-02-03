use proconio::input;

fn main() {
    input! {
        d: usize,
        n: usize,
        ranges: [(usize, usize); n],
    }

    let mut a = vec![0; d + 2];
    for &(l, r) in ranges.iter() {
        a[l] += 1;
        a[r + 1] -= 1;
    }
    for i in 1..=d {
        a[i] += a[i - 1];
    }

    a[1..=d].iter().for_each(|a| println!("{a}"));
}
