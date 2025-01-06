use proconio::input;
use proconio::marker::Chars;

enum State {
    None,
    One(i64),
    Slash(i64),
    Two(i64, i64),
}

fn score(prev: i64, cnt: i64) -> i64 {
    prev.max(cnt * 2 + 1)
}

fn main() {
    input! {
        _n: usize,
        s: Chars,
    }

    let ans = s
        .iter()
        .fold((State::None, 1), |(state, acc), &c| match state {
            State::None => match c {
                '1' => (State::One(1), acc),
                _ => (State::None, acc),
            },
            State::One(ones) => match c {
                '1' => (State::One(ones + 1), acc),
                '/' => (State::Slash(ones), acc),
                _ => (State::None, acc),
            },
            State::Slash(ones) => match c {
                '1' => (State::One(1), acc),
                '2' => (State::Two(ones, 1), score(acc, 1)),
                _ => (State::None, acc),
            },
            State::Two(ones, twos) => match c {
                '1' => (State::One(1), acc),
                '2' => (State::Two(ones, twos + 1), score(acc, ones.min(twos + 1))),
                _ => (State::None, acc),
            },
        })
        .1;

    println!("{ans}");
}
