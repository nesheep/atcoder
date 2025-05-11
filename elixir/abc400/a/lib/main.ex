defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([a]), do: String.to_integer(a)

  defp solve(a) when rem(400, a) == 0, do: div(400, a)
  defp solve(_), do: -1
end
