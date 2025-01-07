use itertools::Itertools;
use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! { s: Chars }

    let ans = s[1..]
        .iter()
        .fold((0i64, vec![]), |(cnt, mut ans), &c| match c {
            '-' => (cnt + 1, ans),
            '|' => {
                ans.push(cnt);
                (0, ans)
            }
            _ => (cnt, ans),
        })
        .1;

    println!("{}", ans.iter().join(" "));
}
