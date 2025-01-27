use proconio::input;

fn get(a: &[usize], i: usize, j: usize) -> usize {
    a[i * (i + 1) / 2 + j] - 1
}

fn combine(a: &[usize], i: usize, j: usize) -> usize {
    if i >= j {
        get(a, i, j)
    } else {
        get(a, j, i)
    }
}

fn main() {
    input! {
        n: usize,
        a: [usize; n * (n + 1) / 2],
    }

    let ans = (0..n).fold(0, |i, j| combine(&a, i, j)) + 1;

    println!("{ans}");
}
