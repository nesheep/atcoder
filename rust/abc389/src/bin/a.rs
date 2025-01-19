use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { s: Chars }

    let (a, b) = match s[..] {
        [a, 'x', b] => (a.to_digit(10).unwrap(), b.to_digit(10).unwrap()),
        _ => unreachable!(),
    };

    println!("{}", a * b);
}
