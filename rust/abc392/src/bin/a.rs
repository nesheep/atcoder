use proconio::input;

fn main() {
    input! {
        a1: usize,
        a2: usize,
        a3: usize,
    }

    let ans = [a1 * a2 == a3, a1 * a3 == a2, a2 * a3 == a1]
        .iter()
        .any(|&x| x);

    println!("{}", if ans { "Yes" } else { "No" });
}
