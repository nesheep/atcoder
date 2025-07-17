defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([_, k | a]) do
    Enum.reduce(a, fn x, acc ->
      (x * acc)
      |> then(&if &1 < 10 ** k, do: &1, else: 1)
    end)
  end
end
