defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([a, b, _]) when a == b, do: true
  defp solve([_, b, c]) when b == c, do: true
  defp solve([a, _, c]) when c == a, do: true
  defp solve(_), do: false

  defp output(true), do: "Yes"
  defp output(false), do: "No"
end
