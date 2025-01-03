use std::collections::HashSet;

use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        h: usize,
        w: usize,
        x: usize,
        y: usize,
        s: [Chars; h],
        t: Chars,
    }

    let mut set = HashSet::new();
    let mut p = (x - 1, y - 1);

    for c in t {
        let next = match c {
            'U' => (p.0.checked_sub(1).unwrap_or(p.0), p.1),
            'D' => (p.0 + 1, p.1),
            'L' => (p.0, p.1.checked_sub(1).unwrap_or(p.1)),
            'R' => (p.0, p.1 + 1),
            _ => p,
        };
        if next.0 >= h || next.1 >= w || s[next.0][next.1] == '#' {
            continue;
        }
        p = next;
        if s[p.0][p.1] == '@' {
            set.insert(p);
        }
    }

    println!("{} {} {}", p.0 + 1, p.1 + 1, set.len());
}
