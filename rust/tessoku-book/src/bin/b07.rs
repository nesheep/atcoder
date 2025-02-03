use proconio::input;

fn main() {
    input! {
        t: usize,
        n: usize,
        ranges: [(usize, usize); n],
    }

    let mut a = vec![0; t + 1];
    for &(l, r) in ranges.iter() {
        a[l] += 1;
        a[r] -= 1;
    }
    for i in 1..=t {
        a[i] += a[i - 1];
    }

    a[..t].iter().for_each(|a| println!("{a}"));
}
