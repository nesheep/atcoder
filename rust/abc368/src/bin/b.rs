use proconio::input;

fn main() {
    input! {
        n: usize,
        mut a: [usize; n],
    }

    let mut ans = 0;
    loop {
        let cnt = a.iter().filter(|&&v| v > 0).count();
        match cnt {
            0 | 1 => break,
            _ => {
                a.sort_by(|a, b| b.cmp(a));
                a[0] -= 1;
                a[1] -= 1;
                ans += 1;
            }
        }
    }

    println!("{ans}");
}
