use proconio::input;

fn main() {
    input! {
        n: usize,
        m: usize,
        a: [i64; n],
        b: [i64; m],
    }

    let max = *a.iter().max().unwrap().max(b.iter().max().unwrap());
    let mut t = vec![-1; max as usize + 1];

    for (i, &x) in a.iter().enumerate().rev() {
        t[x as usize] = i as i64 + 1;
    }

    for i in 1..t.len() {
        match (t[i - 1], t[i]) {
            (-1, _) => (),
            (_, -1) => t[i] = t[i - 1],
            (a, b) => t[i] = a.min(b),
        }
    }

    let ans = b.iter().map(|&x| t[x as usize]);

    ans.for_each(|x| println!("{}", x));
}
