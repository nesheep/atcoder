use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        n: usize,
        m: usize,
        s: [Chars; n],
        t: [Chars; m],
    }

    let mut ans = (0, 0);
    for a in 0..=n - m {
        for b in 0..=n - m {
            let mut f = true;
            for i in 0..m {
                for j in 0..m {
                    if s[a + i][b + j] != t[i][j] {
                        f = false;
                    }
                }
            }
            if f {
                ans = (a + 1, b + 1);
                break;
            }
        }
    }

    println!("{} {}", ans.0, ans.1);
}
