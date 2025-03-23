use proconio::input;

fn main() {
    input! { n: usize }
    let a = "-".repeat((n - 1) / 2);
    let b = if n % 2 == 0 { "==" } else { "=" };
    println!("{a}{b}{a}");
}
