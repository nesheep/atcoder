use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        n: usize,
        s: Chars,
    }

    let ans = match n % 2 {
        0 => false,
        _ => (0..n)
            .map(|i| match (n + 1) / 2 - 1 {
                m if i < m => '1',
                m if i > m => '2',
                _ => '/',
            })
            .zip(s.iter())
            .all(|(a, b)| a == *b),
    };

    println!("{}", if ans { "Yes" } else { "No" });
}
