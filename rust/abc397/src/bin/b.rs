use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { s: Chars }

    let ans = match s.iter().fold((0, 0), |(i, a), &c| match (i, c) {
        (i, 'i') if i % 2 == 0 => (i + 1, a),
        (i, 'o') if i % 2 == 1 => (i + 1, a),
        _ => (i + 2, a + 1),
    }) {
        (i, a) if i % 2 == 1 => a + 1,
        (_, a) => a,
    };

    println!("{ans}");
}
