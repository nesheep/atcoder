defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([_, m, _ | tail]) do
    {_, ans} =
      tail
      |> Enum.chunk_every(2)
      |> Enum.reduce({%{}, []}, fn [a, b], {map, acc} ->
        map = Map.update(map, a, [b], &[b | &1])
        acc = if Enum.count(Map.get(map, a, [])) == m, do: [a | acc], else: acc
        {map, acc}
      end)

    Enum.reverse(ans)
  end

  defp output([]), do: nil
  defp output(list), do: Enum.join(list, " ") |> IO.puts()
end
