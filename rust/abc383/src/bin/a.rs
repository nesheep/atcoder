use proconio::input;

fn main() {
    input! {
        n: usize,
        tv: [(i64, i64); n],
    }

    let ans = tv
        .iter()
        .fold((0, 0), |(acc, prev), (t, v)| {
            ((acc - (t - prev)).max(0) + v, *t)
        })
        .0;

    println!("{ans}");
}
