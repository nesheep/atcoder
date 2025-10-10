defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([n]), do: String.to_integer(n)

  defp f(i), do: if(rem(i, 2) == 0, do: 1, else: -1) * i ** 3

  defp solve(n) do
    1..n
    |> Enum.map(&f/1)
    |> Enum.sum()
  end
end
