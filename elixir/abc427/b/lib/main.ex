defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()
  defp parse_input([x]), do: String.to_integer(x)
  defp solve(x), do: a(x)

  defp f(x) when is_integer(x) do
    x
    |> to_string()
    |> String.graphemes()
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum()
  end

  defp a(x) when x == 0, do: 1
  defp a(x) when x == 1, do: f(a(0))
  defp a(x), do: Enum.reduce(1..(x - 1), a(1), fn _, acc -> acc + f(acc) end)
end
