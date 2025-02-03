use proconio::input;

fn main() {
    input! {
        n: usize,
        points: [(usize, usize); n],
        q: usize,
        questions: [(usize, usize, usize, usize); q],
    }

    let mut t = vec![vec![0; 1501]; 1501];
    for &(x, y) in points.iter() {
        t[x][y] += 1;
    }
    for i in 0..1500 {
        for j in 0..1500 {
            t[i + 1][j + 1] += t[i + 1][j];
        }
    }
    for i in 0..1500 {
        for j in 0..1500 {
            t[i + 1][j + 1] += t[i][j + 1];
        }
    }

    let ans = questions
        .iter()
        .map(|&(y1, x1, y2, x2)| t[y2][x2] - t[y1 - 1][x2] - t[y2][x1 - 1] + t[y1 - 1][x1 - 1]);

    ans.for_each(|a| println!("{a}"));
}
