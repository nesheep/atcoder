defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
  end

  @trials 20_000

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(v) do
    [n | xy] = Enum.map(v, &String.to_integer/1)
    {n, Enum.chunk_every(xy, 2)}
  end

  defp solve({n, xy}) do
    xy
    |> solve_by_greedy_algo()
    |> optimize(n)
    |> Enum.map(&elem(&1, 1))
  end

  defp solve_by_greedy_algo(xy) do
    Stream.unfold({nil, MapSet.new()}, fn
      :halt -> nil
      {nil, visited} -> {{hd(xy), 1}, {hd(xy), MapSet.put(visited, 1)}}
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
      nil -> {{hd(xy), 1}, :halt}
      {qc, qi} -> {q, {qc, MapSet.put(visited, qi)}}
    end
  end

  defp optimize(acc, n) do
    Enum.reduce(1..@trials, acc, fn _, acc ->
      score1 = score(acc)
      reversed = reverse_range(acc, generate_range(n - 1))
      score2 = score(reversed)
      if score2 < score1, do: reversed, else: acc
    end)
  end

  defp reverse_range(list, {start_idx, end_idx}) when start_idx <= end_idx do
    {before, rest} = Enum.split(list, start_idx)
    {middle, after_part} = Enum.split(rest, end_idx - start_idx + 1)
    before ++ Enum.reverse(middle) ++ after_part
  end

  defp generate_range(n) when n > 1 do
    a = :rand.uniform(n)
    b = :rand.uniform(n)
    if a == b, do: generate_range(n), else: {min(a, b), max(a, b)}
  end

  defp distance([x1, y1], [x2, y2]) do
    :math.sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2)
  end

  defp score(v) do
    v
    |> Enum.map(&elem(&1, 0))
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.map(fn [a, b] -> distance(a, b) end)
    |> Enum.sum()
  end

  defp output(result), do: Enum.each(result, &IO.puts/1)
end
