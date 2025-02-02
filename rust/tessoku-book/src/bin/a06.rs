use proconio::input;

fn main() {
    input! {
        n: usize,
        q: usize,
        a: [usize; n],
        questions: [(usize, usize); q],
    }

    let mut b = vec![0; n + 1];
    for i in 0..n {
        b[i + 1] = b[i] + a[i];
    }

    let ans = questions.iter().map(|&(l, r)| b[r] - b[l - 1]);

    ans.for_each(|a| println!("{a}"));
}
