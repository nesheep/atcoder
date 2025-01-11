use std::collections::HashMap;

use proconio::input;

fn main() {
    input! {
        a1: i32,
        a2: i32,
        a3: i32,
        a4: i32,
    }

    let a = [a1, a2, a3, a4];
    let mut map = HashMap::new();
    a.into_iter().for_each(|x| *map.entry(x).or_insert(0) += 1);
    let ans = map.values().map(|&x| x / 2).sum::<i32>();

    println!("{ans}");
}
