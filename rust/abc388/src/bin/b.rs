use proconio::input;

fn main() {
    input! {
        n: usize,
        d: i64,
        snakes: [(i64, i64); n],
    }

    let ans = (1..=d).map(|k| snakes.iter().map(move |&(t, l)| t * (l + k)).max().unwrap());

    ans.for_each(|a| println!("{a}"));
}
