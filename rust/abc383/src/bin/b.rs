use proconio::input;
use proconio::marker::Chars;

fn distance(a: &(usize, usize), b: &(usize, usize)) -> usize {
    ((a.0 as i64 - b.0 as i64).abs() + (a.1 as i64 - b.1 as i64).abs()) as usize
}

fn main() {
    input! {
        h: usize,
        _w: usize,
        d: usize,
        s: [Chars; h],
    }

    let t = s
        .iter()
        .enumerate()
        .flat_map(|(x, r)| {
            r.iter().enumerate().map(move |(y, &c)| match c {
                '.' => Some((x, y)),
                _ => None,
            })
        })
        .filter(Option::is_some)
        .map(Option::unwrap)
        .collect::<Vec<_>>();

    let mut ans = 0;
    for a in t.iter() {
        for b in t.iter() {
            if a == b {
                continue;
            }
            let cnt = t
                .iter()
                .filter(|c| distance(c, a) <= d || distance(c, b) <= d)
                .count();
            ans = ans.max(cnt);
        }
    }

    println!("{ans}");
}
