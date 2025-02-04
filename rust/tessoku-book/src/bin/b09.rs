use proconio::input;

fn main() {
    input! {
        n: usize,
        ranges: [(usize, usize, usize, usize); n]
    }

    let mut g = vec![vec![0i64; 1501]; 1501];
    for &(x1, y1, x2, y2) in ranges.iter() {
        g[x1][y1] += 1;
        g[x1][y2] -= 1;
        g[x2][y1] -= 1;
        g[x2][y2] += 1;
    }
    for i in 1..=1500 {
        g[i][0] += g[i - 1][0];
        g[0][i] += g[0][i - 1];
    }
    for i in 1..=1500 {
        for j in 1..=1500 {
            g[i][j] += g[i - 1][j] + g[i][j - 1] - g[i - 1][j - 1];
        }
    }

    let ans = g
        .into_iter()
        .flat_map(|g| g.into_iter())
        .filter(|&g| g > 0)
        .count();

    println!("{ans}");
}
