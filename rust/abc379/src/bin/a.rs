use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { n: Chars }

    if let [a, b, c] = n[..] {
        println!("{b}{c}{a} {c}{a}{b}");
    }
}
