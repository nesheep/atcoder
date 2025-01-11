use proconio::input;

fn calc(d: i64, (q, r): (i64, i64)) -> i64 {
    (d / q + if d % q > r { 1 } else { 0 }) * q + r
}

fn main() {
    input! {
        n: usize,
        rules: [(i64, i64); n],
        q: usize,
        questions: [(usize, i64); q],
    }

    let ans = questions.iter().map(|q| calc(q.1, rules[q.0 - 1]));

    ans.for_each(|a| println!("{a}"));
}
