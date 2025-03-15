use proconio::input;

fn distance(p1: (i64, i64), p2: (i64, i64)) -> i64 {
    (p1.0 - p2.0).pow(2) + (p1.1 - p2.1).pow(2)
}

fn main() {
    input! {
        p1: (i64, i64),
        p2: (i64, i64),
        p3: (i64, i64),
    }

    let mut d = [distance(p1, p2), distance(p1, p3), distance(p2, p3)];
    d.sort();
    let ans = d[1] + d[0] == d[2];

    println!("{}", if ans { "Yes" } else { "No" });
}
