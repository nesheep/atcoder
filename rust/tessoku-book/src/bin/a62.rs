use proconio::input;

fn dfs(graph: &Vec<Vec<usize>>, v: usize, visited: &mut Vec<bool>) {
    visited[v] = true;
    for &u in graph[v].iter() {
        if visited[u] {
            continue;
        }
        dfs(graph, u, visited);
    }
}

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

    let mut visited = vec![false; n + 1];
    dfs(&graph, 1, &mut visited);
    let ans = visited[1..].iter().all(|&x| x);

    println!(
        "{}",
        if ans {
            "The graph is connected."
        } else {
            "The graph is not connected."
        }
    );
}
