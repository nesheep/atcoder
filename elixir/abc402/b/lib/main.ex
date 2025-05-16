defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> Enum.each(&IO.puts/1)
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([_ | q]), do: parse_query(q, [])
  defp parse_query([], acc), do: acc
  defp parse_query(["1", x | t], acc), do: parse_query(t, [{:q1, x} | acc])
  defp parse_query(["2" | t], acc), do: parse_query(t, [:q2 | acc])

  defp solve(q), do: solve(q, {[], 0})
  defp solve([], {acc, c}), do: acc |> Enum.take(c)
  defp solve([{:q1, x} | t], {acc, c}), do: solve(t, {[x | acc], c})
  defp solve([:q2 | t], {acc, c}), do: solve(t, {acc, c + 1})
end
