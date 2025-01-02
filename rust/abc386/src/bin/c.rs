use proconio::input;
use proconio::marker::Chars;

fn solve0(s: &[char], t: &[char]) -> bool {
    let mut d = 0;
    for i in 0..s.len() {
        if s[i] != t[i] {
            d += 1;
        }
    }
    d <= 1
}

fn solve1(s: &[char], t: &[char]) -> bool {
    let mut f = false;
    for i in 0..s.len() {
        if !f && s[i] != t[i] {
            f = true;
        }
        if f && s[i] != t[i + 1] {
            return false;
        }
    }
    true
}

fn main() {
    input! {
        _k: i8,
        s: Chars,
        t: Chars,
    }

    let ans = match t.len() as i64 - s.len() as i64 {
        0 => solve0(&s, &t),
        1 => solve1(&s, &t),
        -1 => solve1(&t, &s),
        _ => false,
    };

    println!("{}", if ans { "Yes" } else { "No" });
}
