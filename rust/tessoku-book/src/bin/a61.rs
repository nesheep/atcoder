use itertools::Itertools;
use proconio::input;

fn main() {
    input! {
        n: usize,
        m: usize,
        edges: [(usize, usize); m],
    }

    let mut graph = vec![vec![]; n + 1];
    for &(a, b) in edges.iter() {
        graph[a].push(b);
        graph[b].push(a);
    }

    for i in 1..=n {
        let s = graph[i].iter().sorted().map(ToString::to_string).join(", ");
        println!("{i}: {{{s}}}");
    }
}
