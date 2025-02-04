use itertools::Itertools;
use proconio::input;

fn main() {
    input! {
        h: usize,
        w: usize,
        n: usize,
        ranges: [(usize, usize, usize, usize); n]
    }

    let mut g = vec![vec![0i64; w + 2]; h + 2];
    for &(x1, y1, x2, y2) in ranges.iter() {
        g[x1][y1] += 1;
        g[x1][y2 + 1] -= 1;
        g[x2 + 1][y1] -= 1;
        g[x2 + 1][y2 + 1] += 1;
    }
    for i in 1..=h {
        for j in 1..=w {
            g[i][j] += g[i - 1][j] + g[i][j - 1] - g[i - 1][j - 1];
        }
    }

    g[1..=h]
        .iter()
        .map(|r| r[1..=w].iter().join(" "))
        .for_each(|r| println!("{r}"));
}
