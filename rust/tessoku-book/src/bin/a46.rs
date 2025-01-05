use proconio::input;

fn dist(a: (i64, i64), b: (i64, i64)) -> i64 {
    (a.0 - b.0).pow(2) + (a.1 - b.1).pow(2)
}

fn solve(n: usize, cities: &[(i64, i64)]) -> Vec<usize> {
    let mut done = vec![false; n];
    let mut ans = vec![0; n + 1];

    done[0] = true;
    for i in 1..n {
        let mut dmin = std::i64::MAX;
        let mut next = 0;
        for j in 0..n {
            if !done[j] {
                let d = dist(cities[ans[i - 1]], cities[j]);
                if d < dmin {
                    dmin = d;
                    next = j;
                }
            }
        }
        done[next] = true;
        ans[i] = next;
    }

    ans
}

fn main() {
    input! {
        n: usize,
        cities: [(i64, i64); n],
    }

    let ans = solve(n, &cities);
    ans.iter().for_each(|a| println!("{}", a + 1));
}
