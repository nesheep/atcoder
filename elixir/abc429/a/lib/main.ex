defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([n, m]) do
    1..n
    |> Stream.map(&output(&1, m))
    |> Stream.each(&IO.puts/1)
    |> Stream.run()
  end

  defp output(i, m) when i <= m, do: "OK"
  defp output(_, _), do: "Too Many Requests"
end
