use proconio::input;

fn main() {
    input! { x: f64 }

    let ans = match x {
        38.0.. => 1,
        37.5.. => 2,
        _ => 3,
    };

    println!("{ans}");
}
