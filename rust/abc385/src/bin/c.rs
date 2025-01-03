use proconio::input;

fn main() {
    input! {
        n: usize,
        h: [i64; n],
    }

    let mut ans = 1;
    for i in 1..n {
        for j in 0..i {
            let mut cnt = 0;
            let mut g = -1;
            let mut k = j;
            while k < n {
                if h[k] == g {
                    cnt += 1;
                    ans = ans.max(cnt);
                } else {
                    cnt = 1;
                    g = h[k];
                }
                k += i;
            }
        }
    }

    println!("{ans}");
}
