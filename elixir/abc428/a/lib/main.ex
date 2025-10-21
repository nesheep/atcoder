defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([s, a, b, x]) do
    Enum.reduce(1..x, 0, fn
      i, acc when rem(i - 1, a + b) < a -> acc + s
      _, acc -> acc
    end)
  end
end
