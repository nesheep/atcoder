defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([n | l]) do
    a = l |> Enum.find_index(fn v -> v == 1 end)
    b = l |> Enum.reverse() |> Enum.find_index(fn v -> v == 1 end)
    if a == nil or b == nil, do: 0, else: n - b - a - 1
  end
end
