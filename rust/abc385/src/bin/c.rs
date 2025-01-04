use proconio::input;

fn main() {
    input! {
        n: usize,
        h: [i64; n],
    }

    let mut ans = 1;
    for i in 1..n {
        let mut dp = vec![1; n];
        for j in 0..n {
            if j >= i && h[j] == h[j - i] {
                dp[j] = dp[j - i] + 1;
            }
        }
        ans = ans.max(*dp.iter().max().unwrap());
    }

    println!("{ans}");
}
