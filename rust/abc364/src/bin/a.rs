use proconio::input;

fn main() {
    input! {
        n: usize,
        s: [String; n]
    }

    let ans = s
        .windows(2)
        .rev()
        .skip(1)
        .filter(|x| x[0] == x[1])
        .map(|x| x[0].clone())
        .all(|v| v.as_str() != "sweet");

    println!("{}", if ans { "Yes" } else { "No" });
}
