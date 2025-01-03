use proconio::input;

fn main() {
    input! {
        n: usize,
        a: [i64; n],
    }

    let mut ans = false;
    for (ix, &x) in a.iter().enumerate() {
        for (iy, &y) in a.iter().enumerate() {
            for (iz, &z) in a.iter().enumerate() {
                if ix != iy && ix != iz && iy != iz && x + y + z == 1000 {
                    ans = true;
                    break;
                }
            }
        }
    }

    println!("{}", if ans { "Yes" } else { "No" });
}
