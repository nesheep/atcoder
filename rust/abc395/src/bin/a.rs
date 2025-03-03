use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [usize; n],
    }

    let ans = a.windows(2).all(|x| x[0] < x[1]);

    println!("{}", if ans { "Yes" } else { "No" });
}
