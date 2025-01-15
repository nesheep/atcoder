use proconio::input;

fn main() {
    input! {
        n: usize,
        c: i64,
        t: [i64; n],
    }

    let mut ans = 1;
    let mut ct = t[0];
    for i in 1..n {
        if t[i] - ct >= c {
            ans += 1;
            ct = t[i];
        }
    }

    println!("{ans}");
}
