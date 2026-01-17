defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> output()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x) do
    [n, m | ab] = Enum.map(x, &String.to_integer/1)

    g =
      ab
      |> Enum.chunk_every(2)
      |> Enum.reduce(Graph.new(type: :undirected), fn [a, b], g ->
        Graph.add_edge(g, a, b)
      end)

    {n, m, g}
  end

  defp output({n, _, g}) do
    Enum.each(1..n, fn v ->
      neighbors = Graph.neighbors(g, v) |> Enum.sort()
      IO.puts("#{v}: {#{Enum.join(neighbors, ", ")}}")
    end)
  end
end
