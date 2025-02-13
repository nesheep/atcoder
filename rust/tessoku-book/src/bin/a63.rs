use std::collections::VecDeque;

use proconio::input;

fn edges_to_graph(edges: &[(usize, usize)], n: usize) -> Vec<Vec<usize>> {
    let mut graph = vec![vec![]; n];
    for &(a, b) in edges.iter() {
        graph[a].push(b);
        graph[b].push(a);
    }
    graph
}

fn bfs(graph: &Vec<Vec<usize>>, n: usize, s: usize) -> Vec<i64> {
    let mut q = VecDeque::new();
    let mut dist = vec![None; n];

    q.push_back(s);
    dist[s] = Some(0);

    while let Some(v) = q.pop_front() {
        for &u in graph[v].iter() {
            if dist[u].is_some() {
                continue;
            }
            q.push_back(u);
            dist[u] = Some(dist[v].unwrap() + 1);
        }
    }

    dist.iter().map(|&x| x.unwrap_or(-1)).collect()
}

fn main() {
    input! {
        n: usize,
        m: usize,
        edges: [(usize, usize); m],
    }

    let graph = edges_to_graph(&edges, n + 1);
    let dist = bfs(&graph, n + 1, 1);

    dist[1..].iter().for_each(|&x| println!("{x}"));
}
