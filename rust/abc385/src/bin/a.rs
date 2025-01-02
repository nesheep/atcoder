use proconio::input;

fn solve(a: i32, b: i32, c: i32) -> bool {
    if a == b && b == c {
        true
    } else {
        let mut v = vec![a, b, c];
        v.sort();
        v[0] + v[1] == v[2]
    }
}

fn main() {
    input! {
        a: i32,
        b: i32,
        c: i32,
    }

    let ans = if solve(a, b, c) { "Yes" } else { "No" };
    println!("{ans}");
}
