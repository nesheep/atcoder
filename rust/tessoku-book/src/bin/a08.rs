use proconio::input;

fn main() {
    input! {
        h: usize,
        w: usize,
        x: [[usize; w]; h],
        q: usize,
        questions: [(usize, usize, usize, usize); q],
    }

    let mut t = vec![vec![0; w + 1]; h + 1];
    for i in 0..h {
        for j in 0..w {
            t[i + 1][j + 1] = t[i + 1][j] + x[i][j];
        }
    }
    for i in 0..h {
        for j in 0..w {
            t[i + 1][j + 1] += t[i][j + 1];
        }
    }

    let ans = questions
        .iter()
        .map(|&(y1, x1, y2, x2)| t[y2][x2] - t[y1 - 1][x2] - t[y2][x1 - 1] + t[y1 - 1][x1 - 1]);

    ans.for_each(|a| println!("{a}"));
}
