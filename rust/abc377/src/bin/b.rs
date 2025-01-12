use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { s: [Chars; 8] }

    let mut c = [true; 8];
    let mut r = [true; 8];
    for x in 0..8 {
        for y in 0..8 {
            if s[x][y] == '#' {
                c[x] = false;
                r[y] = false;
            }
        }
    }

    let mut ans = 0;
    for x in 0..8 {
        for y in 0..8 {
            if c[x] && r[y] {
                ans += 1;
            }
        }
    }

    println!("{ans}");
}
