use proconio::input;

fn main() {
    input! {
        n: usize,
        k: usize,
        a: [usize; n],
    }

    let ans = a
        .iter()
        .scan(0, |r, &x| {
            while *r < n - 1 && a[*r + 1] - x <= k {
                *r += 1;
            }
            Some(*r)
        })
        .enumerate()
        .map(|(l, r)| r - l)
        .sum::<usize>();

    println!("{ans}");
}
