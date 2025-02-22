use proconio::input;

fn main() {
    input! {
        n: usize,
    }

    let ans = match n {
        n if n % 4 != 0 => 365,
        n if n % 100 != 0 => 366,
        n if n % 400 != 0 => 365,
        _ => 366,
    };

    println!("{ans}");
}
