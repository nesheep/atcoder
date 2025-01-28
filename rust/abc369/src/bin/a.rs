use proconio::input;

fn main() {
    input! {
        a: i64,
        b: i64,
    }

    let ans = match (a - b).abs() {
        0 => 1,
        d if d % 2 == 1 => 2,
        _ => 3,
    };

    println!("{ans}");
}
