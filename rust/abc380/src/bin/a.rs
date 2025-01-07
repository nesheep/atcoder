use proconio::input;
use proconio::marker::Chars;

fn count(s: &[char], t: char) -> usize {
    s.iter().filter(|&&c| c == t).count()
}

fn main() {
    input! { s: Chars }

    let ans = match (count(&s, '1'), count(&s, '2'), count(&s, '3')) {
        (1, 2, 3) => true,
        _ => false,
    };

    println!("{}", if ans { "Yes" } else { "No" });
}
