use proconio::input;

fn main() {
    input! {
        a1: usize,
        a2: usize,
        a3: usize,
        a4: usize,
        a5: usize,
    }

    let ans = match (a1, a2, a3, a4, a5) {
        (2, 1, 3, 4, 5) => true,
        (1, 3, 2, 4, 5) => true,
        (1, 2, 4, 3, 5) => true,
        (1, 2, 3, 5, 4) => true,
        _ => false,
    };

    println!("{}", if ans { "Yes" } else { "No" });
}
