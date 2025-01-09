use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        _n: usize,
        k: usize,
        s: Chars,
    }

    let ans = s
        .iter()
        .fold((0usize, 0usize), |(cnt, ans), &c| match c {
            'O' if cnt + 1 >= k => (0, ans + 1),
            'O' => (cnt + 1, ans),
            _ => (0, ans),
        })
        .1;

    println!("{ans}");
}
