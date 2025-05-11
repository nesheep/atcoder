defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([n, m]), do: {String.to_integer(n), String.to_integer(m)}

  defp solve({n, m}) do
    Enum.reduce_while(0..m, 0, fn i, acc ->
      x = acc + n ** i
      if x > 1_000_000_000, do: {:halt, :inf}, else: {:cont, x}
    end)
  end

  defp output(:inf), do: "inf"
  defp output(x), do: x
end
