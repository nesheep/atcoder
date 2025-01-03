use proconio::input;

fn main() {
    input! { n: i64, k: i64 }

    let mut ans = 0;
    for x in 1..=n {
        for y in 1..=n {
            let z = k - x - y;
            if z >= 1 && z <= n {
                ans += 1;
            }
        }
    }

    println!("{ans}");
}
