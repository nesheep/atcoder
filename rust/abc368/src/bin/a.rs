use proconio::input;

fn main() {
    input! {
        n: usize,
        k: usize,
        a: [usize; n],
    }

    let (l, r) = a.split_at(n - k);
    let b = r.iter().chain(l.iter());

    let ans = b.map(|&x| x.to_string()).collect::<Vec<_>>().join(" ");
    println!("{ans}");
}
