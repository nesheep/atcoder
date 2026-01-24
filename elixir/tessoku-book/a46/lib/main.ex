defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(v) do
    [n | xy] = Enum.map(v, &String.to_integer/1)
    {n, Enum.chunk_every(xy, 2)}
  end

  defp solve({_n, xy}) do
    Stream.unfold({nil, MapSet.new()}, fn
      :halt -> nil
      {nil, visited} -> {1, {hd(xy), MapSet.put(visited, 1)}}
      {pc, visited} -> next(pc, visited, xy)
    end)
    |> Enum.to_list()
  end

  defp next(pc, visited, xy) do
    q =
      xy
      |> Stream.with_index(1)
      |> Stream.filter(fn {_, i} -> !MapSet.member?(visited, i) end)
      |> Stream.map(fn {kc, _} = k -> {k, distance(pc, kc)} end)
      |> Enum.min_by(&elem(&1, 1), fn -> {nil, nil} end)
      |> elem(0)

    case q do
      nil -> {1, :halt}
      {qc, qi} -> {qi, {qc, MapSet.put(visited, qi)}}
    end
  end

  defp distance([x1, y1], [x2, y2]) do
    :math.sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2)
  end

  defp output(result), do: Enum.each(result, &IO.puts/1)
end
