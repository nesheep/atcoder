use proconio::input;

fn f(a: &[usize], t: usize) -> usize {
    a.iter().map(|&x| t / x).sum::<usize>()
}

fn main() {
    input! {
        n: usize,
        k: usize,
        a: [usize; n],
    }

    let ans = {
        let mut l = 0;
        let mut r = 1_000_000_000;
        loop {
            if l == r {
                break l;
            }
            let m = (l + r) / 2;
            if f(&a, m) < k {
                l = m + 1;
            } else {
                r = m;
            }
        }
    };

    println!("{ans}");
}
