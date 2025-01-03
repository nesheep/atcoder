use proconio::input;

fn main() {
    input! {
        n: usize,
        r: i64,
        results: [(i8, i64); n],
    }

    let ans = results.iter().fold(r, |t, &(d, a)| match d {
        1 if t >= 1600 && t < 2800 => t + a,
        2 if t >= 1200 && t < 2400 => t + a,
        _ => t,
    });

    println!("{ans}");
}
