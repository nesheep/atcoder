use itertools::Itertools;
use proconio::input;
use proconio::marker::Chars;

#[derive(Clone, Copy, Debug)]
enum State {
    None,
    Zero(usize),
    One(usize),
}

fn next_state(state: State, c: char) -> State {
    match state {
        State::None => match c {
            '0' => State::Zero(1),
            '1' => State::One(1),
            _ => unreachable!(),
        },
        State::Zero(cnt) => match c {
            '0' => State::Zero(cnt + 1),
            '1' => State::One(1),
            _ => unreachable!(),
        },
        State::One(cnt) => match c {
            '0' => State::Zero(1),
            '1' => State::One(cnt + 1),
            _ => unreachable!(),
        },
    }
}

fn main() {
    input! {
        _n: usize,
        k: usize,
        s: Chars,
    }

    let mut v = s
        .iter()
        .enumerate()
        .fold((State::None, vec![]), |(state, mut v), (i, &c)| {
            let next = next_state(state, c);
            match (state, next) {
                (State::Zero(_), State::One(_)) => v.push(state),
                (State::One(_), State::Zero(_)) => v.push(state),
                _ => (),
            }
            if i == s.len() - 1 {
                v.push(next);
            }
            (next, v)
        })
        .1;

    match v.first().unwrap() {
        State::Zero(_) => v.swap(k * 2 - 1, k * 2 - 2),
        State::One(_) => v.swap(k * 2 - 2, k * 2 - 3),
        _ => unreachable!(),
    };

    let ans = v
        .iter()
        .map(|state| match state {
            &State::Zero(cnt) => "0".repeat(cnt),
            &State::One(cnt) => "1".repeat(cnt),
            _ => unreachable!(),
        })
        .join("");

    println!("{ans}");
}
